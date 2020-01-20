package kuberneteswatcher_test

import (
	"k8s.io/client-go/kubernetes/fake"
	kuberneteswatcher "statusbay/watcher/kubernetes"
)

func NewControllerRevisionMock(client *fake.Clientset, podsManager *kuberneteswatcher.PodsManager) *kuberneteswatcher.ControllerRevisionManager {
	var podManager *kuberneteswatcher.PodsManager
	podManager = podsManager
	if podManager == nil {
		eventManager := kuberneteswatcher.NewEventsManager(client)
		podManager = kuberneteswatcher.NewPodsManager(client, eventManager)
		podManager.Serve()
	}
	controllerRevisionManager := kuberneteswatcher.NewControllerReisionManager(client, podManager)
	return controllerRevisionManager
}
