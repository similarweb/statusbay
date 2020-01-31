package kuberneteswatcher_test

import (
	"testing"
	"time"

	appsV1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	kuberneteswatcher "statusbay/watcher/kubernetes"

	"k8s.io/client-go/kubernetes/fake"

	gtestutil "statusbay/testutil"
	"statusbay/watcher/kubernetes/testutil"
)

const DesiredNumberScheduled = 1

func createDaemonSetMock(client *fake.Clientset, name string, labels map[string]string) *appsV1.DaemonSet {
	daemonset := &appsV1.DaemonSet{
		Spec: appsV1.DaemonSetSpec{
			Template: v1.PodTemplateSpec{
				ObjectMeta: metaV1.ObjectMeta{
					Name: name,
					Labels: map[string]string{
						"app":  "application",
						"name": name,
					},
				},
				Spec: v1.PodSpec{},
			},
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app":  "application",
					"name": name,
				},
			},
		},
		ObjectMeta: metaV1.ObjectMeta{
			Name:   name,
			Labels: labels,
			Annotations: map[string]string{
				"statusbay.io/application-name":       "custom-application-name",
				"statusbay.io/report-deploy-by":       "test@similarweb.com",
				"statusbay.io/report-slack-channels":  "#channel",
				"statusbay.io/alerts-statuscake-tags": "fluentd",
			},
		},
	}
	daemonset.Status.DesiredNumberScheduled = DesiredNumberScheduled
	client.AppsV1().DaemonSets("pe").Create(daemonset)
	return daemonset
}
func NewDaemonSetManagerMock(client *fake.Clientset) (*kuberneteswatcher.DaemonsetManager, *testutil.MockStorage, *gtestutil.MockSlack) {
	maxDeploymentTime, _ := time.ParseDuration("10m")
	eventManager := NewEventsMock(client)
	registryManager, storage, slack := NewRegistryMock()
	serviceManager := NewServiceManagerMockMock(client)
	podManager := kuberneteswatcher.NewPodsManager(client, eventManager)
	controllerRevisionManager := NewControllerRevisionMock(client, podManager)
	daemonsetManager := kuberneteswatcher.NewDaemonsetManager(client, eventManager, registryManager, serviceManager, controllerRevisionManager, maxDeploymentTime)
	daemonsetManager.Serve()
	serviceManager.Serve()
	podManager.Serve()
	return daemonsetManager, storage, slack
}

func TestDaemonsetWatch(t *testing.T) {
	client := fake.NewSimpleClientset()
	_, storage, slack := NewDaemonSetManagerMock(client)
	labels := map[string]string{"app": "application"}
	time.Sleep(time.Second)

	daemonsetObj := createDaemonSetMock(client, "test-daemonset", labels)
	// create matching revision object
	controllerRevisionHash := "6848fd6f74"
	revision := &appsV1.ControllerRevision{
		ObjectMeta: metaV1.ObjectMeta{
			Name: "controllerrevision",
			Labels: map[string]string{
				"statusbay.io/application-name":       "custom-application-name",
				"app":                                 "application",
				"name":                                daemonsetObj.GetName(),
				appsV1.ControllerRevisionHashLabelKey: controllerRevisionHash,
			},
		},
	}
	resourceGeneration := daemonsetObj.ObjectMeta.Generation
	revision.Revision = resourceGeneration
	client.AppsV1().ControllerRevisions("pe").Create(revision)
	time.Sleep(time.Second)
	// create matchin pods to the revision hash
	_ = createRunningPod(client, daemonsetObj.GetName(), controllerRevisionHash)
	time.Sleep(time.Second)
	// verify daemonset deployed
	application := storage.MockWriteDeployment[1]
	_ = application.Schema.Resources.Daemonsets["test-daemonset"]
	t.Run("slack_message", func(t *testing.T) {
		if len(slack.PostMessageRequest) != DesiredNumberScheduled {
			t.Fatalf("unexpected slack report, got %d expected %d", len(slack.PostMessageRequest), DesiredNumberScheduled)
		}
	})
	t.Run("daemonset_schema_data", func(t *testing.T) {
		if application.Status != "running" {
			t.Fatalf("unexpected apply status, got %s expected %s", application.Status, "running")
		}
		if application.Schema.Application != "custom-application-name" {
			t.Fatalf("unexpected application name, got %s expected %s", application.Schema.Application, "custom-application-name")
		}
		if application.Schema.Namespace != "pe" {
			t.Fatalf("unexpected application namespace, got %s expected %s", application.Schema.Namespace, "pe")
		}
		if application.Schema.DeploymentDescription != kuberneteswatcher.DeploymentStatusDescriptionRunning {
			t.Fatalf("unexpected status description, got %s expected %s", application.Schema.Namespace, kuberneteswatcher.DeploymentStatusDescriptionRunning)
		}
	})
}
