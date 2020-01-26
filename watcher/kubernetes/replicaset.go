package kuberneteswatcher

import (
	"context"
	"statusbay/serverutil"

	log "github.com/sirupsen/logrus"

	appsV1 "k8s.io/api/apps/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"

	"k8s.io/client-go/kubernetes"
)

// WatchReplica defined replicaset watch received message
type WatchReplica struct {
	DesiredReplicas int32
	ListOptions     metaV1.ListOptions

	Registry  *DeploymentData
	Namespace string
	Ctx       context.Context
}

// ReplicaSetManager defined replicaset manager struct
type ReplicaSetManager struct {
	eventManager *EventsManager
	Watch        chan WatchReplica
	client       kubernetes.Interface
	podsManager  *PodsManager
	int64
}

// NewReplicasetManager create new replicaset instance
func NewReplicasetManager(kubernetesClientset kubernetes.Interface, eventManager *EventsManager, podsManager *PodsManager) *ReplicaSetManager {
	return &ReplicaSetManager{
		client:       kubernetesClientset,
		eventManager: eventManager,
		podsManager:  podsManager,
		Watch:        make(chan WatchReplica),
	}
}

// Serve will start listening replicaset request
func (rm *ReplicaSetManager) Serve() serverutil.StopFunc {

	ctx, cancelFn := context.WithCancel(context.Background())
	stopped := make(chan bool)
	go func() {
		for {
			select {
			case replicaSets := <-rm.Watch:
				rm.watch(replicaSets)
			case <-ctx.Done():
				log.Warn("Replicaset Manager has been shut down")
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

// watch will start watch on replicaset changes
func (rm *ReplicaSetManager) watch(replicaData WatchReplica) {

	go func() {

		log.WithFields(log.Fields{
			"deployment": replicaData.Registry.Deployment.Name,
			"namespace":  replicaData.Namespace,
		}).Info("Start watch on replicasets")

		log.WithFields(log.Fields{
			"deployment_name": replicaData.Registry.Deployment.Name,
			"namespace":       replicaData.Namespace,
			"list_option":     replicaData.ListOptions,
		}).Warn("Start watch on replicasets with list options")

		//List of replicaset changes events
		firstInit := map[string]bool{}

		watcher, err := rm.client.AppsV1().ReplicaSets(replicaData.Namespace).Watch(replicaData.ListOptions)

		if err != nil {
			log.WithError(err).WithFields(log.Fields{
				"deployment":  replicaData.Registry.Deployment.Name,
				"namespace":   replicaData.Namespace,
				"list_option": replicaData.ListOptions.String(),
			}).Error("Error when trying to start watch on replicasets")
			return
		}

		for {
			select {
			case event, watch := <-watcher.ResultChan():
				if !watch {
					log.WithFields(log.Fields{
						"deployment":   replicaData.Registry.Deployment.Name,
						"list_options": replicaData.ListOptions.String(),
						"namespace":    replicaData.Namespace,
					}).Warn("Replicaset watch was stopped. Channel was closed")

					return
				}

				replicaset, ok := event.Object.(*appsV1.ReplicaSet)
				if !ok {
					log.WithField("object", event.Object).Warn("Failed to parse replicaset watch data")
					continue
				}

				if _, found := firstInit[replicaset.Name]; !found {
					firstInit[replicaset.Name] = true
					replicaData.Registry.InitReplicaset(replicaset.GetName())

					if value, found := replicaset.Spec.Selector.MatchLabels["pod-template-hash"]; found {

						podListOptions := metaV1.ListOptions{LabelSelector: labels.SelectorFromSet(map[string]string{
							"pod-template-hash": value,
						}).String()}

						rm.podsManager.Watch <- WatchData{
							ListOptions:  podListOptions,
							RegistryData: replicaData.Registry,
							Namespace:    replicaData.Namespace,
							Ctx:          replicaData.Ctx,
						}

					} else {
						log.WithFields(log.Fields{
							"deployment": replicaData.Registry.Deployment.Name,
							"name":       replicaset.GetName(),
							"namespace":  replicaset.GetNamespace(),
							"selector":   replicaset.Spec.Selector.String(),
						}).Error("Selector `pod-template-hash` not found in replicaset. cannot start watch on pods")
						continue
					}

					eventListOptions := metaV1.ListOptions{FieldSelector: labels.SelectorFromSet(map[string]string{
						"involvedObject.name": replicaset.Name,
						"involvedObject.kind": "ReplicaSet",
					}).String(),
					}

					rm.watchEvents(replicaData.Ctx, replicaData.Registry, eventListOptions, replicaset.GetName(), replicaData.Namespace)

				}

				replicaData.Registry.UpdateReplicasetStatus(replicaset.GetName(), replicaset.Status)

			case <-replicaData.Ctx.Done():
				log.WithFields(log.Fields{
					"selector":  replicaData.ListOptions.String(),
					"namespace": replicaData.Namespace,
				}).Debug("Replicaset watch was stopped. Got ctx done signal")
				watcher.Stop()
				return
			}
		}
	}()

}

// watchEvents will start watch on replicaset event messages changes
func (rm *ReplicaSetManager) watchEvents(ctx context.Context, registryDeployment *DeploymentData, listOptions metaV1.ListOptions, replicasetName, namespace string) {

	log.WithFields(log.Fields{
		"deployment": registryDeployment.Deployment.Name,
		"replicaset": replicasetName,
		"namespace":  namespace,
	}).Info("Start watch on replicaset events")

	watchData := WatchEvents{
		ListOptions: listOptions,
		Namespace:   namespace,
		Ctx:         ctx,
	}
	eventChan := rm.eventManager.Watch(watchData)
	go func() {

		for {
			select {
			case event := <-eventChan:
				registryDeployment.UpdateReplicasetEvents(replicasetName, event)
			case <-ctx.Done():
				log.WithFields(log.Fields{
					"deployment": registryDeployment.Deployment.Name,
					"replicaset": replicasetName,
					"namespace":  namespace,
				}).Info("Stop watch on replicaset events")
				return
			}

		}
	}()

}
