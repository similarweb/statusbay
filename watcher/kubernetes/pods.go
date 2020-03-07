package kuberneteswatcher

import (
	"bufio"
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
	wg.Add(1)

	go func() {
		for {
			select {
			case data := <-pm.Watch:
				pm.watch(data)
			case <-ctx.Done():
				log.Warn("pods manager has been shut down")
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

		watchData.LogEntry.Info("initializing pods watcher")

		lg.WithField("list_option", watchData.ListOptions).Debug("pod list options")

		watcher, err := pm.client.CoreV1().Pods(watchData.Namespace).Watch(watchData.ListOptions)
		if err != nil {
			lg.WithError(err).WithField("list_option", watchData.ListOptions.String()).Error("error when trying to start watch on pods")
			return
		}
		for {
			select {
			case event, watch := <-watcher.ResultChan():
				if !watch {
					lg.WithFields(log.Fields{
						"list_options": watchData.ListOptions.String(),
					}).Warn("pods watch was stopped, channel was closed")
					return
				}

				pod, ok := event.Object.(*v1.Pod)
				if !ok {
					lg.Warn("failed to parse pod watch data")
					continue
				}

				podLog := lg.WithFields(log.Fields{
					"pod": pod.Name,
				})

				//If it is the first time that we got the pod, we are start watch on pod events & send the pod to registry
				if found := pm.loadPodFirstInit(pod.Name); !found {

					podLog.Debug("new pod discovered")

					pm.storePodFirstInit(pod.Name, true)
					watchData.RegistryData.NewPod(pod)
					eventListOptions := metaV1.ListOptions{FieldSelector: labels.SelectorFromSet(map[string]string{
						"involvedObject.name": pod.GetName(),
						"involvedObject.kind": "Pod",
					}).String(),
					}
					pm.watchEvents(watchData.Ctx, *podLog, watchData.RegistryData, eventListOptions, pod.Namespace, pod.GetName())

					for _, volume := range pod.Spec.Volumes {
						pvc := volume.VolumeSource.PersistentVolumeClaim
						// There are some cases where pvc is nil , when Kubernetes creates it's own certificates on the pods
						// It mounts another volume for system use which does not have PersistentVolumeClaim
						if pvc != nil {
							podLog.WithFields(log.Fields{"pvc": pvc.ClaimName}).Debug("pods watcher found a new pvc")
							PvcEventListOptions := metaV1.ListOptions{FieldSelector: labels.SelectorFromSet(map[string]string{
								"metadata.name": pvc.ClaimName}).String()}

							pm.pvcManager.Watch <- WatchPvcData{
								LogEntry:     *podLog,
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

				if status == string(v1.PodRunning) {
					for _, container := range pod.Spec.Containers {
						pm.podLogs(watchData.Ctx, *podLog, watchData.RegistryData, watchData.Namespace, pod.Name, container.Name)
					}
				}

				podLog.WithFields(log.Fields{
					"count": len(pod.Status.ContainerStatuses),
				}).Debug("list of pod status container statuses")

				for _, container := range pod.Status.ContainerStatuses {

					containerLog := podLog.WithFields(log.Fields{
						"container_name": container.Name,
						"container_id":   container.ContainerID,
					})

					if container.State.Waiting != nil {

						message := container.State.Waiting.Reason
						containerLog.WithField("message", message).Debug("container status is waiting")
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
						containerLog.WithField("message", message).Debug("container status is terminated")
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
				podLog.WithField("status", status).Debug("pod status")
				watchData.RegistryData.UpdatePod(pod, status)

			case <-watchData.Ctx.Done():
				watchData.LogEntry.Info("pod watcher was stopped, got ctx done signal")
				watcher.Stop()
				return
			}
		}
	}()

}

// watchEvents will start watch on pod event messages changes
func (pm *PodsManager) watchEvents(ctx context.Context, lg log.Entry, registryData RegistryData, listOptions metaV1.ListOptions, namespace, podName string) {

	lg.Info("starting to watch pod events")

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
				registryData.UpdatePodEvents(podName, "", event)
			case <-ctx.Done():
				lg.Info("stop watching pod events")
				return
			}

		}
	}()

}

// podLogs open a container logs steam for getting the STDOUT of container
func (pm *PodsManager) podLogs(ctx context.Context, lg log.Entry, registryData RegistryData, namespace, podName, containerName string) {

	// create new container object in registry,
	// if error returned, it mean that the object already exists, mean the steam already created
	err := registryData.NewContainer(podName, containerName)
	if err != nil {
		return
	}
	lgContainer := lg.WithField("container_name", containerName)

	go func() {
		lgContainer.Info("open log stream")
		podLogOpts := v1.PodLogOptions{
			Container:  containerName,
			Follow:     true,
			Previous:   false,
			Timestamps: true,
		}
		req := pm.client.CoreV1().Pods(namespace).GetLogs(podName, &podLogOpts)
		streamIO, err := req.Stream()
		if err != nil {
			lgContainer.Error(err)
			return
		}
		defer streamIO.Close()
		r := bufio.NewReader(streamIO)

		for {
			select {
			case <-ctx.Done():
				lgContainer.Info("close log stream")
				return
			default:
				bytes, err := r.ReadBytes('\n')
				if err != nil {
					lgContainer.WithError(err).Info("failed to read stream bytes")
					continue
				}
				registryData.AddContainerLog(podName, containerName, string(bytes))
			}
		}
	}()

}
