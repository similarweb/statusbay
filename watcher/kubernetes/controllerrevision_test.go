package kuberneteswatcher_test

import (
	"context"
	"fmt"
	kuberneteswatcher "statusbay/watcher/kubernetes"
	"sync"

	appsV1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

type MockControllerRevisionManager struct {
	// Kubernetes client
	client *fake.Clientset

	Error error
	// to find pods related
	podsManager *kuberneteswatcher.PodsManager
}

// NewControllerRevisionManagerMock a new mock of Controller Revision manager.
func NewControllerRevisionManagerMock(client *fake.Clientset, podsManager *kuberneteswatcher.PodsManager) *MockControllerRevisionManager {
	var podManager *kuberneteswatcher.PodsManager
	podManager = podsManager

	var wg *sync.WaitGroup
	ctx := context.Background()
	if podManager == nil {
		eventManager := kuberneteswatcher.NewEventsManager(client)
		podManager = kuberneteswatcher.NewPodsManager(client, eventManager)
		podManager.Serve(ctx, wg)
		eventManager.Serve(ctx, wg)
	}
	//controllerRevisionManager := kuberneteswatcher.NewControllerRevisionManager(client, podManager)
	return &MockControllerRevisionManager{
		client:      client,
		podsManager: podsManager,
	}
}

// stringInMap checks if one of the key's value in a map equals to a string
func stringInMap(str string, dict map[string]string) bool {
	for _, v := range dict {
		if v == str {
			return true
		}
	}
	return false
}

// WatchControllerRevisionPods dummy interface.
func (mcr *MockControllerRevisionManager) WatchControllerRevisionPods(ctx context.Context, registryData kuberneteswatcher.RegistryData, resourceGeneration int64, controllerRevisionlabels map[string]string, controllerRevisionHashlabelKey string, controllerRevisionPodLabelValuePerfix string, namespace string) error {
	return mcr.Error
}

// WatchControllerRevisionPodsRetry Implement a check in the interface to check whether a  controllerRevisionHashlabelKey is valid.
func (mcr *MockControllerRevisionManager) WatchControllerRevisionPodsRetry(ctx context.Context, registryData kuberneteswatcher.RegistryData, resourceGeneration int64, controllerRevisionlabels map[string]string, controllerRevisionHashlabelKey string, controllerRevisionPodLabelValuePerfix string, namespace string, backOffParams *kuberneteswatcher.BackoffParams) error {
	// expectedOptionsOfControllerRevisionHashlabelKey in a Map
	crhlk := map[string]string{"daemonset": appsV1.DefaultDaemonSetUniqueLabelKey, "statefulset": "controller.kubernetes.io/hash"}
	if !stringInMap(controllerRevisionHashlabelKey, crhlk) {
		mcr.Error = fmt.Errorf(fmt.Sprintf(
			"The value:%s of controllerRevisionHashlabelKey is not valid,The following values are valid [ statefulset: %s, daemonset: %s ]", controllerRevisionHashlabelKey, crhlk["statefulset"], crhlk["daemonset"]))

		return mcr.Error
	}
	return mcr.Error
}

// createControllerRevisionMock will create a mock a ControllerRevision Object.
func createControllerRevisionMock(client *fake.Clientset, name string, namespace string, controllerRevisionHash string, controllerRevisionHashlabelKey string, labels map[string]string) *appsV1.ControllerRevision {
	revision := &appsV1.ControllerRevision{
		ObjectMeta: metaV1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			Labels:    labels,
		},
	}
	revision.ObjectMeta.Labels[controllerRevisionHashlabelKey] = controllerRevisionHash
	client.AppsV1().ControllerRevisions(namespace).Create(revision)
	return revision
}

// createPod with a specific status in a specific namespace with Labels.
func createPod(client *fake.Clientset, phase v1.PodPhase, name string, namespace string, labels map[string]string) *v1.Pod {
	PodStatus := v1.PodStatus{
		Phase: phase,
	}
	pod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:   name,
			Labels: labels,
		},
		Status: PodStatus,
	}
	client.CoreV1().Pods(namespace).Create(pod)
	return pod
}
