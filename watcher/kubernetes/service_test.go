package kuberneteswatcher_test

import (
	"context"
	"fmt"
	kuberneteswatcher "statusbay/watcher/kubernetes"
	"statusbay/watcher/kubernetes/common"
	"testing"
	"time"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes/fake"
)

func createServiceMock(client *fake.Clientset, name string) {
	service := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: "pe",
		},
	}

	client.CoreV1().Services("pe").Create(service)
}

func NewServiceManagerMockMock(client *fake.Clientset) *kuberneteswatcher.ServiceManager {

	serviceManager := kuberneteswatcher.NewServiceManager(client)
	return serviceManager

}

func TestServiceWatch(t *testing.T) {
	registory, storageMock := NewRegistryMock()

	registryDeploymentData := createMockDeploymentData(registory, common.DeploymentStatusRunning)

	ctx := context.Background()

	client := fake.NewSimpleClientset()

	serviceManager := NewServiceManagerMockMock(client)
	serviceManager.Serve()

	createServiceMock(client, "service-1")

	serviceManager.Watch <- kuberneteswatcher.WatchData{
		ListOptions:  metav1.ListOptions{},
		RegistryData: registryDeploymentData,
		Namespace:    "pe",
		Ctx:          ctx,
	}

	time.Sleep(time.Second * 5)

	deployment := storageMock.MockWriteDeployment[1].Schema.Resources.Deployments["application"]

	// TODO.. complete the test when the task https://trello.com/c/VheJxFTE/42-add-deployment-service-to-the-db is completed
	fmt.Println(deployment)

}
