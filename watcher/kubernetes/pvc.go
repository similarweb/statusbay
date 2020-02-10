package kuberneteswatcher

import (
	"context"
	"fmt"
	"statusbay/serverutil"

	log "github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
)

// PvcManager manages the Pvc
type PvcManager struct {
	eventManager *EventsManager
	client       kubernetes.Interface
	Watch        chan WatchData
}

// NewPvcManager creates a new pvc manager objects
func NewPvcManager(kubernetesClientset kubernetes.Interface, eventManager *EventsManager) *PvcManager {
	return &PvcManager{
		client:       kubernetesClientset,
		eventManager: eventManager,
		Watch:        make(chan WatchData),
	}
}

// Serve will start listening on PVC requests
func (pm *PvcManager) Serve() serverutil.StopFunc {

	ctx, cancelFn := context.WithCancel(context.Background())
	stopped := make(chan bool)
	go func() {
		for {
			select {
			case data := <-pm.Watch:
				pm.watchEvents(data)
			case <-ctx.Done():
				log.Warn("Pvc Manager has been shut down")
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

// watchEvents will start watch on pod event messages changes
func (pm *PvcManager) watchEvents(watchData WatchData) {

	log.WithFields(log.Fields{
		fmt.Sprintf("%T", watchData): watchData.RegistryData.GetName(),
		"pod":                        watchData.RegistryData.GetName,
		"namespace":                  watchData.Namespace,
	}).Info("Started watching on Pvc events")

	watchPvcData := WatchEvents{
		ListOptions: watchData.ListOptions,
		Namespace:   watchData.Namespace,
		Ctx:         watchData.Ctx,
	}

	eventChan := pm.eventManager.Watch(watchPvcData)
	go func() {

		for {
			select {
			case event := <-eventChan:
				log.Info(event)
			case <-watchData.Ctx.Done():
				log.WithFields(log.Fields{
					fmt.Sprintf("%T", watchData): watchData.RegistryData.GetName(),
					"pod":                        watchData.RegistryData.GetName,
					"namespace":                  watchData.Namespace,
				}).Info("Stopped watching Pvc events")
				return
			}

		}
	}()

}
