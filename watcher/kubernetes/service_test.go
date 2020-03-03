package kuberneteswatcher_test

import (
	"context"
	kuberneteswatcher "statusbay/watcher/kubernetes"
	"statusbay/watcher/kubernetes/common"
	"sync"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes/fake"
)

func createServiceMock(client *fake.Clientset, name, namespace string) {
	service := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
	}

	client.CoreV1().Services(namespace).Create(service)
}

func NewServiceManagerMockMock(client *fake.Clientset) *kuberneteswatcher.ServiceManager {

	eventManager := NewEventsMock(client)
	serviceManager := kuberneteswatcher.NewServiceManager(client, eventManager)
	return serviceManager

}

func TestServiceWatch(t *testing.T) {
	registry, storageMock := NewRegistryMock()

	registryRow := registry.NewApplication("nginx", "default", map[string]string{}, common.ApplyStatusRunning)
	namespace := "default"
	apply := kuberneteswatcher.ApplyEvent{
		Event:        "create",
		ApplyName:    "nginx-deployment",
		ResourceName: "resourceName",
		Namespace:    namespace,
		Kind:         "deployment",
		Hash:         1234,
		Annotations:  map[string]string{},
		Labels:       map[string]string{},
	}

	registryDeploymentData := createMockDeploymentData(registry, registryRow, apply, "10m")
	lg := log.WithField("test", "TestServiceWatch")
	ctx := context.Background()

	client := fake.NewSimpleClientset()

	serviceManager := NewServiceManagerMockMock(client)

	var wg sync.WaitGroup

	serviceManager.Serve(ctx, &wg)

	serviceManager.Watch <- kuberneteswatcher.WatchData{
		ListOptions:  metav1.ListOptions{},
		RegistryData: registryDeploymentData,
		Namespace:    namespace,
		Ctx:          ctx,
		LogEntry:     *lg,
	}

	time.Sleep(time.Second)
	createServiceMock(client, "service-1", namespace)
	createServiceMock(client, "service-2", namespace)
	time.Sleep(time.Second)

	event1 := &v1.Event{Message: "message", InvolvedObject: v1.ObjectReference{Kind: "Service", Name: "service-1"}, ObjectMeta: metav1.ObjectMeta{CreationTimestamp: metav1.Time{Time: time.Now()}}}
	client.CoreV1().Events(namespace).Create(event1)

	time.Sleep(time.Second * 2)

	deployment := storageMock.MockWriteDeployment["1"].Schema.Resources.Deployments["resourceName"]

	if len(deployment.Services) != 2 {
		t.Fatalf("unexpected services count, got %d expected %d", len(deployment.Services), 2)
	}

	if len(*deployment.Services["service-1"].Events) != 1 {
		t.Fatalf("unexpected event count, got %d expected %d", len(*deployment.Services["service-1"].Events), 1)
	}

}
