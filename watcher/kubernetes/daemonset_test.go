package kuberneteswatcher_test

import (
	"context"
	"sync"
	"testing"
	"time"

	appsV1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	kuberneteswatcher "statusbay/watcher/kubernetes"

	"k8s.io/client-go/kubernetes/fake"

	"statusbay/watcher/kubernetes/common"
	"statusbay/watcher/kubernetes/testutil"
)

func updateDaemonsetMock(client *fake.Clientset, namespace string, daemonset *appsV1.DaemonSet) {

	daemonset.Status.DesiredNumberScheduled = 2
	client.AppsV1().DaemonSets(namespace).Update(daemonset)
}

const DesiredNumberScheduled = 1

func createDaemonSetMock(client *fake.Clientset, name string, labels map[string]string) *appsV1.DaemonSet {
	daemonset := &appsV1.DaemonSet{
		Spec: appsV1.DaemonSetSpec{
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
		ObjectMeta: metaV1.ObjectMeta{
			Name:   name,
			Labels: labels,
			Annotations: map[string]string{
				"statusbay.io/application-name":       "custom-application-name",
				"statusbay.io/report-deploy-by":       "test@example.com",
				"statusbay.io/report-slack-channels":  "#channel",
				"statusbay.io/alerts-statuscake-tags": "fluentd",
			},
		},
	}
	daemonset.Status.DesiredNumberScheduled = DesiredNumberScheduled
	client.AppsV1().DaemonSets("pe").Create(daemonset)
	return daemonset
}
func NewDaemonSetManagerMock(client *fake.Clientset) (*kuberneteswatcher.DaemonsetManager, *testutil.MockStorage, *MockControllerRevisionManager) {
	maxDeploymentTime, _ := time.ParseDuration("10m")
	eventManager := NewEventsMock(client)
	registryManager, storage := NewRegistryMock()
	serviceManager := NewServiceManagerMockMock(client)
	pvcManager := NewPvcManagerMock(client)
	podManager := kuberneteswatcher.NewPodsManager(client, eventManager, pvcManager)
	controllerRevisionManager := NewControllerRevisionManagerMock(client, podManager)
	runningApplies := registryManager.LoadRunningApplies()
	daemonsetManager := kuberneteswatcher.NewDaemonsetManager(client, eventManager, registryManager, serviceManager, controllerRevisionManager, runningApplies, maxDeploymentTime)

	var wg sync.WaitGroup
	ctx := context.Background()

	eventManager.Serve(ctx, &wg)
	serviceManager.Serve(ctx, &wg)
	podManager.Serve(ctx, &wg)
	daemonsetManager.Serve(ctx, &wg)
	return daemonsetManager, storage, controllerRevisionManager
}

func TestDaemonsetWatch(t *testing.T) {
	client := fake.NewSimpleClientset()
	_, storage, controllerRevisionManager := NewDaemonSetManagerMock(client)
	controllerRevisionHash := "6848fd6f74"
	labels := map[string]string{
		"app": "application",
	}
	namespace := "pe"

	// Create daemonset object
	daemonsetObj := createDaemonSetMock(client, "test-daemonset", labels)
	time.Sleep(time.Second)

	svc := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "service-1",
			Labels:    labels,
			Namespace: namespace,
		},
	}

	// Update the number of replica
	updateDaemonsetMock(client, namespace, daemonsetObj)

	pod := createPod(client,
		v1.PodRunning,
		daemonsetObj.GetName(),
		namespace,
		labels)
	time.Sleep(time.Second)

	client.CoreV1().Services(namespace).Create(svc)

	pod.ObjectMeta.Labels[appsV1.ControllerRevisionHashLabelKey] = controllerRevisionHash

	// create matching revision object
	deamonsetLabels := map[string]string{
		"statusbay.io/application-name": "custom-application-name",
		"app":                           "application",
		"name":                          daemonsetObj.GetName(),
	}

	revision := createControllerRevisionMock(
		client,
		"controllerrevision",
		"pe",
		controllerRevisionHash,
		appsV1.DefaultDaemonSetUniqueLabelKey,
		deamonsetLabels)

	revision.Revision = daemonsetObj.ObjectMeta.Generation
	client.AppsV1().ControllerRevisions(namespace).Create(revision)

	time.Sleep(time.Second)
	// create matchin pods to the revision hash

	NotValidControllerRevisionHashlabelKey := controllerRevisionManager.Error
	// verify daemonset deployed
	application := storage.MockWriteDeployment["1"]
	_ = application.Schema.Resources.Daemonsets["test-daemonset"]

	t.Run("controller_revision_valid_hash_label_key", func(t *testing.T) {

		if NotValidControllerRevisionHashlabelKey != nil {
			t.Fatalf(NotValidControllerRevisionHashlabelKey.Error())
		}
	})

	t.Run("daemonset_schema_data", func(t *testing.T) {
		if application.Status != "running" {
			t.Fatalf("unexpected apply status, got %s expected %s", application.Status, "running")
		}
		if application.Schema.Application != "custom-application-name" {
			t.Fatalf("unexpected application name, got %s expected %s", application.Schema.Application, "custom-application-name")
		}
		if application.Schema.Namespace != namespace {
			t.Fatalf("unexpected application namespace, got %s expected %s", application.Schema.Namespace, namespace)
		}
		if application.Schema.DeploymentDescription != common.ApplyStatusDescriptionRunning {
			t.Fatalf("unexpected status description, got %s expected %s", application.Schema.Namespace, common.ApplyStatusDescriptionRunning)
		}
	})

	// t.Run("service", func(t *testing.T) {
	// 	if len(daemonsetData.Services) != 1 {
	// 		t.Fatalf("unexpected service count, got %d expected %d", len(daemonsetData.Services), 1)
	// 	}
	// })
}
