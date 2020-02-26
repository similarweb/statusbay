package kuberneteswatcher

import (
	"context"
	"sync"

	log "github.com/sirupsen/logrus"

	v1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
)

// ServiceManager defined service manager struct
type ServiceManager struct {
	eventManager *EventsManager
	client       kubernetes.Interface
	Watch        chan WatchData

	dashboardURL string
}

// NewServiceManager create new service instance
func NewServiceManager(kubernetesClientset kubernetes.Interface, eventManager *EventsManager) *ServiceManager {
	return &ServiceManager{
		client:       kubernetesClientset,
		eventManager: eventManager,

		Watch: make(chan WatchData),
	}
}

// Serve will start listening on pods request
func (sm *ServiceManager) Serve(ctx context.Context, wg *sync.WaitGroup) {

	go func() {
		for {
			select {
			case data := <-sm.Watch:
				sm.watch(data)
			case <-ctx.Done():
				log.Warn("service manager has been shut down")
				wg.Done()
				return
			}
		}
	}()

}

// watch will start watch on pods changes
func (sm *ServiceManager) watch(watchData WatchData) {

	go func() {

		watchData.LogEntry.Info("start watching service")
		watchData.LogEntry.WithField("list_option", watchData.ListOptions).Debug("start watch on service with list options")

		watcher, err := sm.client.CoreV1().Services(watchData.Namespace).Watch(watchData.ListOptions)

		if err != nil {
			watchData.LogEntry.WithError(err).Error("could not start watch on service")
			return
		}
		firstInit := map[string]bool{}

		for {
			select {
			case event, watch := <-watcher.ResultChan():
				if !watch {
					watchData.LogEntry.Warn("service watcher was stopped, channel was closed")
					return
				}
				svc, isOk := event.Object.(*v1.Service)
				if !isOk {
					watchData.LogEntry.WithField("object", event.Object).Warn("failed to parse service watcher data")
					continue
				}
				if _, found := firstInit[svc.GetName()]; !found {
					eventListOptions := metaV1.ListOptions{FieldSelector: labels.SelectorFromSet(map[string]string{
						"involvedObject.name": svc.GetName(),
						"involvedObject.kind": "Service",
					}).String(),
					}
					watchData.RegistryData.NewService(svc)
					sm.watchEvents(watchData.Ctx, watchData.LogEntry, watchData.RegistryData, eventListOptions, svc.GetName(), watchData.Namespace)
				}

			case <-watchData.Ctx.Done():
				watchData.LogEntry.Debug("service watcher was stopped, got ctx done signal")
				watcher.Stop()
				return
			}
		}

	}()

}

// watchEvents will start watch on deployment event messages changes
func (sm *ServiceManager) watchEvents(ctx context.Context, lg log.Entry, registryDeployment RegistryData, listOptions metaV1.ListOptions, serviceName, namespace string) {
	lg.Info("initializing events watcher")

	watchData := WatchEvents{
		ListOptions: listOptions,
		Namespace:   namespace,
		Ctx:         ctx,
		LogEntry:    lg,
	}

	eventChan := sm.eventManager.Watch(watchData)
	go func() {

		for {
			select {
			case event := <-eventChan:
				registryDeployment.UpdateServiceEvents(serviceName, event)
			case <-ctx.Done():
				lg.Info("stopping events watcher")
				return
			}
		}
	}()
}
