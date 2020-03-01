package kuberneteswatcher_test

import (
	"context"
	"fmt"
	notifierCommon "statusbay/notifiers/common"
	kuberneteswatcher "statusbay/watcher/kubernetes"
	"statusbay/watcher/kubernetes/common"
	"statusbay/watcher/kubernetes/testutil"
	"sync"
	"testing"
	"time"

	appsV1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NewRegistryMock() (*kuberneteswatcher.RegistryManager, *testutil.MockStorage) {

	saveInterval, _ := time.ParseDuration("1s")
	checkFinishDelay := 10 * time.Microsecond
	collectDataAfterApplyFinish := 10 * time.Microsecond

	storageMock := testutil.NewMockStorage()
	reporter := kuberneteswatcher.NewReporter([]notifierCommon.Notifier{})
	registry := kuberneteswatcher.NewRegistryManager(saveInterval, checkFinishDelay, collectDataAfterApplyFinish, storageMock, reporter, "mock-cluster")

	var wg sync.WaitGroup
	ctx := context.Background()

	registry.Serve(ctx, &wg)
	reporter.Serve(ctx, &wg)
	return registry, storageMock

}

func TestNewApplicationDeployment(t *testing.T) {

	registry, _ := NewRegistryMock()

	fakeDeployment := GetFakeDeployment(300)
	registryRow := registry.NewApplication("nginx", fakeDeployment.GetNamespace(), fakeDeployment.GetAnnotations(), common.ApplyStatusRunning)

	testCases := []struct {
		description string
		mutate      func(row *kuberneteswatcher.RegistryRow) interface{}
		expected    interface{}
	}{
		{
			"application name",
			func(row *kuberneteswatcher.RegistryRow) interface{} { return row.DBSchema.Application },
			"nginx",
		},
		{
			"namespace",
			func(row *kuberneteswatcher.RegistryRow) interface{} { return row.DBSchema.Namespace },
			"pe",
		},
		{
			"deploy by",
			func(row *kuberneteswatcher.RegistryRow) interface{} { return row.DBSchema.DeployBy },
			"foo@example.com",
		},
		{
			"report to count",
			func(row *kuberneteswatcher.RegistryRow) interface{} { return len(row.DBSchema.ReportTo) },
			2,
		},
	}

	t.Run("new_application", func(t *testing.T) {

		for _, tc := range testCases {
			value := tc.mutate(registryRow)
			if value != tc.expected {
				t.Fatalf("unexpected %s, got %s expected %s", tc.description, value, tc.expected)
			}
		}

	})

	t.Run("get_application", func(t *testing.T) {

		row := registry.Get("nginx", "pe", "")
		for _, tc := range testCases {
			value := tc.mutate(row)
			if value != tc.expected {
				t.Fatalf("unexpected %s, got %s expected %s", tc.description, value, tc.expected)
			}
		}

	})

	t.Run("get_application", func(t *testing.T) {
		uri := registryRow.GetURI()
		uriExpected := fmt.Sprintf("application/%s", registryRow.GetApplyID())
		if uri != uriExpected {
			t.Fatalf("unexpected deployment count, got %s expected %s", uri, uriExpected)
		}

	})

}

func TestAddDeployment(t *testing.T) {

	registry, _ := NewRegistryMock()

	fakeDeployment := GetFakeDeployment(300)
	annotations := map[string]string{
		"statusbay.io/report-deploy-by":      "foo@example.com",
		"statusbay.io/report-slack-channels": "#channel",
	}
	registryRow := registry.NewApplication("nginx", fakeDeployment.GetNamespace(), annotations, common.ApplyStatusRunning)

	// data := registryRow.AddDeployment("nginx-deployment", "pe", labels, annotations, 3, 300)

	apply := kuberneteswatcher.ApplyEvent{
		Event:        "create",
		ApplyName:    "nginx-deployment",
		ResourceName: "resourceName",
		Namespace:    "default",
		Kind:         "deployment",
		Hash:         1234,
		Annotations:  fakeDeployment.GetAnnotations(),
		Labels:       fakeDeployment.GetLabels(),
	}

	data := createMockDeploymentData(registry, registryRow, apply, "10m")

	// TODO:: add report check
	t.Run("deployment_data", func(t *testing.T) {

		testCases := []struct {
			description string
			mutate      func(row *kuberneteswatcher.DeploymentData) interface{}
			expected    interface{}
		}{
			{
				"deployment name",
				func(deploymentData *kuberneteswatcher.DeploymentData) interface{} {
					return deploymentData.Deployment.Name
				},
				"nginx-deployment",
			},
			{
				"namespace",
				func(deploymentData *kuberneteswatcher.DeploymentData) interface{} {
					return deploymentData.Deployment.Namespace
				},
				"default",
			},
			{
				"desired state",
				func(deploymentData *kuberneteswatcher.DeploymentData) interface{} {
					return deploymentData.Deployment.DesiredState
				},
				int32(3),
			},
			{
				"annotations count",
				func(deploymentData *kuberneteswatcher.DeploymentData) interface{} {
					return len(deploymentData.Deployment.Annotations)
				},
				2,
			},
		}

		for _, tc := range testCases {
			value := tc.mutate(data)
			if value != tc.expected {
				t.Fatalf("unexpected %s, got %s expected %s", tc.description, value, tc.expected)
			}
		}

	})

	t.Run("registry_deployment", func(t *testing.T) {

		row := registry.Get("nginx", "pe", "")
		if len(row.DBSchema.Resources.Deployments) != 1 {
			t.Fatalf("unexpected deployment count, got %d expected %d", len(row.DBSchema.Resources.Deployments), 1)
		}

	})

}

func TestDeploymentData(t *testing.T) {
	registry, _ := NewRegistryMock()

	annotations := map[string]string{
		"statusbay.io/report-deploy-by":      "foo@example.com",
		"statusbay.io/report-slack-channels": "#channel",
	}

	registryRow := registry.NewApplication("nginx", "default", annotations, common.ApplyStatusRunning)

	apply := kuberneteswatcher.ApplyEvent{
		Event:        "create",
		ApplyName:    "nginx-deployment",
		ResourceName: "resourceName",
		Namespace:    "default",
		Kind:         "deployment",
		Hash:         1234,
		Annotations:  annotations,
		Labels:       map[string]string{},
	}

	data := createMockDeploymentData(registry, registryRow, apply, "10m")

	t.Run("update_deployment_Status", func(t *testing.T) {
		deploymentStatus := appsV1.DeploymentStatus{
			Replicas: 10,
		}
		data.UpdateDeploymentStatus(deploymentStatus)

		if data.Status.Replicas != 10 {
			t.Fatalf("unexpected deployment status")
		}

	})

	t.Run("update_deployment_events", func(t *testing.T) {
		event1 := kuberneteswatcher.EventMessages{}
		event2 := kuberneteswatcher.EventMessages{}
		data.UpdateDeploymentEvents(event1)
		data.UpdateDeploymentEvents(event2)

		if len(data.Events) != 2 {
			t.Fatalf("unexpected deployment event count, got %d expected %d", len(data.Events), 2)

		}

	})

	t.Run("init_replicaset", func(t *testing.T) {
		data.InitReplicaset("replica")
		data.InitReplicaset("replica")

		if len(data.Replicaset) != 1 {
			t.Fatalf("unexpected replicaset count, got %d expected %d", len(data.Replicaset), 1)
		}

	})

	t.Run("update_replicaset_events", func(t *testing.T) {
		data.InitReplicaset("replica")
		data.UpdateReplicasetEvents("replica", kuberneteswatcher.EventMessages{})
		data.UpdateReplicasetEvents("replica", kuberneteswatcher.EventMessages{})

		if len(*data.Replicaset["replica"].Events) != 2 {
			t.Fatalf("unexpected replicaset event count, got %d expected %d", len(*data.Replicaset["replica"].Events), 2)
		}
		err := data.UpdateReplicasetEvents("replica1", kuberneteswatcher.EventMessages{})
		if err == nil {
			t.Fatalf("expected error when trying to set event to none replicaset")
		}

	})

	t.Run("update_replicaset_status", func(t *testing.T) {

		data.InitReplicaset("replica")
		data.UpdateReplicasetStatus("replica", appsV1.ReplicaSetStatus{Replicas: 1})

		if data.Replicaset["replica"].Status.Replicas != 1 {
			t.Fatalf("unexpected replicaset status, got %d expected %d", data.Replicaset["replica"].Status.Replicas, 1)
		}
		err := data.UpdateReplicasetStatus("replica1", appsV1.ReplicaSetStatus{})
		if err == nil {
			t.Fatalf("expected error when trying to set event to none replicaset")
		}

	})

	t.Run("new_pod", func(t *testing.T) {
		pod := &v1.Pod{
			ObjectMeta: metaV1.ObjectMeta{
				Name: "pod1",
			},
		}
		err := data.NewPod(pod)
		if err != nil {
			t.Fatalf("expected error when trying to set new pod to registry")
		}

		err = data.NewPod(pod)
		if err == nil {
			t.Fatalf("expected error when trying to set existing pod to registry")
		}

	})

	t.Run("update_pod", func(t *testing.T) {

		pod := &v1.Pod{
			ObjectMeta: metaV1.ObjectMeta{
				Name: "pod1",
			},
			Status: v1.PodStatus{
				Phase: v1.PodPending,
			},
		}
		data.NewPod(pod)
		data.UpdatePod(pod, string(v1.PodSucceeded))

		if *data.Pods["pod1"].Phase != string(v1.PodSucceeded) {
			t.Fatalf("unexpected pod Phase, got %s expected %s", *data.Pods["pod1"].Phase, v1.PodSucceeded)
		}

	})

	t.Run("update_pod_events", func(t *testing.T) {
		pod := &v1.Pod{
			ObjectMeta: metaV1.ObjectMeta{
				Name: "pod1",
			},
		}
		data.NewPod(pod)

		eventTime := time.Now().Unix()
		data.UpdatePodEvents("pod1", "", kuberneteswatcher.EventMessages{
			Message: "Message",
			Time:    eventTime,
		})
		data.UpdatePodEvents("pod1", "", kuberneteswatcher.EventMessages{
			Message: "Message",
			Time:    eventTime,
		})
		data.UpdatePodEvents("pod1", "", kuberneteswatcher.EventMessages{
			Message: "Message2",
			Time:    eventTime,
		})

		if len(*data.Pods["pod1"].Events) != 2 {
			t.Fatalf("unexpected pod event count, got %d expected %d", len(*data.Pods["pod1"].Events), 2)
		}
	})

}

func TestSave(t *testing.T) {
	registry, storage := NewRegistryMock()

	annotations := map[string]string{
		"statusbay.io/report-deploy-by":      "foo@example.com",
		"statusbay.io/report-slack-channels": "#channel",
	}
	registryRow := registry.NewApplication("nginx", "default", annotations, common.ApplyStatusRunning)

	apply := kuberneteswatcher.ApplyEvent{
		Event:        "create",
		ApplyName:    "nginx-deployment",
		ResourceName: "resourceName",
		Namespace:    "default",
		Kind:         "deployment",
		Hash:         1234,
		Annotations:  annotations,
		Labels:       map[string]string{},
	}

	createMockDeploymentData(registry, registryRow, apply, "10m")

	apply2 := kuberneteswatcher.ApplyEvent{
		Event:        "create",
		ApplyName:    "nginx-deployment1",
		ResourceName: "resourceName2",
		Namespace:    "default",
		Kind:         "deployment",
		Hash:         1234,
		Annotations:  annotations,
		Labels:       map[string]string{},
	}
	createMockDeploymentData(registry, registryRow, apply2, "10m")

	time.Sleep(time.Second * 5)
	id := "1"

	t.Run("save_new_deployment", func(t *testing.T) {
		testCases := []struct {
			description string
			mutate      func(schema kuberneteswatcher.DBSchema) interface{}
			expected    interface{}
		}{
			{
				"application name",
				func(schema kuberneteswatcher.DBSchema) interface{} { return schema.Application },
				"nginx",
			},
			{
				"deployments count",
				func(schema kuberneteswatcher.DBSchema) interface{} { return len(schema.Resources.Deployments) },
				2,
			},
		}

		for _, tc := range testCases {
			value := tc.mutate(storage.MockWriteDeployment[id].Schema)
			if value != tc.expected {
				t.Fatalf("unexpected %s, got %s expected %s", tc.description, value, tc.expected)
			}
		}

	})

}

func TestDeploymentFinishSuccessful(t *testing.T) {

	registry, storage := NewRegistryMock()

	annotations := map[string]string{
		"statusbay.io/report-deploy-by":      "foo@example.com",
		"statusbay.io/report-slack-channels": "#channel",
	}
	registryRow := registry.NewApplication("nginx", "default", annotations, common.ApplyStatusRunning)

	replicasetStatus := appsV1.ReplicaSetStatus{
		Replicas:      3,
		ReadyReplicas: 3,
	}

	apply := kuberneteswatcher.ApplyEvent{
		Event:        "create",
		ApplyName:    "nginx-deployment",
		ResourceName: "resourceName",
		Namespace:    "default",
		Kind:         "deployment",
		Hash:         1234,
		Annotations:  annotations,
		Labels:       map[string]string{},
	}

	data := createMockDeploymentData(registry, registryRow, apply, "10m")

	data.InitReplicaset("replicaset-name")

	data.UpdateReplicasetStatus("replicaset-name", replicasetStatus)

	time.Sleep(time.Second * 10)
	if storage.MockWriteDeployment["1"].Status != common.ApplySuccessful {
		t.Errorf("unexpected deployment status, got %s expected %s", storage.MockWriteDeployment["1"].Status, common.ApplySuccessful)
	}
}
func TestDeploymentFinishProgressDeadLine(t *testing.T) {

	registry, storage := NewRegistryMock()

	registryRow := registry.NewApplication("nginx", "default", map[string]string{}, common.ApplyStatusRunning)

	replicasetStatus := appsV1.ReplicaSetStatus{
		Replicas:      3,
		ReadyReplicas: 2,
	}
	apply := kuberneteswatcher.ApplyEvent{
		Event:        "create",
		ApplyName:    "nginx-deployment",
		ResourceName: "resourceName",
		Namespace:    "default",
		Kind:         "deployment",
		Hash:         1234,
		Annotations:  map[string]string{},
		Labels:       map[string]string{},
	}
	data := createMockDeploymentData(registry, registryRow, apply, "1s")

	data.InitReplicaset("replicaset-name")

	data.UpdateReplicasetStatus("replicaset-name", replicasetStatus)

	time.Sleep(time.Second * 8)
	if storage.MockWriteDeployment["1"].Status != common.ApplyStatusFailed {
		t.Fatalf("unexpected deployment status, got %s expected %s", storage.MockWriteDeployment["1"].Status, common.ApplyStatusFailed)
	}

	if storage.MockWriteDeployment["1"].Schema.DeploymentDescription != common.ApplyStatusDescriptionProgressDeadline {
		t.Fatalf("unexpected deployment message description, got %s expected %s", storage.MockWriteDeployment["1"].Schema.DeploymentDescription, common.ApplyStatusDescriptionProgressDeadline)
	}
}

func TestGetApplyID(t *testing.T) {

	registry, _ := NewRegistryMock()

	fakeDeployment := GetFakeDeployment(1)
	registryRow := registry.NewApplication("nginx", fakeDeployment.GetNamespace(), fakeDeployment.GetAnnotations(), common.ApplyStatusRunning)
	registryRow.DBSchema.CreationTimestamp = 12345
	applyID := registryRow.GetApplyID()

	if applyID != "8fe4325f717a39f8bdcf772cc81e201102851fb8" {
		t.Fatalf("unexpected apply ID, got %s expected %s", applyID, "8fe4325f717a39f8bdcf772cc81e201102851fb8")

	}
}
