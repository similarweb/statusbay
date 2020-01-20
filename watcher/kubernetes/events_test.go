package kuberneteswatcher_test

import (
	"context"
	"testing"
	"time"

	kuberneteswatcher "statusbay/watcher/kubernetes"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func NewEventsMock(client *fake.Clientset) *kuberneteswatcher.EventsManager {

	eventManager := kuberneteswatcher.NewEventsManager(client)

	return eventManager
}

func TestWatchReceivedCount(t *testing.T) {
	client := fake.NewSimpleClientset()

	eventManager := NewEventsMock(client)

	event1 := &v1.Event{Message: "message", ObjectMeta: metav1.ObjectMeta{Name: "a", CreationTimestamp: metav1.Time{Time: time.Now()}}}
	event2 := &v1.Event{Message: "message1", ObjectMeta: metav1.ObjectMeta{Name: "b", CreationTimestamp: metav1.Time{Time: time.Now()}}}
	event3 := &v1.Event{Message: "message2", ObjectMeta: metav1.ObjectMeta{Name: "c", CreationTimestamp: metav1.Time{Time: time.Now().Add(-time.Hour)}}}

	listOptions := metav1.ListOptions{}

	ctx := context.Background()

	watchData := kuberneteswatcher.WatchEvents{
		ListOptions: listOptions,
		Namespace:   "default",
		Ctx:         ctx,
	}

	eventChan := eventManager.Watch(watchData)

	messageCount := 0
	go func() {
		for {
			select {
			case <-eventChan:
				messageCount = messageCount + 1
			case <-ctx.Done():
				return
			}
		}
	}()

	time.Sleep(time.Second)
	client.CoreV1().Events("default").Create(event1)
	client.CoreV1().Events("default").Create(event2)
	client.CoreV1().Events("default").Create(event3)
	time.Sleep(time.Second)

	if messageCount != 2 {
		t.Fatalf("unexpected count of received events, got %d expected %d", messageCount, 2)
	}

}

func TestWatchMark(t *testing.T) {
	client := fake.NewSimpleClientset()

	eventManager := NewEventsMock(client)

	event1 := &v1.Event{Message: "OOMKill message", ObjectMeta: metav1.ObjectMeta{Name: "a", CreationTimestamp: metav1.Time{Time: time.Now()}}}

	listOptions := metav1.ListOptions{}

	ctx := context.Background()

	watchData := kuberneteswatcher.WatchEvents{
		ListOptions: listOptions,
		Namespace:   "default",
		Ctx:         ctx,
	}

	eventChan := eventManager.Watch(watchData)

	time.Sleep(time.Second)
	client.CoreV1().Events("default").Create(event1)

	event := <-eventChan

	if event.Message != "OOMKill message" {
		t.Fatalf("unexpected event message, got %s expected %s", event.Message, "OOMKill message")
	}

}
