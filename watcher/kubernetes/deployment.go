package kuberneteswatcher

import (
	"context"
	"statusbay/serverutil"
	"time"

	"statusbay/watcher/kubernetes/common"

	"github.com/mitchellh/hashstructure"
	log "github.com/sirupsen/logrus"

	appsV1 "k8s.io/api/apps/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	eventwatch "k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
)

// DeploymentStatusDescription are the various descriptions of the states a deployment can be in.
type DeploymentStatusDescription string

const (

	// DeploymentStatusDescriptionRunning running deployment
	DeploymentStatusDescriptionRunning DeploymentStatusDescription = "Deployment is running"

	// DeploymentStatusDescriptionSuccessful successfully deployment
	DeploymentStatusDescriptionSuccessful DeploymentStatusDescription = "Deployment completed successfully"

	// DeploymentStatusDescriptionProgressDeadline progress deadline ended
	DeploymentStatusDescriptionProgressDeadline DeploymentStatusDescription = "Failed due to progress deadline"
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
func (dm *DeploymentManager) Serve() serverutil.StopFunc {

	ctx, cancelFn := context.WithCancel(context.Background())
	stopped := make(chan bool)
	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Warn("Deployment Manager has been shut down")
				stopped <- true
				return
			}
		}
	}()

	//Continue running deployments from storage state
	runningDeploymentApplication := dm.registryManager.LoadRunningApplies()
	for _, application := range runningDeploymentApplication {
		for _, deploymentData := range application.DBSchema.Resources.Deployments {
			deploymentWatchListOptions := metaV1.ListOptions{LabelSelector: labels.SelectorFromSet(deploymentData.Deployment.Labels).String()}
			go dm.watchDeployment(application.ctx, application.cancelFn, deploymentData, deploymentWatchListOptions, deploymentData.Deployment.Namespace, deploymentData.ProgressDeadlineSeconds)
		}
	}

	dm.watchDeployments(ctx)
	return func() {
		cancelFn()
		<-stopped
	}
}

