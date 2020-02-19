package kuberneteswatcher_test

import (
	"context"
	kuberneteswatcher "statusbay/watcher/kubernetes"
	"statusbay/watcher/kubernetes/common"
	"statusbay/watcher/kubernetes/testutil"
	"sync"
	"testing"

	"time"

	appsV1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func createMockDeploymentData(registryManager *kuberneteswatcher.RegistryManager, registryRow *kuberneteswatcher.RegistryRow, applyEvent kuberneteswatcher.ApplyEvent, progressDeadlineSeconds string) *kuberneteswatcher.DeploymentData {

	maxDeploymentTime, _ := time.ParseDuration(progressDeadlineSeconds)
	client := fake.NewSimpleClientset()

	eventManager := NewEventsMock(client)
	replicasetManager := NewReplicasetMock(client)
	serviceManager := NewServiceManagerMockMock(client)
	deploymentManager := kuberneteswatcher.NewDeploymentManager(client, eventManager, registryManager, replicasetManager, serviceManager, maxDeploymentTime)

	return deploymentManager.AddNewDeployment(applyEvent, registryRow, 3)

}

func updateDeploymentMock(client *fake.Clientset, deployment *appsV1.Deployment) {

	deployment.Status.Replicas = 2
	client.AppsV1().Deployments("pe").Update(deployment)

}

func createDeploymentMock(client *fake.Clientset, name string, labels map[string]string) *appsV1.Deployment {
	replicas := int32(1)
	progressDeadlineSeconds := int32(300)
	deployment := &appsV1.Deployment{
		Spec: appsV1.DeploymentSpec{
			ProgressDeadlineSeconds: &progressDeadlineSeconds,
			Replicas:                &replicas,
			Selector: &metaV1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "application",
				},
			},
		},
		ObjectMeta: metaV1.ObjectMeta{
			Name:   name,
			Labels: labels,
			Annotations: map[string]string{
				"statusbay.io/application-name":       "custom-application-name",
				"statusbay.io/report-deploy-by":       "foo@example.com",
				"statusbay.io/report-slack-channels":  "#channel",
				"statusbay.io/alerts-statuscake-tags": "nginx",
				"statusbay.io/kibana-query":           "application: statusbay AND mode: watcher",
				"statusbay.io/kibana-query-1":         "application: statusbay AND mode: client",
			},
		},
	}

	client.AppsV1().Deployments("pe").Create(deployment)

	return deployment
}

func GetFakeDeployment(progressDeadlineSeconds int32) *appsV1.Deployment {

	fakeDeployment := &appsV1.Deployment{
		Spec: appsV1.DeploymentSpec{
			ProgressDeadlineSeconds: &progressDeadlineSeconds,
		},
		ObjectMeta: metaV1.ObjectMeta{
			Namespace: "pe",
			Annotations: map[string]string{
				"statusbay.io/report-deploy-by":      "foo@example.com",
				"statusbay.io/report-slack-channels": "#channel",
			},
		},
	}

	return fakeDeployment

}

func NewDeploymentManagerMock(client *fake.Clientset) (*kuberneteswatcher.DeploymentManager, *testutil.MockStorage) {

	maxDeploymentTime, _ := time.ParseDuration("10m")

	eventManager := NewEventsMock(client)
	registryManager, storage := NewRegistryMock()
	replicasetManager := NewReplicasetMock(client)
	serviceManager := NewServiceManagerMockMock(client)
	deploymentManager := kuberneteswatcher.NewDeploymentManager(client, eventManager, registryManager, replicasetManager, serviceManager, maxDeploymentTime)

	var wg *sync.WaitGroup
	ctx := context.Background()

	deploymentManager.Serve(ctx, wg)
	serviceManager.Serve(ctx, wg)
	replicasetManager.Serve(ctx, wg)
	return deploymentManager, storage

}

func TestDeploymentsWatch(t *testing.T) {
	client := fake.NewSimpleClientset()
	_, storage := NewDeploymentManagerMock(client)

	labels := map[string]string{
		"app.kubernetes.io/name": "custom-application-name",
	}

	deploymentObj := createDeploymentMock(client, "test-deployment", labels)

	time.Sleep(time.Second)
	updateDeploymentMock(client, deploymentObj)

	replicaset := &appsV1.ReplicaSet{
		Spec: appsV1.ReplicaSetSpec{
			Selector: &metaV1.LabelSelector{
				MatchLabels: map[string]string{
					"pod-template-hash": "1",
				},
			},
		},
		ObjectMeta: metaV1.ObjectMeta{
			Name: "replicaset",
			Labels: map[string]string{
				"app": "application",
			},
		},
	}

	createServiceMock(client, "service")
	client.AppsV1().ReplicaSets("pe").Create(replicaset)
	time.Sleep(time.Second)
	replicaset.Status.Replicas = 2
	client.AppsV1().ReplicaSets("pe").Update(replicaset)

	event1 := &v1.Event{Message: "message", ObjectMeta: metaV1.ObjectMeta{Name: "a", CreationTimestamp: metaV1.Time{Time: time.Now()}}}
	client.CoreV1().Events("pe").Create(event1)

	time.Sleep(time.Second)

	application := storage.MockWriteDeployment["1"]

	deployment := application.Schema.Resources.Deployments["test-deployment"]

	t.Run("deployment_schema_data", func(t *testing.T) {

		if len(deployment.Events) != 1 {
			t.Fatalf("unexpected deployment events, got %d expected %d", len(deployment.Events), 1)
		}
	})

	t.Run("deployment_schema_data", func(t *testing.T) {
		if application.Status != "running" {
			t.Fatalf("unexpected deployment status, got %s expected %s", application.Status, "running")
		}
		if application.Schema.Application != "custom-application-name" {
			t.Fatalf("unexpected application name, got %s expected %s", application.Schema.Application, "custom-application-name")
		}
		if application.Schema.Namespace != "pe" {
			t.Fatalf("unexpected application namespace, got %s expected %s", application.Schema.Namespace, "pe")
		}
		if application.Schema.DeploymentDescription != common.ApplyStatusDescriptionRunning {
			t.Fatalf("unexpected status description, got %s expected %s", application.Schema.Namespace, common.ApplyStatusDescriptionRunning)
		}

	})

	t.Run("replicaset", func(t *testing.T) {
		if len(deployment.Replicaset) != 1 {
			t.Fatalf("unexpected replicaset count, got %d expected %d", len(deployment.Replicaset), 1)
		}
	})

}
