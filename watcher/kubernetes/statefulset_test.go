package kuberneteswatcher_test

import (
	"reflect"
	kuberneteswatcher "statusbay/watcher/kubernetes"
	"statusbay/watcher/kubernetes/testutil"
	"testing"
	"time"

	appsV1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func createStatefulSettMock(client *fake.Clientset, name string, labels map[string]string) *appsV1.StatefulSet {
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
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:   name,
			Labels: labels,
			Annotations: map[string]string{
				"statusbay.io/application-name":          "application",
				"statusbay.io/report-deploy-by":          "testme@similarweb.com",
				"statusbay.io/report-slack-channels":     "#testchannel",
				"statusbay.io/progress-deadline-seconds": "10",
			},
		},
	}
	statefulset.Spec.Replicas = &StatefulSetReplicas
	statefulset, _ = client.AppsV1().StatefulSets("pe").Create(statefulset)
	return statefulset
}

func NewStatefulSetManagerMock(client *fake.Clientset) (*kuberneteswatcher.StatefulsetManager, *testutil.MockStorage) {
	maxDeploymentTime, _ := time.ParseDuration("10m")
	eventManager := kuberneteswatcher.NewEventsManager(client)
	registryManager, Mockstorage, _ := NewRegistryMock()
	serviceManager := NewServiceManagerMockMock(client)
	podManager := kuberneteswatcher.NewPodsManager(client, eventManager)
	controllerRevisionManager := NewControllerRevisionMock(client, podManager)

	statefulsetManager := kuberneteswatcher.NewStatefulsetManager(client, eventManager, registryManager, serviceManager, controllerRevisionManager, maxDeploymentTime)

	eventManager.Serve()
	serviceManager.Serve()
	podManager.Serve()
	statefulsetManager.Serve()

	return statefulsetManager, Mockstorage

}

func TestStatefulsetWatch(t *testing.T) {
	client := fake.NewSimpleClientset()
	_, Mockstorage := NewStatefulSetManagerMock(client)
	name, app := "application", "app"
	labels := map[string]string{"name": name, "app": app}

	time.Sleep(time.Second)

	statefulsetObj := createStatefulSettMock(client, name, labels)
	// Create Revision Hash
	controllerRevisionHash := "RaezoTepie7p"
	revision := &appsV1.ControllerRevision{
		ObjectMeta: metaV1.ObjectMeta{
			Name:      "controllerrevision",
			Namespace: statefulsetObj.GetNamespace(),
			Labels: map[string]string{
				"statusbay.io/application-name":       name,
				"app":                                 "application",
				"name":                                statefulsetObj.GetName(),
				appsV1.ControllerRevisionHashLabelKey: controllerRevisionHash,
			},
		},
	}
	// We need both Resource Generation and revision.Revision in order to compare them in ControllerRevision
	resourceGeneration := statefulsetObj.ObjectMeta.Generation
	revision.Revision = resourceGeneration
	client.AppsV1().ControllerRevisions("pe").Create(revision)
	time.Sleep(time.Second)

	// Create the pods with the same ControllerRevisionHash
	_ = createRunningPod(client, statefulsetObj.GetName(), controllerRevisionHash)
	time.Sleep(time.Second)

	// Trigger Application Apply
	application := Mockstorage.MockWriteDeployment[1]

	expectedReportTo := []string{"testme@similarweb.com", "#testchannel"}
	var expectedProgressDeadLine int64 = 10

	t.Run("running_statefulsets", func(t *testing.T) {

		if len(application.Schema.Resources.Statefulsets) != 1 {
			t.Fatalf("unexpected number of statefulsets running, got %d expected %d", len(application.Schema.Resources.Statefulsets), 1)
		}
	})

	t.Run("statefulset_schema_data", func(t *testing.T) {
		if application.Status != "running" {
			t.Fatalf("unexpected apply status, got %s expected %s", application.Status, "running")
		}
		if application.Schema.Application != name {
			t.Fatalf("unexpected application name, got %s expected %s", application.Schema.Application, name)
		}
		if application.Schema.Namespace != "pe" {
			t.Fatalf("unexpected application namespace, got %s expected %s", application.Schema.Namespace, "pe")
		}
		if application.Schema.DeploymentDescription != kuberneteswatcher.DeploymentStatusDescriptionRunning {
			t.Fatalf("unexpected status description, got %s expected %s", application.Schema.Namespace, kuberneteswatcher.DeploymentStatusDescriptionRunning)
		}
		if application.Schema.DeployBy != "testme@similarweb.com" {
			t.Fatalf("unexpected field deployby , got %s expected %s", application.Schema.DeployBy, statefulsetObj.ObjectMeta.Labels["statusbay.io/report-deploy-by"])
		}
		if !reflect.DeepEqual(application.Schema.ReportTo, expectedReportTo) {
			t.Fatalf("unexpected values for ReportTo field , got %s expected %s", application.Schema.ReportTo, expectedReportTo)
		}
		if application.Schema.Resources.Statefulsets["application"].ProgressDeadlineSeconds != expectedProgressDeadLine {
			t.Fatalf("unexpected values for ProgressDeadline field , got %d expected %d", application.Schema.Resources.Statefulsets["application"].ProgressDeadlineSeconds, expectedProgressDeadLine)
		}
		if len(application.Schema.Resources.Statefulsets["application"].Statefulset.Labels) != 2 {
			t.Fatalf("unexpected amount of Lables values for statefulset labeles field , got %d expected %d", len(application.Schema.Resources.Statefulsets["application"].Statefulset.Labels), 2)
		}

	})
}
