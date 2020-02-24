package kuberneteswatcher_test

import (
	"context"
	"fmt"
	kuberneteswatcher "statusbay/watcher/kubernetes"
	"statusbay/watcher/kubernetes/common"
	"statusbay/watcher/kubernetes/testutil"
	"sync"
	"testing"
	"time"

	appsV1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func updateStatefulsetMock(client *fake.Clientset, namespace string, statefulset *appsV1.StatefulSet) {

	statefulset.Status.Replicas = 2
	client.AppsV1().StatefulSets(namespace).Update(statefulset)
}

func createStatefulSetMock(client *fake.Clientset, name string, namespace string, labels map[string]string) *appsV1.StatefulSet {
	var StatefulSetReplicas int32 = 1
	statefulset := &appsV1.StatefulSet{
		Spec: appsV1.StatefulSetSpec{
			Template: v1.PodTemplateSpec{
				ObjectMeta: metaV1.ObjectMeta{
					Name:   name,
					Labels: labels,
				},
				Spec: v1.PodSpec{},
			},
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:   name,
			Labels: labels,
			Annotations: map[string]string{
				"statusbay.io/application-name":          "application",
				"statusbay.io/report-deploy-by":          "foo@example.com",
				"statusbay.io/report-slack-channels":     "#testchannel",
				"statusbay.io/progress-deadline-seconds": "10",
			},
		},
	}
	statefulset.Spec.Replicas = &StatefulSetReplicas
	statefulset, _ = client.AppsV1().StatefulSets(namespace).Create(statefulset)
	return statefulset
}

func NewStatefulSetManagerMock(client *fake.Clientset) (*kuberneteswatcher.StatefulsetManager, *testutil.MockStorage, *MockControllerRevisionManager) {
	maxDeploymentTime, _ := time.ParseDuration("10m")
	eventManager := kuberneteswatcher.NewEventsManager(client)
	registryManager, Mockstorage := NewRegistryMock()
	serviceManager := NewServiceManagerMockMock(client)
	podManager := kuberneteswatcher.NewPodsManager(client, eventManager)
	controllerRevisionManager := NewControllerRevisionManagerMock(client, podManager)
	runningApplies := registryManager.LoadRunningApplies()
	statefulsetManager := kuberneteswatcher.NewStatefulsetManager(client, eventManager, registryManager, serviceManager, controllerRevisionManager, runningApplies, maxDeploymentTime)

	var wg *sync.WaitGroup
	ctx := context.Background()

	eventManager.Serve(ctx, wg)
	serviceManager.Serve(ctx, wg)
	podManager.Serve(ctx, wg)
	statefulsetManager.Serve(ctx, wg)

	return statefulsetManager, Mockstorage, controllerRevisionManager

}

func TestStatefulsetWatch(t *testing.T) {
	client := fake.NewSimpleClientset()
	_, Mockstorage, controllerRevisionManager := NewStatefulSetManagerMock(client)
	name, app, namespace := "application", "app", "test-ns"
	labels := map[string]string{"name": name, "app": app}
	controllerRevisionHash := "RaezoTepie7p"

	statefulsetObj := createStatefulSetMock(client, name, namespace, labels)
	time.Sleep(time.Second)

	// Update the number of replica
	updateStatefulsetMock(client, namespace, statefulsetObj)

	// First we create the pod with the same ControllerRevisionHash
	pod := createPod(
		client,
		v1.PodRunning,
		statefulsetObj.GetName(),
		namespace,
		labels)
	time.Sleep(time.Second)

	// Add the expected pod label for Statefulset ControllerRevision
	pod.ObjectMeta.Labels[appsV1.ControllerRevisionHashLabelKey] = fmt.Sprintf("%s-%s",
		statefulsetObj.ObjectMeta.Name, controllerRevisionHash)

	// Create Revision Hash
	revision := createControllerRevisionMock(
		client,
		"controllerrevision",
		namespace,
		controllerRevisionHash,
		"controller.kubernetes.io/hash",
		labels)

	// We need both Resource Generation and revision.Revision in order to compare them in ControllerRevision
	revision.Revision = statefulsetObj.ObjectMeta.Generation
	time.Sleep(time.Second)

	event1 := &v1.Event{Message: "message for statefulset", ObjectMeta: metaV1.ObjectMeta{Name: "a", CreationTimestamp: metaV1.Time{Time: time.Now()}}}
	client.CoreV1().Events(namespace).Create(event1)

	NotValidControllerRevisionHashlabelKey := controllerRevisionManager.Error
	application := Mockstorage.MockWriteDeployment["1"]
	_ = application.Schema.Resources.Statefulsets["test-statefulset"]

	var expectedProgressDeadLine int64 = 10

	t.Run("running_statefulsets", func(t *testing.T) {

		if len(application.Schema.Resources.Statefulsets) != 1 {
			t.Fatalf("unexpected number of statefulsets running, got %d expected %d", len(application.Schema.Resources.Statefulsets), 1)
		}
	})

	t.Run("controller_revision_valid_hash_label_key", func(t *testing.T) {

		if NotValidControllerRevisionHashlabelKey != nil {
			t.Fatalf(NotValidControllerRevisionHashlabelKey.Error())
		}
	})

	t.Run("statefulset_schema_data", func(t *testing.T) {
		if application.Status != "running" {
			t.Fatalf("unexpected apply status, got %s expected %s", application.Status, "running")
		}
		if application.Schema.Application != name {
			t.Fatalf("unexpected application name, got %s expected %s", application.Schema.Application, name)
		}
		if application.Schema.Namespace != namespace {
			t.Fatalf("unexpected application namespace, got %s expected %s", application.Schema.Namespace, namespace)
		}
		if application.Schema.DeploymentDescription != common.ApplyStatusDescriptionRunning {
			t.Fatalf("unexpected status description, got %s expected %s", application.Schema.Namespace, common.ApplyStatusDescriptionRunning)
		}
		if application.Schema.DeployBy != "foo@example.com" {
			t.Fatalf("unexpected field deployby , got %s expected %s", application.Schema.DeployBy, statefulsetObj.ObjectMeta.Labels["statusbay.io/report-deploy-by"])
		}
		if application.Schema.Resources.Statefulsets["application"].ProgressDeadlineSeconds != expectedProgressDeadLine {
			t.Fatalf("unexpected values for ProgressDeadline field , got %d expected %d", application.Schema.Resources.Statefulsets["application"].ProgressDeadlineSeconds, expectedProgressDeadLine)
		}
		if len(application.Schema.Resources.Statefulsets["application"].Statefulset.Annotations) != 4 {
			t.Fatalf("unexpected amount of Lables values for statefulset annotations field , got %d expected %d", len(application.Schema.Resources.Statefulsets["application"].Statefulset.Annotations), 4)
		}

	})
}
