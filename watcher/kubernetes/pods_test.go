package kuberneteswatcher_test

import (
	"context"
	kuberneteswatcher "statusbay/watcher/kubernetes"
	"statusbay/watcher/kubernetes/common"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func createPodMock(client *fake.Clientset, name string, status v1.PodStatus, deletionTimestamp *metav1.Time) {
	event1 := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:              name,
			DeletionTimestamp: deletionTimestamp,
		},
		Status: status,
	}

	client.CoreV1().Pods("pe").Create(event1)
}

func NewPodManagerMock() (*fake.Clientset, *kuberneteswatcher.PodsManager) {

	client := fake.NewSimpleClientset()

	eventManager := kuberneteswatcher.NewEventsManager(client)

	podManager := kuberneteswatcher.NewPodsManager(client, eventManager)
	podManager.Serve()
	return client, podManager

}

func TestPodWatch(t *testing.T) {
	registry, storageMock := NewRegistryMock()

	registryDeploymentData := createMockDeploymentData(registry, common.DeploymentStatusRunning)
	lg := log.WithField("test", "TestPodWatch")
	ctx := context.Background()

	client, podManager := NewPodManagerMock()

	podManager.Watch <- kuberneteswatcher.WatchData{
		RegistryData: registryDeploymentData,
		ListOptions:  metav1.ListOptions{},
		Namespace:    "pe",
		Ctx:          ctx,
		LogEntry:     *lg,
	}
	time.Sleep(time.Second)

	nginxWaitingState := v1.ContainerState{
		Waiting: &v1.ContainerStateWaiting{
			Reason:  "waiting Reason message",
			Message: "waiting message text",
		},
	}
	nginxTerminatedtate := v1.ContainerState{
		Terminated: &v1.ContainerStateTerminated{
			Reason:  "crashloopbackoff",
			Message: "Terminated message text",
		},
	}
	nginxPodStatus := v1.PodStatus{
		Phase: v1.PodRunning,
		ContainerStatuses: []v1.ContainerStatus{
			{
				State: nginxWaitingState,
			},
			{
				State: nginxTerminatedtate,
			},
		},
	}

	createPodMock(client, "nginx", nginxPodStatus, nil)
	createPodMock(client, "nginx1", v1.PodStatus{Phase: v1.PodFailed}, nil)
	createPodMock(client, "nginx1", v1.PodStatus{Phase: v1.PodRunning}, nil)
	createPodMock(client, "nginx2", v1.PodStatus{Phase: v1.PodRunning}, &metav1.Time{Time: time.Now()})
	time.Sleep(time.Second * 3)

	pods := storageMock.MockWriteDeployment[1].Schema.Resources.Deployments["application"].Pods
	t.Run("registory_pods", func(t *testing.T) {
		podCount := len(pods)

		if podCount != 3 {
			t.Fatalf("unexpected watch pod count, got %d expected %d", podCount, 3)
		}
	})

	t.Run("pod_status", func(t *testing.T) {

		if *pods["nginx"].Phase != "crashloopbackoff" {
			t.Fatalf("unexpected nginx pod status, got %s expected %s", *pods["nginx"].Phase, "crashloopbackoff")
		}
		if *pods["nginx1"].Phase != string(v1.PodFailed) {
			t.Fatalf("unexpected nginx pod status, got %s expected %s", *pods["nginx"].Phase, string(v1.PodRunning))
		}

	})

	t.Run("pod_waiting_container_status", func(t *testing.T) {
		excepted := "waiting Reason message - waiting message text"
		status := *pods["nginx"].Events
		if status[0].Message != excepted {
			t.Fatalf("unexpected nginx pod status, got %s expected %s", status[0].Message, excepted)
		}
	})

	t.Run("pod_termination_container_status", func(t *testing.T) {
		excepted := "crashloopbackoff - Terminated message text"
		status := *pods["nginx"].Events
		if status[1].Message != excepted {
			t.Fatalf("unexpected nginx pod status, got %s expected %s", status[0].Message, excepted)
		}
	})

	t.Run("pod_deletion_time", func(t *testing.T) {
		status := *pods["nginx2"].Phase
		if status != "Terminated" {
			t.Fatalf("unexpected nginx2 pod status, got %s expected %s", status, "Terminated")
		}
	})

}

func TestPodWatchEvent(t *testing.T) {
	registry, storageMock := NewRegistryMock()

	registryDeploymentData := createMockDeploymentData(registry, common.DeploymentStatusRunning)
	lg := log.WithField("test", "TestPodWatchEvent")
	ctx := context.Background()

	client, podManager := NewPodManagerMock()

	podManager.Watch <- kuberneteswatcher.WatchData{
		RegistryData: registryDeploymentData,
		ListOptions:  metav1.ListOptions{},
		Namespace:    "pe",
		Ctx:          ctx,
		LogEntry:     *lg,
	}
	time.Sleep(time.Second)

	createPodMock(client, "nginx", v1.PodStatus{Phase: v1.PodRunning}, nil)
	time.Sleep(time.Second)
	event1 := &v1.Event{Message: "message", ObjectMeta: metav1.ObjectMeta{Name: "a", CreationTimestamp: metav1.Time{Time: time.Now()}}}
	event2 := &v1.Event{Message: "message", ObjectMeta: metav1.ObjectMeta{Name: "b", CreationTimestamp: metav1.Time{Time: time.Now()}}}
	client.CoreV1().Events("pe").Create(event1)
	client.CoreV1().Events("pe").Create(event2)

	time.Sleep(time.Second)
	pods := storageMock.MockWriteDeployment[1].Schema.Resources.Deployments["application"].Pods

	if len(*pods["nginx"].Events) != 2 {
		t.Fatalf("unexpected watch pod events count, got %d expected %d", len(*pods["nginx"].Events), 2)
	}
}
