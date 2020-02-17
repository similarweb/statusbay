package kuberneteswatcher

import (
	"context"
	"fmt"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	v1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
)

// PodsManager defined pods manager struct
type PodsManager struct {
	client        kubernetes.Interface
	eventManager  *EventsManager
	pvcManager    *PvcManager
	Watch         chan WatchData
	podsFirstInit map[string]bool
	mutex         *sync.RWMutex
}

// NewPodsManager create new pods instance
func NewPodsManager(kubernetesClientset kubernetes.Interface, eventManager *EventsManager, pvcManager *PvcManager) *PodsManager {
	return &PodsManager{
		client:        kubernetesClientset,
		eventManager:  eventManager,
		pvcManager:    pvcManager,
		podsFirstInit: map[string]bool{},
		mutex:         &sync.RWMutex{},
		Watch:         make(chan WatchData),
	}
}

// storePodFirstInit will set if some pod appears for the first time true == first time
func (pm *PodsManager) storePodFirstInit(key string, val bool) {
	pm.mutex.Lock()
	pm.podsFirstInit[key] = val
	pm.mutex.Unlock()
}

// loadPodFirstInit will return true if pod exists or false otherwise
func (pm *PodsManager) loadPodFirstInit(key string) bool {
	pm.mutex.RLock()
	exist := pm.podsFirstInit[key]
	pm.mutex.RUnlock()
	return exist
}

// Serve will start listening on pods request
func (pm *PodsManager) Serve(ctx context.Context, wg *sync.WaitGroup) {

	go func() {
		for {
			select {
			case data := <-pm.Watch:
				pm.watch(data)
			case <-ctx.Done():
				log.Warn("Pods Manager has been shut down")
				wg.Done()
				return
			}
		}
	}()

}

// watch will start watch on pods changes
func (pm *PodsManager) watch(watchData WatchData) {
	go func() {
		log.WithFields(log.Fields{
			fmt.Sprintf("%T", watchData.RegistryData): watchData.RegistryData.GetName(),
			"namespace": watchData.Namespace,
		}).Info("Start watch on pods")

		log.WithFields(log.Fields{
			"namespace":   watchData.Namespace,
			"list_option": watchData.ListOptions,
		}).Debug("Start watch on pods with list options")

		watcher, err := pm.client.CoreV1().Pods(watchData.Namespace).Watch(watchData.ListOptions)
		if err != nil {
			log.WithError(err).WithField("list_option", watchData.ListOptions.String()).Error("Error when trying to start watch on pods")
			return
		}
		for {
			select {
			case event, watch := <-watcher.ResultChan():
				if !watch {
					log.WithFields(log.Fields{
						fmt.Sprintf("%T", watchData.RegistryData): watchData.RegistryData.GetName(),
						"list_options": watchData.ListOptions.String(),
						"namespace":    watchData.Namespace,
					}).Warn("Pods watch was stopped. Channel was closed")
					return
				}

				pod, ok := event.Object.(*v1.Pod)
				if !ok {
					log.Warn("Failed to parse pod watch data")
					continue
				}

				//If it is the first time that we got the pod, we are start watch on pod events & send the pod to registry
				if found := pm.loadPodFirstInit(pod.Name); !found {
					pm.storePodFirstInit(pod.Name, true)
					watchData.RegistryData.NewPod(pod)
					eventListOptions := metaV1.ListOptions{FieldSelector: labels.SelectorFromSet(map[string]string{
						"involvedObject.name": pod.GetName(),
						"involvedObject.kind": "Pod",
					}).String(),
					}
					go pm.watchEvents(watchData.Ctx, watchData.RegistryData, eventListOptions, pod.Namespace, pod.GetName())

					for _, volume := range pod.Spec.Volumes {
						pvc := volume.VolumeSource.PersistentVolumeClaim
						// There are some cases where pvc is nil , when Kubernetes creates it's own certificates on the pods
						// It mounts another volume for system use which does not have PersistentVolumeClaim
						if pvc != nil {
							PvcEventListOptions := metaV1.ListOptions{FieldSelector: labels.SelectorFromSet(map[string]string{
								"metadata.name": pvc.ClaimName}).String()}

							pm.pvcManager.Watch <- WatchPvcData{
								ListOptions:  PvcEventListOptions,
								RegistryData: watchData.RegistryData,
								Namespace:    pod.Namespace,
								Pod:          pod.Name,
								Ctx:          watchData.Ctx,
							}
						}
					}
				}

				status := string(pod.Status.Phase)
				for _, container := range pod.Status.ContainerStatuses {

					if container.State.Waiting != nil {
						message := container.State.Waiting.Reason

						if container.State.Waiting.Message != "" {
							message = fmt.Sprintf("%s - %s", message, container.State.Waiting.Message)
						}

						eventMessage := EventMessages{
							Message: message,
							Time:    time.Now().UnixNano(),
						}
						watchData.RegistryData.UpdatePodEvents(pod.GetName(), "", eventMessage)
						status = container.State.Waiting.Reason
					}

					if container.State.Terminated != nil {

						message := container.State.Terminated.Reason

						if container.State.Terminated.Message != "" {
							message = fmt.Sprintf("%s - %s", message, container.State.Terminated.Message)
						}

						eventMessage := EventMessages{
							Message:             message,
							Time:                container.State.Terminated.StartedAt.UnixNano(),
							ReportingController: container.State.Terminated.ContainerID,
						}
						watchData.RegistryData.UpdatePodEvents(pod.GetName(), "", eventMessage)
						status = container.State.Terminated.Reason
					}
				}

				if pod.GetDeletionTimestamp() != nil {
					status = "Terminated"
				}
				watchData.RegistryData.UpdatePod(pod, status)

			case <-watchData.Ctx.Done():
				log.WithFields(log.Fields{
					"selector":  watchData.ListOptions.String(),
					"namespace": watchData.Namespace,
				}).Info("Pod watch was stopped. Got ctx done signal")
				watcher.Stop()
				return
			}
		}
	}()

}

// watchEvents will start watch on pod event messages changes
func (pm *PodsManager) watchEvents(ctx context.Context, registryData RegistryData, listOptions metaV1.ListOptions, namespace, podName string) {

	log.WithFields(log.Fields{
		fmt.Sprintf("%T", registryData): registryData.GetName(),
		"pod":                           podName,
		"namespace":                     namespace,
	}).Info("Start watch on pod events")

	watchData := WatchEvents{
		ListOptions: listOptions,
		Namespace:   namespace,
		Ctx:         ctx,
	}
	eventChan := pm.eventManager.Watch(watchData)
	go func() {

		for {
			select {
			case event := <-eventChan:
				registryData.UpdatePodEvents(podName, "", event)
			case <-ctx.Done():
				log.WithFields(log.Fields{
					fmt.Sprintf("%T", registryData): registryData.GetName(),
					"pod":                           podName,
					"namespace":                     namespace,
				}).Info("Stop watch on pod events")
				return
			}

		}
	}()

}
