package kuberneteswatcher

import (
	"context"
	"fmt"
	"statusbay/watcher/kubernetes/common"
	"sync"
	"time"

	"github.com/mitchellh/hashstructure"
	log "github.com/sirupsen/logrus"

	appsV1 "k8s.io/api/apps/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
)

// DeploymentManager defined deployment struct
type DeploymentManager struct {

	// Kubernetes client
	client kubernetes.Interface

	// Event manager will be owner to start watch on deployment events
	eventManager *EventsManager

	// Registry manager will be owner to manage the running / new deployment
	registryManager *RegistryManager

	// Will triggered when deployment watch started
	replicaset *ReplicaSetManager

	// Will triggered when deployment watch started
	serviceManager *ServiceManager

	// Max watch time
	maxDeploymentTime int64

	// Initial Running Applies to load on start
	initialRunningApplies []*RegistryRow
}

// NewDeploymentManager create new deployment instance
func NewDeploymentManager(kubernetesClientset kubernetes.Interface, eventManager *EventsManager, registryManager *RegistryManager, replicaset *ReplicaSetManager, serviceManager *ServiceManager, runningApplies []*RegistryRow, maxDeploymentTime time.Duration) *DeploymentManager {
	return &DeploymentManager{
		client:                kubernetesClientset,
		eventManager:          eventManager,
		registryManager:       registryManager,
		replicaset:            replicaset,
		serviceManager:        serviceManager,
		maxDeploymentTime:     int64(maxDeploymentTime.Seconds()),
		initialRunningApplies: runningApplies,
	}
}

// Serve will start listening on deployment request
func (dm *DeploymentManager) Serve(ctx context.Context, wg *sync.WaitGroup) {

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Warn("deployment manager has been shut down")
				wg.Done()
				return
			}
		}
	}()

	//Continue running deployments from storage state
	runningDeploymentApplication := dm.initialRunningApplies
	log.WithField("running_apps", len(runningDeploymentApplication)).Debug("loaded running applications in deployment manager")
	for _, application := range runningDeploymentApplication {
		app := application
		for _, deploymentData := range application.DBSchema.Resources.Deployments {
			depData := deploymentData
			deploymentWatchListOptions := metaV1.ListOptions{LabelSelector: labels.SelectorFromSet(deploymentData.Deployment.Labels).String()}
			app.Log().Logger.WithField("name", depData.GetName()).Debug("begining watching loaded running deployment")
			go func(app *RegistryRow, depData *DeploymentData, listOptions metaV1.ListOptions) {
				dm.watchDeployment(app.ctx, app.cancelFn, app.Log(), depData, listOptions, depData.Deployment.Namespace, depData.ProgressDeadlineSeconds)
			}(app, depData, deploymentWatchListOptions)
		}
	}
	// we dont need anymore that list
	dm.initialRunningApplies = nil
	dm.watchDeployments(ctx)
}

// watchDeployments start watch on all Kubernetes deployments
func (dm *DeploymentManager) watchDeployments(ctx context.Context) {

	deploymentList, _ := dm.client.AppsV1().Deployments("").List(metaV1.ListOptions{})
	deploymentWatchListOptions := metaV1.ListOptions{ResourceVersion: deploymentList.GetResourceVersion()}
	watcher, err := dm.client.AppsV1().Deployments("").Watch(deploymentWatchListOptions)

	if err != nil {
		log.WithError(err).WithField("list_option", deploymentWatchListOptions.String()).Error("could not start deployments watcher")

		return
	}

	go func() {
		log.WithField("resource_version", deploymentList.GetResourceVersion()).Info("starting deployments watcher")
		for {
			select {
			case event, watch := <-watcher.ResultChan():

				if !watch {
					log.WithField("list_options", deploymentWatchListOptions.String()).Info("deployments watcher was stopped, reopening the channel")
					dm.watchDeployments(ctx)
					return
				}

				deployment, ok := event.Object.(*appsV1.Deployment)
				if !ok {
					log.WithField("object", event.Object).Warn("failed to parse deployment watcher data")
					continue
				}

				log.WithFields(log.Fields{
					"name":      deployment.GetName(),
					"namespace": deployment.GetNamespace(),
				}).Debug("deployment event detected")
				deploymentName := GetApplicationName(deployment.GetAnnotations(), deployment.GetName())

				if common.IsSupportedEventType(event.Type) {

					hash, _ := hashstructure.Hash(deployment.Spec, nil)
					apply := ApplyEvent{
						Event:        fmt.Sprintf("%v", event.Type),
						ApplyName:    deploymentName,
						ResourceName: deployment.GetName(),
						Namespace:    deployment.GetNamespace(),
						Kind:         "deployment",
						Hash:         hash,
						Annotations:  deployment.GetAnnotations(),
						Labels:       deployment.GetLabels(),
					}

					applicationRegistry := dm.registryManager.NewApplyEvent(apply)
					if applicationRegistry == nil {
						continue
					}
					deploymentLog := applicationRegistry.Log()
					deploymentLog.WithField("event", event.Type).Info("adding deployment to apply registry")

					registryDeployment := dm.AddNewDeployment(apply, applicationRegistry, *deployment.Spec.Replicas)

					deploymentWatchListOptions := metaV1.ListOptions{LabelSelector: labels.SelectorFromSet(deployment.GetLabels()).String()}

					maxWatchTime := dm.maxDeploymentTime

					go dm.watchDeployment(
						applicationRegistry.ctx,
						applicationRegistry.cancelFn,
						deploymentLog,
						registryDeployment,
						deploymentWatchListOptions,
						deployment.GetNamespace(),
						maxWatchTime)

				} else {
					log.WithFields(log.Fields{
						"event_type":      event.Type,
						"deployment_name": deploymentName,
					}).Info("event type not supported")
				}

			case <-ctx.Done():
				log.Warn("deployment watch was stopped, got ctx done signal")
				watcher.Stop()
				return

			}

		}

	}()

}

