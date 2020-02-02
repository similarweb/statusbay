package kuberneteswatcher_test

import (
	kuberneteswatcher "statusbay/watcher/kubernetes"

	appsV1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func NewControllerRevisionMock(client *fake.Clientset, podsManager *kuberneteswatcher.PodsManager) *kuberneteswatcher.ControllerRevisionManager {
	var podManager *kuberneteswatcher.PodsManager
	podManager = podsManager
	if podManager == nil {
		eventManager := kuberneteswatcher.NewEventsManager(client)
		podManager = kuberneteswatcher.NewPodsManager(client, eventManager)
		podManager.Serve()
	}
	controllerRevisionManager := kuberneteswatcher.NewControllerRevisionManager(client, podManager)
	return controllerRevisionManager
}

func createRunningPod(client *fake.Clientset, name string, controllerRevisionHash string) *v1.Pod {
	runningPodStatus := v1.PodStatus{
		Phase: v1.PodRunning,
	}
	event := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
			Labels: map[string]string{
				"app":                                 "application",
				"name":                                name,
				appsV1.ControllerRevisionHashLabelKey: controllerRevisionHash,
			},
		},
		Status: runningPodStatus,
	}
	client.CoreV1().Pods("pe").Create(event)
	return event
}
