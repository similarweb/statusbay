package kuberneteswatcher

import (
	"context"
	"statusbay/watcher/kubernetes/common"
	"sync"
	"time"

	"github.com/mitchellh/hashstructure"
	log "github.com/sirupsen/logrus"
	appsV1 "k8s.io/api/apps/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	eventwatch "k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
)

// StatefulsetManager Struct definition
type StatefulsetManager struct {
	// Kubernetes client interface
	client kubernetes.Interface

	// Event manager will be owner to start watch on deployment events
	eventManager *EventsManager

	// Registry manager will be owner to manage the running / new statefulset deployment
	registryManager *RegistryManager

	// ServiceManager will
	serviceManager *ServiceManager

	// Holds the revisions of deamonset and statefulset
	controllerRevManager ControllerRevision

	// Max time we want to watch a statefulset deployment
	maxDeploymentTime int64
}

// NewStatefulsetManager creates a new instance to manage statefulset related resources
func NewStatefulsetManager(k8sClient kubernetes.Interface, eventManager *EventsManager, registryManager *RegistryManager, serviceManager *ServiceManager, controllerRevisionManager ControllerRevision, maxDeploymentTime time.Duration) *StatefulsetManager {
	return &StatefulsetManager{
		client:               k8sClient,
		eventManager:         eventManager,
		registryManager:      registryManager,
		serviceManager:       serviceManager,
		controllerRevManager: controllerRevisionManager,
		maxDeploymentTime:    int64(maxDeploymentTime.Seconds()),
	}
}

//Serve Will serve the watch channels of statefulset
func (ssm *StatefulsetManager) Serve(ctx context.Context, wg *sync.WaitGroup) {

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Warn("Statefulsets Manager has been shut down")
				wg.Done()
				return
			}
		}
	}()

	// Continue watching on running statefulsets from storage state
	runningStatefulsetApps := ssm.registryManager.LoadRunningApplies()
	for _, application := range runningStatefulsetApps {
		for _, staefulsetData := range application.DBSchema.Resources.Statefulsets {
			staefulsetWatchListOptions := metaV1.ListOptions{
				LabelSelector: labels.SelectorFromSet(staefulsetData.Statefulset.Labels).String(),
			}
			go ssm.watchStatefulset(application.ctx, application.cancelFn, application.Log(), staefulsetData,
				staefulsetWatchListOptions, staefulsetData.Statefulset.Namespace, staefulsetData.ProgressDeadlineSeconds)
		}
	}
	ssm.watchStatefulsets(ctx)

}

func (ssm *StatefulsetManager) watchStatefulsets(ctx context.Context) {
	statefulsetsList, _ := ssm.client.AppsV1().StatefulSets("").List(metaV1.ListOptions{})
	statefulsetWatchListOptions := metaV1.ListOptions{ResourceVersion: statefulsetsList.GetResourceVersion()}
	watcher, err := ssm.client.AppsV1().StatefulSets("").Watch(statefulsetWatchListOptions)
	if err != nil {
		log.WithError(err).WithField("list_option", statefulsetWatchListOptions.String()).Error("Could not start a watcher on statefulset")
		return
	}
	go func() {
		log.WithField("resource_version", statefulsetsList.GetResourceVersion()).Info("Statefulsets watcher was started")
		for {
			select {
			case event, watch := <-watcher.ResultChan():
				if !watch {
					log.WithField("list_options", statefulsetWatchListOptions.String()).Info("Statefulsets watcher was stopped. Reopen the channel")
					ssm.watchStatefulsets(ctx)
					return
				}
				statefulset, ok := event.Object.(*appsV1.StatefulSet)

				if !ok {
					log.WithField("object", event.Object).Warn("Failed to parse statefulset watcher data")
					continue
				}

				statefulsetName := GetApplicationName(statefulset.GetAnnotations(), statefulset.GetName())

				if event.Type == eventwatch.Modified || event.Type == eventwatch.Added || event.Type == eventwatch.Deleted {
					// handle modified event
					if event.Type == eventwatch.Deleted {
						ssm.registryManager.DeleteAppliedVersion(statefulset.GetName(), statefulset.GetNamespace(), "statefulset")
					} else {
						hash, _ := hashstructure.Hash(statefulset.Spec, nil)
						if !ssm.registryManager.UpdateAppliesVersionHistory(statefulset.GetName(), statefulset.GetNamespace(), "statefulset", hash) {
							continue
						}
					}
					appRegistry := ssm.registryManager.Get(statefulsetName, statefulset.GetNamespace())

					// extract annotation for progressDeadLine since statefulset don't have this feature
					progressDeadLine := GetProgressDeadlineApply(statefulset.GetAnnotations(), ssm.maxDeploymentTime)

					if appRegistry == nil {
						statefulsetStatus := common.DeploymentStatusRunning
						if event.Type == eventwatch.Deleted {
							statefulsetStatus = common.DeploymentStatusDeleted
						}

						appRegistry = ssm.registryManager.NewApplication(statefulsetName,
							statefulset.GetNamespace(),
							statefulset.GetAnnotations(),
							statefulsetStatus)
					}

					registryApply := appRegistry.AddStatefulset(
						statefulsetName,
						statefulset.GetNamespace(),
						statefulset.GetLabels(),
						statefulset.GetAnnotations(),
						*statefulset.Spec.Replicas,
						progressDeadLine)
					statefulsetWatchListOptions := metaV1.ListOptions{
						LabelSelector: labels.SelectorFromSet(statefulset.GetLabels()).String()}

					go ssm.watchStatefulset(
						appRegistry.ctx,
						appRegistry.cancelFn,
						appRegistry.Log(),
						registryApply,
						statefulsetWatchListOptions,
						statefulset.GetNamespace(),
						progressDeadLine)
				} else {
					log.WithFields(log.Fields{
						"event_type":  event.Type,
						"statefulset": statefulsetName,
					}).Info("Event type not supported")
				}
			case <-ctx.Done():
				log.Warn("Statefulset watcher was stopped. Got ctx done signal")
				watcher.Stop()
				return
			}
		}
	}()
}

