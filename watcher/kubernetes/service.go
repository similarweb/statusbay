package kuberneteswatcher

import (
	"context"
	"sync"

	log "github.com/sirupsen/logrus"

	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// ServiceManager defined service manager struct
type ServiceManager struct {
	client kubernetes.Interface
	Watch  chan WatchData

	dashboardURL string
}

// NewServiceManager create new service instance
func NewServiceManager(kubernetesClientset kubernetes.Interface) *ServiceManager {
	return &ServiceManager{
		client: kubernetesClientset,

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
		if len(services.Items) == 0 {
			watchData.RegistryData.GetName()
			watchEvents(watchData.Ctx, watchData.LogEntry)
			return
		}

	}()

}

// watchEvents will start watch on deployment event messages changes
func (dm *DeploymentManager) watchEvents(ctx context.Context, lg log.Entry, registryDeployment *DeploymentData, listOptions metaV1.ListOptions, namespace string) {
	lg.Info("initializing events watcher")

	watchData := WatchEvents{
		ListOptions: listOptions,
		Namespace:   namespace,
		Ctx:         ctx,
		LogEntry:    lg,
	}

	eventChan := dm.eventManager.Watch(watchData)
	go func() {

		for {
			select {
			case event := <-eventChan:
				registryDeployment.UpdateDeploymentEvents(event)
			case <-ctx.Done():
				lg.Info("stopping events watcher")
				return
			}
		}
	}()
}