//watchDeployment will watch on one running deployment
func (dm *DeploymentManager) watchDeployment(ctx context.Context, cancelFn context.CancelFunc, lg log.Entry, registryDeployment *DeploymentData, listOptions metaV1.ListOptions, namespace string, maxWatchTime int64) {

	deploymentLog := lg.WithField("deployment_name", registryDeployment.GetName())
	deploymentLog.Info("initializing deployments watcher")
	deploymentLog.WithField("list_option", listOptions.String()).Debug("list option for deployment filtering")

	watcher, err := dm.client.AppsV1().Deployments(namespace).Watch(listOptions)
	if err != nil {
		deploymentLog.WithError(err).Error("could not start deployments watcher")
		return
	}

	firstInit := true

	for {
		select {
		case event, watch := <-watcher.ResultChan():
			if !watch {
				deploymentLog.Warn("deployment watcher was stopped, channel was closed")
				cancelFn()
				return
			}

			deployment, isOk := event.Object.(*appsV1.Deployment)
			if !isOk {
				deploymentLog.WithField("object", event.Object).Warn("failed to parse deployment watcher data")
				continue
			}

			if firstInit {
				firstInit = false

				eventListOptions := metaV1.ListOptions{FieldSelector: labels.SelectorFromSet(map[string]string{
					"involvedObject.name": deployment.GetName(),
					"involvedObject.kind": "Deployment",
				}).String(),
					TimeoutSeconds: &maxWatchTime,
					// ResourceVersion: deployment.ResourceVersion,
				}
				dm.watchEvents(ctx, *deploymentLog, registryDeployment, eventListOptions, namespace)
				//Starting replicaset watch
				dm.replicaset.Watch <- WatchReplica{
					DesiredReplicas: *deployment.Spec.Replicas,
					ListOptions:     metaV1.ListOptions{TimeoutSeconds: &maxWatchTime, LabelSelector: labels.SelectorFromSet(deployment.Spec.Selector.MatchLabels).String()},
					Registry:        registryDeployment,
					Namespace:       deployment.Namespace,
					Ctx:             ctx,
					LogEntry:        *deploymentLog,
				}

				dm.serviceManager.Watch <- WatchData{
					ListOptions:  metaV1.ListOptions{TimeoutSeconds: &maxWatchTime, LabelSelector: labels.SelectorFromSet(deployment.Spec.Selector.MatchLabels).String()},
					RegistryData: registryDeployment,
					Namespace:    deployment.Namespace,
					Ctx:          ctx,
					LogEntry:     *deploymentLog,
				}
			}

			registryDeployment.UpdateDeploymentStatus(deployment.Status)

		case <-ctx.Done():
			deploymentLog.Debug("deployment watcher was stopped, got ctx done signal")
			watcher.Stop()
			return

		}
	}

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

// AddNewDeployment add new deployment under application
func (dm *DeploymentManager) AddNewDeployment(data ApplyEvent, applicationRegistry *RegistryRow, desiredState int32) *DeploymentData {

	log := applicationRegistry.Log()
	dd := &DeploymentData{
		Deployment: MetaData{
			Name:         data.ApplyName,
			Namespace:    data.Namespace,
			Annotations:  data.Annotations,
			Labels:       data.Labels,
			Metrics:      GetMetricsDataFromAnnotations(data.Annotations),
			Alerts:       GetAlertsDataFromAnnotations(data.Annotations),
			DesiredState: desiredState,
		},
		Pods:                    make(map[string]DeploymenPod, 0),
		Replicaset:              make(map[string]Replicaset, 0),
		ProgressDeadlineSeconds: GetProgressDeadlineApply(data.Annotations, dm.maxDeploymentTime),
	}
	applicationRegistry.DBSchema.Resources.Deployments[data.ResourceName] = dd

	log.Info("deployment was associated with application")

	return dd

}