// watchDeployments start watch on all Kubernetes deployments
func (dm *DeploymentManager) watchDeployments(ctx context.Context) {

	deploymentList, _ := dm.client.AppsV1().Deployments("").List(metaV1.ListOptions{})
	deploymentWatchListOptions := metaV1.ListOptions{ResourceVersion: deploymentList.GetResourceVersion()}
	watcher, err := dm.client.AppsV1().Deployments("").Watch(deploymentWatchListOptions)

	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"list_option": deploymentWatchListOptions.String(),
		}).Error("Could not start watch on deployment")
		return
	}

	go func() {
		log.WithField("resource_version", deploymentList.GetResourceVersion()).Info("Deployments watch was started")
		for {
			select {
			case event, watch := <-watcher.ResultChan():

				if !watch {
					log.WithFields(log.Fields{
						"list_options": deploymentWatchListOptions.String(),
					}).Info("Deployments watch was stopped. Channel was closed")
					dm.watchDeployments(ctx)
					return
				}

				deployment, ok := event.Object.(*appsV1.Deployment)
				if !ok {
					log.WithField("object", event.Object).Warn("Failed to parse deployment watch data")
					continue
				}

				deploymentName := GetApplicationName(deployment.GetAnnotations(), deployment.GetName())

				if event.Type == eventwatch.Modified || event.Type == eventwatch.Added || event.Type == eventwatch.Deleted {

					if event.Type == eventwatch.Deleted {
						dm.registryManager.DeleteAppliedVersion(deployment.GetName(), deployment.GetNamespace())
					} else {
						hash, _ := hashstructure.Hash(deployment.Spec, nil)
						if !dm.registryManager.UpdateAppliesVersionHistory(deployment.GetName(), deployment.GetNamespace(), hash) {
							continue
						}
					}

					applicationRegistry := dm.registryManager.Get(deploymentName, deployment.GetNamespace())

					// extract annotation for progressDeadLine since Daemonset don't have hat feature.
					progressDeadLine := GetProgressDeadlineApply(deployment.GetAnnotations(), dm.maxDeploymentTime)

					if applicationRegistry == nil {

						deploymentStatus := common.DeploymentStatusRunning
						if event.Type == eventwatch.Deleted {
							deploymentStatus = common.DeploymentStatusDeleted
						}
						applicationRegistry = dm.registryManager.NewApplication(deploymentName,
							deployment.GetNamespace(),
							deployment.GetAnnotations(),
							deploymentStatus)
					}

					registryDeployment := applicationRegistry.AddDeployment(deployment.GetName(),
						deployment.GetNamespace(),
						deployment.GetLabels(),
						deployment.GetAnnotations(),
						*deployment.Spec.Replicas,
						progressDeadLine)
					deploymentWatchListOptions := metaV1.ListOptions{LabelSelector: labels.SelectorFromSet(deployment.GetLabels()).String()}

					maxWatchTime := dm.maxDeploymentTime

					go dm.watchDeployment(applicationRegistry.ctx,
						applicationRegistry.cancelFn,
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

// watchDeployment will watch on one running deployment
func (dm *DeploymentManager) watchDeployment(ctx context.Context, cancelFn context.CancelFunc, registryDeployment *DeploymentData, listOptions metaV1.ListOptions, namespace string, maxWatchTime int64) {

	log.WithFields(log.Fields{
		"deployment": registryDeployment.Deployment.Name,
		"namespace":  namespace,
	}).Info("Starting watch on deployment")

	log.WithFields(log.Fields{
		"deployment":  registryDeployment.Deployment.Name,
		"list_option": listOptions.String(),
		"namespace":   namespace,
	}).Debug("Deployment watch list option")

	watcher, err := dm.client.AppsV1().Deployments(namespace).Watch(listOptions)
	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"deployment": registryDeployment.Deployment.Name,
			"namespace":  namespace,
		}).Error("Could not start watch on deployment")
		return
	}

	firstInit := true

	for {
		select {
		case event, watch := <-watcher.ResultChan():
			if !watch {
				log.WithFields(log.Fields{
					"deployment": registryDeployment.Deployment.Name,
					"namespace":  namespace,
				}).Warn("Deployment watch was stopped. Channel was closed")
				cancelFn()
				return
			}

			deployment, ok := event.Object.(*appsV1.Deployment)
			if !ok {
				log.WithFields(log.Fields{
					"deployment": registryDeployment.Deployment.Name,
					"namespace":  namespace,
				}).WithField("object", event.Object).Warn("Failed to parse deployment watch data")
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
				dm.watchEvents(ctx, registryDeployment, eventListOptions, namespace)

				//Starting replicaset watch
				dm.replicaset.Watch <- WatchReplica{
					DesiredReplicas: *deployment.Spec.Replicas,
					ListOptions:     metaV1.ListOptions{TimeoutSeconds: &maxWatchTime, LabelSelector: labels.SelectorFromSet(deployment.Spec.Selector.MatchLabels).String()},
					Registry:        registryDeployment,
					Namespace:       deployment.Namespace,
					Ctx:             ctx,
				}

				dm.serviceManager.Watch <- WatchData{
					ListOptions:  metaV1.ListOptions{TimeoutSeconds: &maxWatchTime, LabelSelector: labels.SelectorFromSet(deployment.Spec.Selector.MatchLabels).String()},
					RegistryData: registryDeployment,
					Namespace:    deployment.Namespace,
					Ctx:          ctx,
				}
			}

			registryDeployment.UpdateDeploymentStatus(deployment.Status)

		case <-ctx.Done():
			log.WithFields(log.Fields{
				"selector":  listOptions.String(),
				"namespace": namespace,
			}).Debug("Deployment watch was stopped. Got ctx done signal")
			watcher.Stop()
			return

		}
	}

}

// watchEvents will start watch on deployment event messages changes
func (dm *DeploymentManager) watchEvents(ctx context.Context, registryDeployment *DeploymentData, listOptions metaV1.ListOptions, namespace string) {

	log.WithFields(log.Fields{
		"deployment": registryDeployment.Deployment.Name,
		"namespace":  namespace,
	}).Info("Start watch on deployment events")

	watchData := WatchEvents{
		ListOptions: listOptions,
		Namespace:   namespace,
		Ctx:         ctx,
	}
	eventChan := dm.eventManager.Watch(watchData)
	go func() {

		for {
			select {
			case event := <-eventChan:
				registryDeployment.UpdateDeploymentEvents(event)
			case <-ctx.Done():
				log.WithFields(log.Fields{
					"deployment": registryDeployment.Deployment.Name,
					"namespace":  namespace,
				}).Info("Stop watch on deployment events")
				return
			}
		}
	}()
}
