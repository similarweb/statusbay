package kuberneteswatcher

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"strings"
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
	client                 kubernetes.Interface
	eventManager           *EventsManager
	pvcManager             *PvcManager
	Watch                  chan WatchData
	podsFirstInit          map[string]bool
	podsLogsFirstInit      map[string]bool
	mutexPodsFirstInit     *sync.RWMutex
	mutexPodsLogsFirstInit *sync.RWMutex
	absoluteLogsPodPath    string
}

// NewPodsManager create new pods instance
func NewPodsManager(kubernetesClientset kubernetes.Interface, eventManager *EventsManager, pvcManager *PvcManager, absoluteLogsPodPath string) *PodsManager {
	return &PodsManager{
		client:                 kubernetesClientset,
		eventManager:           eventManager,
		pvcManager:             pvcManager,
		podsFirstInit:          map[string]bool{},
		podsLogsFirstInit:      map[string]bool{},
		mutexPodsFirstInit:     &sync.RWMutex{},
		mutexPodsLogsFirstInit: &sync.RWMutex{},
		Watch:                  make(chan WatchData),
		absoluteLogsPodPath:    absoluteLogsPodPath,
	}
}

// storePodFirstInit will set if some pod appears for the first time true == first time
func (pm *PodsManager) storePodFirstInit(key string, val bool) {
	pm.mutexPodsFirstInit.Lock()
	pm.podsFirstInit[key] = val
	pm.mutexPodsFirstInit.Unlock()
}

// loadPodFirstInit will return true if pod exists or false otherwise
func (pm *PodsManager) loadPodFirstInit(key string) bool {
	pm.mutexPodsFirstInit.RLock()
	exist := pm.podsFirstInit[key]
	pm.mutexPodsFirstInit.RUnlock()
	return exist
}

// storePodFirstInit will set if some pod appears for the first time true == first time
func (pm *PodsManager) storePodLogsFirstInit(key string, val bool) {
	pm.mutexPodsLogsFirstInit.Lock()
	pm.podsLogsFirstInit[key] = val
	pm.mutexPodsLogsFirstInit.Unlock()
}

// loadPodFirstInit will return true if pod exists or false otherwise
func (pm *PodsManager) loadPodLogsFirstInit(key string) bool {
	pm.mutexPodsLogsFirstInit.RLock()
	exist := pm.podsLogsFirstInit[key]
	pm.mutexPodsLogsFirstInit.RUnlock()
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
					go pm.watchEvents(watchData.Ctx, *podLog, watchData.RegistryData, eventListOptions, pod.Namespace, pod.GetName())

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
						folder := pm.getFolderLogsPath(watchData.ApplyID, pod.GetName())
						pullLogPathFile := pm.getPodLogFilePath(folder, container.Name)
						if found := pm.loadPodLogsFirstInit(pullLogPathFile); !found {
							pm.storePodLogsFirstInit(pullLogPathFile, true)
							os.MkdirAll(folder, os.ModePerm)
							pm.podLogs(watchData.Ctx, *podLog, pullLogPathFile, watchData.Namespace, pod.Name, container.Name)
						}
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

// getFolderLogsPath returns the absolute folder path of the current applyID for pod continers logs
func (pm *PodsManager) getFolderLogsPath(applyID, podName string) string {
	return fmt.Sprintf("%s/%s/%s", pm.absoluteLogsPodPath, applyID, podName)
}

// getPodLogFilePath returns the container file path
func (pm *PodsManager) getPodLogFilePath(folder, containerName string) string {
	return fmt.Sprintf("%s/%s.log", folder, containerName)
}

// podLogs open a container logs steam for getting the STDOUT of container
func (pm *PodsManager) podLogs(ctx context.Context, lg log.Entry, absoluteFilePath, namespace, podName, containerName string) {

	lgContainer := lg.WithFields(log.Fields{
		"container_name": containerName,
		"file_path":      absoluteFilePath,
	})

	go func() {
		file, err := os.OpenFile(fmt.Sprintf("%s", absoluteFilePath), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
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
			lgContainer.WithError(err).Error("could not open container log stream")
			return
		}
		defer streamIO.Close()
		r := bufio.NewReader(streamIO)

		for {
			select {
			case <-ctx.Done():
				lgContainer.Info("close log stream")
				err = file.Close()
				if err != nil {
					lgContainer.WithError(err).Error("could not close pod logs file")
				}
				return
			default:
				bytes, err := r.ReadBytes('\n')
				if err != nil {
					lgContainer.WithError(err).Info("failed to read stream bytes")
					continue
				}

				line := strings.NewReader(string(bytes))
				_, err = io.Copy(file, line)
				if err != nil {
					lgContainer.WithError(err).Error(file)
				}

			}
		}
	}()

}
