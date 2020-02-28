package kuberneteswatcher_test

import (
	"context"
	"errors"
	kuberneteswatcher "statusbay/watcher/kubernetes"
	"sync"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
	coreV1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

type MockRegistryData struct {
	callCount int
}

// All functions under the RegistryData interface
func (mrd *MockRegistryData) UpdatePodEvents(podName string, pvcName string, event kuberneteswatcher.EventMessages) error {
	mrd.callCount++
	return nil
}

func (mrd *MockRegistryData) UpdatePod(pod *v1.Pod, status string) error {
	return errors.New("Implement me")
}

func (mrd *MockRegistryData) NewPod(pod *v1.Pod) error {
	return errors.New("Implement me")
}

func (mrd *MockRegistryData) NewService(pod *v1.Service) error {
	return errors.New("Implement me")
}

func (mrd *MockRegistryData) UpdateServiceEvents(name string, event kuberneteswatcher.EventMessages) error {
	return errors.New("Implement me")
}

func (mrd *MockRegistryData) GetName() string {
	return "www"
}

func createPvcMock(client *fake.Clientset, namespace, volumeName, pvcName string) *coreV1.PersistentVolumeClaim {
	pvc := &v1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name:      pvcName,
			Namespace: namespace,
		},
		Spec: v1.PersistentVolumeClaimSpec{VolumeName: volumeName},
	}

	pvc, _ = client.CoreV1().PersistentVolumeClaims(namespace).Create(pvc)
	return pvc
}

// NewPvcManagerMock creates a pvcManager mock object.
func NewPvcManagerMock(client *fake.Clientset) *kuberneteswatcher.PvcManager {
	eventManager := kuberneteswatcher.NewEventsManager(client)
	pvcManager := kuberneteswatcher.NewPvcManager(client, eventManager)

	// Start the pvcManger
	var wg *sync.WaitGroup
	ctx := context.Background()
	pvcManager.Serve(ctx, wg)
	eventManager.Serve(ctx, wg)

	return pvcManager
}

func TestPvcWatch(t *testing.T) {
	client := fake.NewSimpleClientset()
	pvcManager := NewPvcManagerMock(client)
	namespace, pvcName := "namespace", "pvc"
	lg := log.WithField("test", "TestPvcWatch")
	ctx := context.Background()

	MockRegistryData := &MockRegistryData{}

	pvcManager.Watch <- kuberneteswatcher.WatchPvcData{
		ListOptions:  metav1.ListOptions{},
		RegistryData: MockRegistryData,
		Namespace:    namespace,
		Ctx:          ctx,
		Pod:          "pvc-pod",
		LogEntry:     *lg,
	}

	time.Sleep(2 * time.Second)
	createPvcMock(client, namespace, "www", pvcName)
	time.Sleep(2 * time.Second)

	event1 := &v1.Event{Message: "message number 1", ObjectMeta: metaV1.ObjectMeta{Name: "www", CreationTimestamp: metaV1.Time{Time: time.Now()}}}
	event2 := &v1.Event{Message: "message number 2", ObjectMeta: metaV1.ObjectMeta{Name: "www2", CreationTimestamp: metaV1.Time{Time: time.Now()}}}
	event3 := &v1.Event{Message: "message number 3", ObjectMeta: metaV1.ObjectMeta{Name: "www3", CreationTimestamp: metaV1.Time{Time: time.Now()}}}
	client.CoreV1().Events(namespace).Create(event1)
	client.CoreV1().Events(namespace).Create(event2)
	client.CoreV1().Events(namespace).Create(event3)
	time.Sleep(time.Second)

	expectedEvents := 3
	t.Run("pvc_events_count", func(t *testing.T) {
		if MockRegistryData.callCount != expectedEvents {
			t.Fatalf("Unexpected number of pvc events running, got %d expected %d", MockRegistryData.callCount, expectedEvents)
		}
	})
}
