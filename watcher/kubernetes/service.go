package kuberneteswatcher

import (
	"context"
	"sync"

	log "github.com/sirupsen/logrus"

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

		services, err := sm.client.CoreV1().Services(watchData.Namespace).List(watchData.ListOptions)
		if err != nil {
			watchData.LogEntry.WithError(err).WithField("list_option", watchData.ListOptions.String()).Error("error when trying to start watch on services")
			return
		}

		watchData.LogEntry.WithField("service_count", len(services.Items)).Info("found related services")

		for _, svc := range services.Items {

			eventListOptions := metaV1.ListOptions{FieldSelector: labels.SelectorFromSet(map[string]string{
				"involvedObject.name": svc.GetName(),
				"involvedObject.kind": "Service",
			}).String(),
			}
			watchData.RegistryData.NewService(&svc)
			sm.watchEvents(watchData.Ctx, watchData.LogEntry, watchData.RegistryData, eventListOptions, svc.GetName(), watchData.Namespace)
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
