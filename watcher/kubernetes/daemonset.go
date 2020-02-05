//reference: https://gitlab.similarweb.io/elad.kaplan/statusier-open-source/blob/test_replicaset/watcher/kubernetes/daemonset.go
package kuberneteswatcher

import (
	"context"
	"statusbay/serverutil"
	"statusbay/watcher/kubernetes/common"
	"strconv"
	"time"

	"github.com/mitchellh/hashstructure"
	log "github.com/sirupsen/logrus"
	appsV1 "k8s.io/api/apps/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	eventwatch "k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
)

type DaemonsetManager struct {
	// Kubernetes client
	client kubernetes.Interface

	// Event manager will be owner to start watch on deployment events
	eventManager *EventsManager

	// Registry manager will be owner to manage the running / new deployment
	registryManager *RegistryManager

	// Will triggered when deployment watch started
	serviceManager *ServiceManager

	// will be trigged when pod deployment watch start
	podsManager *PodsManager

	//
	controllerRevManager ControllerRevision
	// Max watch time
	maxDeploymentTime int64
}

//NewDaemonsetManager  create new instance to manage damonset related things
func NewDaemonsetManager(k8sClient kubernetes.Interface, eventManager *EventsManager, registryManager *RegistryManager, serviceManager *ServiceManager, podsManager *PodsManager, controllerRevisionManager ControllerRevision, maxDeploymentTime time.Duration) *DaemonsetManager {
	return &DaemonsetManager{
		client:               k8sClient,
		eventManager:         eventManager,
		registryManager:      registryManager,
		serviceManager:       serviceManager,
		podsManager:          podsManager,
		controllerRevManager: controllerRevisionManager,
		maxDeploymentTime:    int64(maxDeploymentTime.Seconds()),
	}
}

func (dsm *DaemonsetManager) Serve() serverutil.StopFunc {
	ctx, cancelFn := context.WithCancel(context.Background())
	stopped := make(chan bool)
	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Warn("Daemonset Manager has been shut down")
				stopped <- true
				return
			}
		}
	}()
	// continue running daemonsets from storage state
	runningDaemonsetsApps := dsm.registryManager.LoadRunningApplies()
	for _, application := range runningDaemonsetsApps {
		for _, daemonsetData := range application.DBSchema.Resources.Daemonsets {
			daemonsetWatchListOptions := metaV1.ListOptions{
				LabelSelector: labels.SelectorFromSet(daemonsetData.Metadata.Labels).String(),
			}
			go dsm.watchDaemonset(
				application.ctx,
				application.cancelFn,
				daemonsetData,
				daemonsetWatchListOptions,
				daemonsetData.Metadata.Namespace,
				daemonsetData.ProgressDeadlineSeconds,
			)
		}
	}
	dsm.watchDaemonsets(ctx)
	return func() {
		cancelFn()
		<-stopped
	}

}

func (dsm *DaemonsetManager) watchDaemonsets(ctx context.Context) {
	daemonsetsList, _ := dsm.client.AppsV1().DaemonSets("").List(metaV1.ListOptions{})
	daemonsetWatchListOptions := metaV1.ListOptions{ResourceVersion: daemonsetsList.GetResourceVersion()}
	watcher, err := dsm.client.AppsV1().DaemonSets("").Watch(daemonsetWatchListOptions)
	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"list_option": daemonsetWatchListOptions.String(),
		}).Error("Could not start watch on daemonset")
		return
	}
	go func() {
		log.WithField("resource_version", daemonsetsList.GetResourceVersion()).Info("Daemonsets watch was started")
		for {
			select {
			case event, watch := <-watcher.ResultChan():
				if !watch {
					log.WithFields(log.Fields{
						"list_options": daemonsetWatchListOptions.String(),
					}).Info("Daemonsets watch was stopped. Channel was closed")
					dsm.watchDaemonsets(ctx)
					return
				}
				daemonset, ok := event.Object.(*appsV1.DaemonSet)
				if !ok {
					log.WithField("object", event.Object).Warn("Failed to parse daemonset watch data")
					continue
				}
				daemonsetName := daemonset.GetName()
				applicationName := GetMetadata(daemonset.GetAnnotations(), "statusbay.io/application-name")
				if applicationName != "" {
					daemonsetName = applicationName
				}
				// extract annotation for progressDeadLine since Daemonset don't have hat feature.
				progressDeadLineAnnotations := GetMetadata(daemonset.GetAnnotations(), "statusbay.io/progress-deadline-seconds")
				progressDeadLine, err := strconv.ParseInt(progressDeadLineAnnotations, 10, 64)
				if err != nil {
					progressDeadLine = dsm.maxDeploymentTime
				}
				if event.Type == eventwatch.Modified ||
					event.Type == eventwatch.Added ||
					event.Type == eventwatch.Deleted {
					// handle modified update
					if event.Type == eventwatch.Deleted {
						dsm.registryManager.DeleteAppliedVersion(daemonset.GetName(), daemonset.GetNamespace())
					} else {
						hash, _ := hashstructure.Hash(daemonset.Spec, nil)
						if !dsm.registryManager.UpdateAppliesVersionHistory(
							daemonset.GetName(),
							daemonset.GetNamespace(),
							hash) {
							continue
						}
					}
					appRegistry := dsm.registryManager.Get(daemonsetName, daemonset.GetNamespace())
					if appRegistry == nil {
						daemonsetStatus := common.DeploymentStatusRunning
						if event.Type == eventwatch.Deleted {
							daemonsetStatus = common.DeploymentStatusDeleted
						}
						appRegistry = dsm.registryManager.NewApplication(daemonsetName,
							daemonset.GetName(),
							daemonset.GetNamespace(),
							"cluster-name",
							daemonset.GetAnnotations(),
							daemonsetStatus)

					}
					registryApply := appRegistry.AddDaemonset(
						daemonset.GetName(),
						daemonset.GetNamespace(),
						daemonset.GetLabels(),
						daemonset.GetAnnotations(),
						daemonset.Status.DesiredNumberScheduled,
						progressDeadLine)
					daemonsetWatchListOptions := metaV1.ListOptions{
						LabelSelector: labels.SelectorFromSet(daemonset.GetLabels()).String()}
					maxWatchTime := dsm.maxDeploymentTime

					if progressDeadLine > dsm.maxDeploymentTime {
						maxWatchTime = progressDeadLine
					}
					go dsm.watchDaemonset(
						appRegistry.ctx,
						appRegistry.cancelFn,
						registryApply,
						daemonsetWatchListOptions,
						daemonset.GetNamespace(),
						maxWatchTime)
				} else {
					log.WithFields(log.Fields{
						"event_type": event.Type,
						"deamonset":  daemonsetName,
					}).Info("Event type not supported")
				}
			case <-ctx.Done():
				log.Warn("Daemonset watch was stopped. Got ctx done signal")
				watcher.Stop()
				return
			}
		}
	}()
}

