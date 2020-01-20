package kuberneteswatcher

import (
	"context"
	"fmt"
	"statusbay/serverutil"

	log "github.com/sirupsen/logrus"

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
func (sm *ServiceManager) Serve() serverutil.StopFunc {

	ctx, cancelFn := context.WithCancel(context.Background())
	stopped := make(chan bool)
	go func() {
		for {
			select {
			case data := <-sm.Watch:
				sm.watch(data)
			case <-ctx.Done():
				log.Warn("Service Manager has been shut down")
				stopped <- true
				return
			}
		}
	}()

	return func() {
		cancelFn()
		<-stopped
	}
}

// watch will start watch on pods changes
func (sm *ServiceManager) watch(watchData WatchData) {

	go func() {
		log.WithFields(log.Fields{
			fmt.Sprintf("%T", watchData.RegistryData): watchData.RegistryData.GetName(),
			"namespace": watchData.Namespace,
		}).Info("Start watch on service")

		log.WithFields(log.Fields{
			"namespace":   watchData.Namespace,
			"list_option": watchData.ListOptions,
		}).Debug("Start watch on service with list options")

		services, err := sm.client.CoreV1().Services(watchData.Namespace).List(watchData.ListOptions)
		if err != nil {
			log.WithError(err).WithField("list_option", watchData.ListOptions.String()).Error("Error when trying to start watch on services")
			return
		}

		if len(services.Items) == 0 {
			log.WithError(err).WithField("list_option", watchData.ListOptions.String()).Info("services not found")
			return
		}

		// todo:: add service name to the DB

	}()

}
