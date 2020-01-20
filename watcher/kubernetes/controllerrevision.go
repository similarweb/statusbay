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

type BackoffParams struct {
	InitialInterval time.Duration
	Multiplier      float64
	MaxElapsedTime  time.Duration
}

func NewBackOffParams() *BackoffParams {
	return &BackoffParams{
		InitialInterval: 0,
		Multiplier:      1.5,
		MaxElapsedTime:  60 * time.Second,
	}
}

type ControllerRevision interface {
	WatchControllerRevisionPods(ctx context.Context, registryData RegistryData, resourceGeneration int64, controllerRevisionlabels map[string]string, namespace string) error
	WatchControllerRevisionPodsRetry(ctx context.Context, registryData RegistryData, resourceGeneration int64, controllerRevisionlabels map[string]string, namespace string, backOffParams *BackoffParams)
}

type ControllerRevisionManager struct {
	// Kubernetes client
	client kubernetes.Interface

	// to find pods related
	podsManager *PodsManager
}

// NewControllerReisionManager create new instance of controllerRecision manager
func NewControllerReisionManager(client kubernetes.Interface, podsManager *PodsManager) *ControllerRevisionManager {
	return &ControllerRevisionManager{
		client:      client,
		podsManager: podsManager,
	}
}

// WatchControllerRevisionPodsRetry perform exponential backoff retry on WatchControllerRevisionPods
func (cr *ControllerRevisionManager) WatchControllerRevisionPodsRetry(ctx context.Context, registryData RegistryData, resourceGeneration int64, controllerRevisionlabels map[string]string, namespace string, backOff *BackoffParams) {
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
		if err = cr.WatchControllerRevisionPods(ctx, registryData, resourceGeneration, controllerRevisionlabels, namespace); err != nil {
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
		return
	}
}

// WatchControllerRevisionPods finds the correct pods to watch
// 1. search a controllerrevision resource that is related to (statefulset or daemonset) using the version id and labels.
// 2. once found, extract the controller-revision-hash value and look for pods with this annotation
// 3. watch those pods.
func (cr *ControllerRevisionManager) WatchControllerRevisionPods(ctx context.Context, registryData RegistryData, resourceGeneration int64, controllerRevisionlabels map[string]string, namespace string) error {
	// find controller revision that fits the resource version
	revisions, err := cr.client.AppsV1().ControllerRevisions(namespace).List(metaV1.ListOptions{
		LabelSelector: labels.SelectorFromSet(controllerRevisionlabels).String(),
	})
	if err != nil {
		log.WithFields(log.Fields{
			fmt.Sprintf("%T", registryData): registryData.GetName(),
			"namespace":                     namespace,
			"revision":                      resourceGeneration,
			"error":                         err,
		}).Error("Cannot list revisions")
		return errors.New("Cannot list revisions")
	}
	// find the revision hash inside controller revision
	for _, revision := range revisions.Items {
		if revision.Revision == resourceGeneration {
			controllerRevisionHash, valExist := revision.ObjectMeta.Labels[appsV1.DefaultDaemonSetUniqueLabelKey]
			if !valExist {
				log.WithFields(log.Fields{
					fmt.Sprintf("%T", registryData): registryData.GetName(),
					"namespace":                     namespace,
					"revision":                      resourceGeneration,
				}).Warn("Cannot find controller-revision-hash label inside ControllerRevision. cannot start watch on pods")
				return errors.New("Cannot find controller-revision-hash label inside ControllerRevision. cannot start watch on pods")
			}
			log.WithFields(log.Fields{
				fmt.Sprintf("%T", registryData): registryData.GetName(),
				"namespace":                     namespace,
				"revision":                      resourceGeneration,
				"controllerRevisionHash":        controllerRevisionHash,
			}).Debug("ControllerRevision found  find controller-revision-hash calling Pods Manager")
			// start watch on pods
			podListOptions := metaV1.ListOptions{LabelSelector: labels.SelectorFromSet(map[string]string{
				appsV1.DefaultDaemonSetUniqueLabelKey: controllerRevisionHash,
			}).String()}
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
		fmt.Sprintf("%T", registryData): registryData.GetName(),
		"namespace":                     namespace,
		"revision":                      resourceGeneration,
	}).Error("Cannot find resourceVersion in ControllerRevision. cannot start watch on pods")
	return errors.New("Cannot find resourceVersion in ControllerRevision. cannot start watch on pods")
}