// watchDaemonset will watch a specific daemonset and its related resources (controller revision + pods)
func (dsm *DaemonsetManager) watchDaemonset(ctx context.Context, cancelFn context.CancelFunc, daemonsetData *DaemonsetData, listOptions metaV1.ListOptions, namespace string, maxWatchTime int64) {
	log.WithFields(log.Fields{
		"daemonset": daemonsetData.GetName(),
		"namespace": namespace,
	}).Info("Starting watch on Daemonset")

	log.WithFields(log.Fields{
		"daemonset":   daemonsetData.GetName(),
		"list_option": listOptions.String(),
		"namespace":   namespace,
	}).Debug("Daemonset watch list option")
	watcher, err := dsm.client.AppsV1().DaemonSets(namespace).Watch(listOptions)
	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"daemonset": daemonsetData.GetName(),
			"namesapce": namespace,
		}).Error("Could not start watch on daemonset")
		return
	}
	firstInit := true
	for {
		select {
		case event, watch := <-watcher.ResultChan():
			if !watch {
				log.WithFields(log.Fields{
					"daemonset": daemonsetData.GetName(),
					"namespace": namespace,
				}).Warn("Daemonset watch was stopped. Channel was closed")
				cancelFn()
				return
			}
			daemonset, isOk := event.Object.(*appsV1.DaemonSet)
			if !isOk {
				log.WithFields(log.Fields{
					"daemonset": daemonsetData.GetName(),
					"namespace": namespace,
				}).WithField("object", event.Object).Warn("Failed to parse daemonset watch data")
				continue
			}
			if firstInit {
				firstInit = false
				eventListOptions := metaV1.ListOptions{FieldSelector: labels.SelectorFromSet(map[string]string{
					"involvedObject.name": daemonset.GetName(),
					"involvedObject.kind": "DaemonSet",
				}).String(),
					TimeoutSeconds: &maxWatchTime,
				}
				dsm.watchEvents(ctx, daemonsetData, eventListOptions, namespace)
				// start pods watch
				dsm.watchPods(ctx, daemonsetData, daemonset, namespace)
				// start service watch
				dsm.serviceManager.Watch <- WatchData{
					ListOptions:  metaV1.ListOptions{TimeoutSeconds: &maxWatchTime, LabelSelector: labels.SelectorFromSet(daemonset.Spec.Selector.MatchLabels).String()},
					RegistryData: daemonsetData,
					Namespace:    daemonset.Namespace,
					Ctx:          ctx,
				}
			}
			daemonsetData.UpdateApplyStatus(daemonset.Status)
		case <-ctx.Done():
			log.WithFields(log.Fields{
				"selector":  listOptions.String(),
				"namespace": namespace,
			}).Debug("Daemonset watch was stopped. Got ctx done signal")
			watcher.Stop()
			return
		}
	}
}

// watchPods will trigger a controller revision manager to watch for the related pods
func (dsm *DaemonsetManager) watchPods(ctx context.Context, daemonsetData *DaemonsetData, daemonset *appsV1.DaemonSet, namespace string) {
	resourceGeneration := daemonset.ObjectMeta.Generation
	revisionLabels := map[string]string{"name": daemonsetData.GetName()}
	dsm.controllerRevManager.WatchControllerRevisionPodsRetry(ctx, daemonsetData, resourceGeneration, revisionLabels, namespace, nil)
}

// watchEvents will watch for events related to the Daemonset Resource
func (dsm *DaemonsetManager) watchEvents(ctx context.Context, daemonsetData *DaemonsetData, listOptions metaV1.ListOptions, namespace string) {
	log.WithFields(log.Fields{
		"daemonset": daemonsetData.GetName(),
		"namespace": namespace,
	}).Info("Start watch on daemonset events")
	watchData := WatchEvents{
		ListOptions: listOptions,
		Namespace:   namespace,
		Ctx:         ctx,
	}
	eventChan := dsm.eventManager.Watch(watchData)
	go func() {
		for {
			select {
			case event := <-eventChan:
				daemonsetData.UpdateDaemonsetEvents(event)
			case <-ctx.Done():
				log.WithFields(log.Fields{
					"daemonset": daemonsetData.GetName(),
					"namespace": namespace,
				}).Info("Stop watch on daemonset events")
				return
			}
		}
	}()
}
