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
	Watch         chan WatchData
	podsFirstInit map[string]bool
	mutex         *sync.RWMutex
}

// NewPodsManager create new pods instance
func NewPodsManager(kubernetesClientset kubernetes.Interface, eventManager *EventsManager) *PodsManager {
	return &PodsManager{
		client:        kubernetesClientset,
		eventManager:  eventManager,
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

		lg := watchData.LogEntry.WithFields(log.Fields{
			"name": watchData.RegistryData.GetName(),
		})

		watchData.LogEntry.Info("Start watch on pods")

		lg.WithField("list_option", watchData.ListOptions).Debug("Pod list options")

		watcher, err := pm.client.CoreV1().Pods(watchData.Namespace).Watch(watchData.ListOptions)
		if err != nil {
			lg.WithError(err).WithField("list_option", watchData.ListOptions.String()).Error("Error when trying to start watch on pods")
			return
		}
		for {
			select {
			case event, watch := <-watcher.ResultChan():
				if !watch {
					lg.WithFields(log.Fields{
						"list_options": watchData.ListOptions.String(),
					}).Warn("Pods watch was stopped. Channel was closed")
					return
				}

				pod, ok := event.Object.(*v1.Pod)
				if !ok {
					lg.Warn("Failed to parse pod watch data")
					continue
				}

				podLog := lg.WithFields(log.Fields{
					"pod": pod.Name,
				})

				//If it is the first time that we got the pod, we are start watch on pod events & send the pod to registry
				if found := pm.loadPodFirstInit(pod.Name); !found {

					podLog.Debug("New pod found")

					pm.storePodFirstInit(pod.Name, true)
					watchData.RegistryData.NewPod(pod)
					eventListOptions := metaV1.ListOptions{FieldSelector: labels.SelectorFromSet(map[string]string{
						"involvedObject.name": pod.GetName(),
						"involvedObject.kind": "Pod",
					}).String(),
					}
					go pm.watchEvents(watchData.Ctx, *podLog, watchData.RegistryData, eventListOptions, pod.Namespace, pod.GetName())

				}

				status := string(pod.Status.Phase)
				podLog.WithFields(log.Fields{
					"count": len(pod.Status.ContainerStatuses),
				}).Debug("List of pod status Container statuses")

				for _, container := range pod.Status.ContainerStatuses {

					containerLog := podLog.WithFields(log.Fields{
						"container_name": container.Name,
						"container_id":   container.ContainerID,
					})

					if container.State.Waiting != nil {

						message := container.State.Waiting.Reason
						containerLog.WithField("message", message).Debug("Container statue is waiting")
						if container.State.Waiting.Message != "" {
							message = fmt.Sprintf("%s - %s", message, container.State.Waiting.Message)
						}

						eventMessage := EventMessages{
							Message: message,
							Time:    time.Now().UnixNano(),
						}
						watchData.RegistryData.UpdatePodEvents(pod.GetName(), eventMessage)
						status = container.State.Waiting.Reason
					}

					if container.State.Terminated != nil {

						message := container.State.Terminated.Reason
						containerLog.WithField("message", message).Debug("Container statue is termenated")
						if container.State.Terminated.Message != "" {
							message = fmt.Sprintf("%s - %s", message, container.State.Terminated.Message)
						}

						eventMessage := EventMessages{
							Message:             message,
							Time:                container.State.Terminated.StartedAt.UnixNano(),
							ReportingController: container.State.Terminated.ContainerID,
						}
						watchData.RegistryData.UpdatePodEvents(pod.GetName(), eventMessage)
						status = container.State.Terminated.Reason
					}
				}

				if pod.GetDeletionTimestamp() != nil {
					status = "Terminated"
				}
				podLog.WithField("status", status).Debug("Pod status")
				watchData.RegistryData.UpdatePod(pod, status)

			case <-watchData.Ctx.Done():
				watchData.LogEntry.Info("Pod watcher was stopped. Got ctx done signal")
				watcher.Stop()
				return
			}
		}
	}()

}

// watchEvents will start watch on pod event messages changes
func (pm *PodsManager) watchEvents(ctx context.Context, lg log.Entry, registryData RegistryData, listOptions metaV1.ListOptions, namespace, podName string) {

	lg.Info("Start watch on pod events")

	watchData := WatchEvents{
		ListOptions: listOptions,
		Namespace:   namespace,
		Ctx:         ctx,
		LogEntry:    lg,
	}
	eventChan := pm.eventManager.Watch(watchData)
	go func() {

		for {
			select {
			case event := <-eventChan:
				registryData.UpdatePodEvents(podName, event)
			case <-ctx.Done():
				lg.Info("Stop watch on pod events")
				return
			}

		}
	}()

}
