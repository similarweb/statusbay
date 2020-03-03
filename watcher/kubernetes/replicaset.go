package kuberneteswatcher

import (
	"context"
	"sync"

	log "github.com/sirupsen/logrus"

	appsV1 "k8s.io/api/apps/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"

	"k8s.io/client-go/kubernetes"
)

// WatchReplica defined replicaset watch received message
type WatchReplica struct {
	LogEntry        log.Entry
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
func (rm *ReplicaSetManager) Serve(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)

	go func() {
		for {
			select {
			case replicaSets := <-rm.Watch:
				rm.watch(replicaSets)
			case <-ctx.Done():
				log.Warn("replicaset manager has been shut down")
				wg.Done()
				return
			}
		}
	}()

}

// watch will start watch on replicaset changes
func (rm *ReplicaSetManager) watch(replicaData WatchReplica) {

	go func() {

		replicaData.LogEntry.Info("start watch on replicasets")
		replicaData.LogEntry.WithField("list_option", replicaData.ListOptions).Debug("list option for replicaset filtering")

		//List of replicaset changes events
		firstInit := map[string]bool{}

		watcher, err := rm.client.AppsV1().ReplicaSets(replicaData.Namespace).Watch(replicaData.ListOptions)

		if err != nil {
			replicaData.LogEntry.WithField("list_option", replicaData.ListOptions).WithError(err).Error("error occured trying to start watch on replicasets")
			return
		}

		for {
			select {
			case event, watch := <-watcher.ResultChan():
				if !watch {
					replicaData.LogEntry.Warn("replicaset watch was stopped, channel was closed")

					return
				}

				replicaset, ok := event.Object.(*appsV1.ReplicaSet)
				if !ok {
					replicaData.LogEntry.WithField("object", event.Object).Warn("failed to parse replicaset watch data")
					continue
				}

				lg := replicaData.LogEntry.WithFields(log.Fields{
					"replicaset_name": replicaset.GetName(),
				})

				if _, found := firstInit[replicaset.Name]; !found {

					lg.Debug("discovered new replicaset")
					firstInit[replicaset.Name] = true
					replicaData.Registry.InitReplicaset(replicaset.GetName())

					if value, found := replicaset.Spec.Selector.MatchLabels["pod-template-hash"]; found {
						lg.Debug("selector `pod-template-hash` found in replicaset")

						podListOptions := metaV1.ListOptions{LabelSelector: labels.SelectorFromSet(map[string]string{
							"pod-template-hash": value,
						}).String()}

						rm.podsManager.Watch <- WatchData{
							ListOptions:  podListOptions,
							RegistryData: replicaData.Registry,
							Namespace:    replicaData.Namespace,
							Ctx:          replicaData.Ctx,
							LogEntry:     *lg,
						}

					} else {
						lg.Warn("selector `pod-template-hash` not found in replicaset")
						continue
					}

					eventListOptions := metaV1.ListOptions{FieldSelector: labels.SelectorFromSet(map[string]string{
						"involvedObject.name": replicaset.Name,
						"involvedObject.kind": "ReplicaSet",
					}).String(),
					}

					rm.watchEvents(replicaData.Ctx, *lg, replicaData.Registry, eventListOptions, replicaset.GetName(), replicaData.Namespace)

				}

				replicaData.Registry.UpdateReplicasetStatus(replicaset.GetName(), replicaset.Status)

			case <-replicaData.Ctx.Done():
				replicaData.LogEntry.Debug("replicaset watch was stopped, got ctx done signal")
				watcher.Stop()
				return
			}
		}
	}()

}

// watchEvents will start watch on replicaset event messages changes
func (rm *ReplicaSetManager) watchEvents(ctx context.Context, lg log.Entry, registryDeployment *DeploymentData, listOptions metaV1.ListOptions, replicasetName, namespace string) {

	lg.Info("start watching replicaset events")
	watchData := WatchEvents{
		ListOptions: listOptions,
		Namespace:   namespace,
		Ctx:         ctx,
		LogEntry:    lg,
	}

	eventChan := rm.eventManager.Watch(watchData)
	go func() {

		for {
			select {
			case event := <-eventChan:
				registryDeployment.UpdateReplicasetEvents(replicasetName, event)
			case <-ctx.Done():
				lg.Info("stop watching replicaset events")
				return
			}

		}
	}()

}
