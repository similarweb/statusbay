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
}

// NewDeploymentManager create new deployment instance
func NewDeploymentManager(kubernetesClientset kubernetes.Interface, eventManager *EventsManager, registryManager *RegistryManager, replicaset *ReplicaSetManager, serviceManager *ServiceManager, maxDeploymentTime time.Duration) *DeploymentManager {
	return &DeploymentManager{
		client:            kubernetesClientset,
		eventManager:      eventManager,
		registryManager:   registryManager,
		replicaset:        replicaset,
		serviceManager:    serviceManager,
		maxDeploymentTime: int64(maxDeploymentTime.Seconds()),
	}
}

// Serve will start listening on deployment request
func (dm *DeploymentManager) Serve(ctx context.Context, wg *sync.WaitGroup) {

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Warn("Deployment Manager has been shut down")
				wg.Done()
				return
			}
		}
	}()

	//Continue running deployments from storage state
	runningDeploymentApplication := dm.registryManager.LoadRunningApplies()
	for _, application := range runningDeploymentApplication {
		for _, deploymentData := range application.DBSchema.Resources.Deployments {
			deploymentWatchListOptions := metaV1.ListOptions{LabelSelector: labels.SelectorFromSet(deploymentData.Deployment.Labels).String()}
			go dm.watchDeployment(application.ctx, application.cancelFn, application.Log(), deploymentData, deploymentWatchListOptions, deploymentData.Deployment.Namespace, deploymentData.ProgressDeadlineSeconds)
		}
	}

	dm.watchDeployments(ctx)

}

// watchDeployments start watch on all Kubernetes deployments
func (dm *DeploymentManager) watchDeployments(ctx context.Context) {

	deploymentList, _ := dm.client.AppsV1().Deployments("").List(metaV1.ListOptions{})
	deploymentWatchListOptions := metaV1.ListOptions{ResourceVersion: deploymentList.GetResourceVersion()}
	watcher, err := dm.client.AppsV1().Deployments("").Watch(deploymentWatchListOptions)

	if err != nil {
		log.WithError(err).WithField("list_option", deploymentWatchListOptions.String()).Error("Could not start a watcher on deployment")

		return
	}

	go func() {
		log.WithField("resource_version", deploymentList.GetResourceVersion()).Info("Deployments watcher was started")
		for {
			select {
			case event, watch := <-watcher.ResultChan():

				if !watch {
					log.WithField("list_options", deploymentWatchListOptions.String()).Info("Deployments watcher was stopped. Reopen the channel")
					dm.watchDeployments(ctx)
					return
				}

				deployment, ok := event.Object.(*appsV1.Deployment)
				if !ok {
					log.WithField("object", event.Object).Warn("Failed to parse deployment watcher data")
					continue
				}

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
					}

					applicationRegistry := dm.registryManager.NewApplyEvent(apply)
					if applicationRegistry == nil {
						continue
					}
					registryDeployment := dm.AddNewDeployment(apply, applicationRegistry, *deployment.Spec.Replicas)

					deploymentWatchListOptions := metaV1.ListOptions{LabelSelector: labels.SelectorFromSet(deployment.GetLabels()).String()}

					maxWatchTime := dm.maxDeploymentTime

					go dm.watchDeployment(
						applicationRegistry.ctx,
						applicationRegistry.cancelFn,
						applicationRegistry.Log(),
						registryDeployment,
						deploymentWatchListOptions,
						deployment.GetNamespace(),
						maxWatchTime)

				} else {
					log.WithFields(log.Fields{
						"event_type":      event.Type,
						"deployment_name": deploymentName,
					}).Info("Event type not supported")
				}

			case <-ctx.Done():
				log.Warn("Deployment watch was stopped. Got ctx done signal")
				watcher.Stop()
				return

			}

		}

	}()

}

//watchDeployment will watch on one running deployment
func (dm *DeploymentManager) watchDeployment(ctx context.Context, cancelFn context.CancelFunc, lg log.Entry, registryDeployment *DeploymentData, listOptions metaV1.ListOptions, namespace string, maxWatchTime int64) {

	deploymentLog := lg.WithField("deployment_name", registryDeployment.GetName())
	deploymentLog.Info("Starting watch on deployment")
	deploymentLog.WithField("list_option", listOptions.String()).Debug("List option for deployment filtering")

	watcher, err := dm.client.AppsV1().Deployments(namespace).Watch(listOptions)
	if err != nil {
		deploymentLog.Error("Could not start watch on deployment")
		return
	}

	firstInit := true

	for {
		select {
		case event, watch := <-watcher.ResultChan():
			if !watch {
				deploymentLog.Warn("Deployment watcher was stopped. Channel was closed")
				cancelFn()
				return
			}

			deployment, isOk := event.Object.(*appsV1.Deployment)
			if !isOk {
				deploymentLog.WithField("object", event.Object).Warn("Failed to parse deployment watcher data")
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
			deploymentLog.Debug("Deployment watcher was stopped. Got ctx done signal")
			watcher.Stop()
			return

		}
	}

}

// watchEvents will start watch on deployment event messages changes
func (dm *DeploymentManager) watchEvents(ctx context.Context, lg log.Entry, registryDeployment *DeploymentData, listOptions metaV1.ListOptions, namespace string) {
	lg.Info("Started the event watcher on deployment events")

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
				lg.Info("Stop watch on deployment events")
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
			Metrics:      GetMetricsDataFromAnnotations(data.Annotations),
			Alerts:       GetAlertsDataFromAnnotations(data.Annotations),
			DesiredState: desiredState,
		},
		Pods:                    make(map[string]DeploymenPod, 0),
		Replicaset:              make(map[string]Replicaset, 0),
		ProgressDeadlineSeconds: GetProgressDeadlineApply(data.Annotations, dm.maxDeploymentTime),
	}
	applicationRegistry.DBSchema.Resources.Deployments[data.ApplyName] = dd

	log.Info("Deployment was associated to the application")

	return dd

}
