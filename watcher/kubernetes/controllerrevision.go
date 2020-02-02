package kuberneteswatcher

import (
	"context"
	"errors"
	"fmt"
	"time"

	backoff "github.com/cenkalti/backoff/v4"
	log "github.com/sirupsen/logrus"
	appsV1 "k8s.io/api/apps/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
)

//BackoffParams parameters
type BackoffParams struct {
	InitialInterval time.Duration
	Multiplier      float64
	MaxElapsedTime  time.Duration
}

//NewBackOffParams Parameters init
func NewBackOffParams() *BackoffParams {
	return &BackoffParams{
		InitialInterval: 0,
		Multiplier:      1.5,
		MaxElapsedTime:  60 * time.Second,
	}
}

//ControllerRevision Main interface
type ControllerRevision interface {
	WatchControllerRevisionPods(ctx context.Context, registryData RegistryData, resourceGeneration int64, controllerRevisionlabels map[string]string, controllerRevisionHashlabelKey string, controllerRevisionPodLabelValuePerfix string, namespace string) error
	WatchControllerRevisionPodsRetry(ctx context.Context, registryData RegistryData, resourceGeneration int64, controllerRevisionlabels map[string]string, controllerRevisionHashlabelKey string, controllerRevisionPodLabelValuePerfix string, namespace string, backOffParams *BackoffParams) error
}

//ControllerRevisionManager Manager to interfact with Kubernetes kind
type ControllerRevisionManager struct {
	// Kubernetes client
	client kubernetes.Interface

	// to find pods related
	podsManager *PodsManager
}

// NewControllerRevisionManager create new instance of controllerRevision manager
func NewControllerRevisionManager(client kubernetes.Interface, podsManager *PodsManager) *ControllerRevisionManager {
	return &ControllerRevisionManager{
		client:      client,
		podsManager: podsManager,
	}
}

// WatchControllerRevisionPodsRetry perform exponential backoff retry on WatchControllerRevisionPods
func (cr *ControllerRevisionManager) WatchControllerRevisionPodsRetry(ctx context.Context, registryData RegistryData, resourceGeneration int64, controllerRevisionlabels map[string]string, controllerRevisionHashlabelKey string, controllerRevisionPodLabelValuePerfix string, namespace string, backOff *BackoffParams) error {
	defaultParams := NewBackOffParams()
	if backOff != nil {
		defaultParams = backOff
	}
	b := backoff.NewExponentialBackOff()
	b.InitialInterval = defaultParams.InitialInterval
	b.MaxElapsedTime = defaultParams.MaxElapsedTime
	b.Multiplier = defaultParams.Multiplier
	ticker := backoff.NewTicker(b)
	var err error
	for range ticker.C {
		if err = cr.WatchControllerRevisionPods(ctx, registryData, resourceGeneration, controllerRevisionlabels, controllerRevisionHashlabelKey, controllerRevisionPodLabelValuePerfix, namespace); err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).Warn("retrying backoff")
			continue
		}
		ticker.Stop()
		break
	}
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Stopping retrying backoff Fail")
		return err

	}
	return err
}

// WatchControllerRevisionPods finds the correct pods to watch
// 1. search a controllerrevision resource that is related to (statefulset or daemonset) using the version id and labels.
// 2. once found, extract the controller-revision-hash value and look for pods with this annotation
// 3. watch those pods.
func (cr *ControllerRevisionManager) WatchControllerRevisionPods(ctx context.Context, registryData RegistryData, resourceGeneration int64, controllerRevisionlabels map[string]string, controllerRevisionHashlabelKey string, controllerRevisionPodLabelValuePerfix string, namespace string) error {
	registryDataName := registryData.GetName()
	// find controller revision that fits the resource version`
	revisions, err := cr.client.AppsV1().ControllerRevisions(namespace).List(metaV1.ListOptions{
		LabelSelector: labels.SelectorFromSet(controllerRevisionlabels).String()})

	if err != nil {
		log.WithFields(log.Fields{
			fmt.Sprintf("%T", registryData): registryDataName,
			"namespace":                     namespace,
			"revision":                      resourceGeneration,
			"error":                         err,
		}).Error("Cannot list revisions")
		return errors.New("Cannot list revisions")
	}

	// Get the revision hash inside controller revision
	for _, revision := range revisions.Items {
		if revision.Revision == resourceGeneration {
			log.WithField("controller_revision_hash_label_key", controllerRevisionHashlabelKey).Info(
				"Searching for controllerRevisionHash from the controllerRevisionHashlabelKey")
			controllerRevisionHash, valExist := revision.ObjectMeta.Labels[controllerRevisionHashlabelKey]
			if !valExist {
				log.WithFields(log.Fields{
					fmt.Sprintf("%T", registryData):      registryDataName,
					"namespace":                          namespace,
					"controller_revision_hash_label_key": controllerRevisionHashlabelKey,
					"revision":                           resourceGeneration,
				}).Warn("Cannot find controllerRevision label inside ControllerRevision kind. cannot start watch on pods")
				return errors.New("Cannot find controllerRevisionHashLabelKey lables, inside ControllerRevision. cannot start watch on pods")
			}
			log.WithFields(log.Fields{
				fmt.Sprintf("%T", registryData):      registryDataName,
				"namespace":                          namespace,
				"revision":                           resourceGeneration,
				"controller_revision_hash_label_key": controllerRevisionHashlabelKey,
				"controllerRevisionHash":             controllerRevisionHash,
			}).Debug("ControllerRevision found controller-revision-hash calling Pods Manager")

			log.WithField("controllerRevisionPodLabelValuePerfix", controllerRevisionPodLabelValuePerfix).Debug(
				"Going to check for ConrollerRevisionPodLabel Value with prefix")

			controllerRevisionPodLabelValue := controllerRevisionHash
			if controllerRevisionPodLabelValuePerfix != "" {
				controllerRevisionPodLabelValue = fmt.Sprintf("%s-%s", registryData.GetName(), controllerRevisionHash)
			}

			log.WithFields(log.Fields{
				"controller_revision_pod_label_key":   appsV1.ControllerRevisionHashLabelKey,
				"controller_revision_pod_label_value": controllerRevisionPodLabelValue}).Debug("Going to watch pods with the following fields")

			// Start watching pods with the specific appsV1.ControllerRevisionHashLabelKey
			podLabelSelector := map[string]string{appsV1.ControllerRevisionHashLabelKey: controllerRevisionPodLabelValue}
			podListOptions := metaV1.ListOptions{LabelSelector: labels.SelectorFromSet(podLabelSelector).String()}
			cr.podsManager.Watch <- WatchData{
				ListOptions:  podListOptions,
				RegistryData: registryData,
				Namespace:    namespace,
				Ctx:          ctx,
			}
			return nil
		}
	}
	log.WithFields(log.Fields{
		fmt.Sprintf("%T", registryData):      registryDataName,
		"namespace":                          namespace,
		"revision":                           resourceGeneration,
		"controller_revision_hash_label_key": controllerRevisionHashlabelKey,
	}).Error("Cannot find resourceVersion in ControllerRevision. cannot start watch on pods")
	return errors.New("Cannot find resourceVersion in ControllerRevision. cannot start watch on pods")
}