//  watchStatefulset will watch a specific statefulset and its related resources (controller revision + pods)
func (ssm *StatefulsetManager) watchStatefulset(ctx context.Context, cancelFn context.CancelFunc, lg log.Entry, registryStatefulset *StatefulsetData, listOptions metaV1.ListOptions, namespace string, maxWatchTime int64) {

	statefulsetLog := lg.WithField("statefulset_name", registryStatefulset.GetName())

	statefulsetLog.Info("Starting Statefulset watcher")
	statefulsetLog.WithField("list_option", listOptions.String()).Debug("List option for statefulset filtering")

	watcher, err := ssm.client.AppsV1().StatefulSets(namespace).Watch(listOptions)
	if err != nil {
		statefulsetLog.WithError(err).Error("Could not start statefulset watcher")
		return
	}
	firstInit := true
	for {
		select {
		case event, watch := <-watcher.ResultChan():
			if !watch {
				statefulsetLog.Warn("Statefulset watcher was stopped. Channel was closed")
				cancelFn()
				return
			}
			statefulset, isOk := event.Object.(*appsV1.StatefulSet)
			if !isOk {
				statefulsetLog.WithField("object", event.Object).Warn("Failed to parse statefulset watcher data")
				continue
			}
			if firstInit {
				firstInit = false
				eventListOptions := metaV1.ListOptions{FieldSelector: labels.SelectorFromSet(map[string]string{
					"involvedObject.name": registryStatefulset.Statefulset.Name,
					"involvedObject.kind": "StatefulSet",
				}).String(),
					TimeoutSeconds: &maxWatchTime,
				}

				// Start watching on Events of statefulset
				ssm.watchEvents(ctx, *statefulsetLog, registryStatefulset, eventListOptions, namespace)

				// Use the Controller revision to find the pods with specific controller-revision-hash for the statefulset
				ssm.controllerRevManager.WatchControllerRevisionPodsRetry(ctx, *statefulsetLog,
					registryStatefulset,
					statefulset.ObjectMeta.Generation,
					statefulset.Spec.Selector.MatchLabels,
					"controller.kubernetes.io/hash",
					statefulset.ObjectMeta.Name,
					namespace,
					nil)

				// Start watching services of statefulset
				ssm.serviceManager.Watch <- WatchData{
					ListOptions:  metaV1.ListOptions{TimeoutSeconds: &maxWatchTime, LabelSelector: labels.SelectorFromSet(statefulset.Spec.Selector.MatchLabels).String()},
					RegistryData: registryStatefulset,
					Namespace:    statefulset.Namespace,
					Ctx:          ctx,
					LogEntry:     *statefulsetLog,
				}
			}
			registryStatefulset.UpdateApplyStatus(statefulset.Status)
		case <-ctx.Done():
			statefulsetLog.Debug("Statefulset watcher was stopped. Got ctx done signal")
			watcher.Stop()
			return
		}
	}
}

// watchEvents will watch for events relate d to the Statefulset Resources
func (ssm *StatefulsetManager) watchEvents(ctx context.Context, lg log.Entry, registryStatefulset *StatefulsetData, listOptions metaV1.ListOptions, namespace string) {

	lg.Info("Started the event watcher on statefulset events")
	watchData := WatchEvents{
		ListOptions: listOptions,
		Namespace:   namespace,
		Ctx:         ctx,
		LogEntry:    lg,
	}

	eventChan := ssm.eventManager.Watch(watchData)
	go func() {
		for {
			select {
			case event := <-eventChan:
				registryStatefulset.UpdateStatefulsetEvents(event)
			case <-ctx.Done():
				lg.Info("Stopped the event watcher on statefulset events")
				return
			}
		}
	}()
}
