package kubernetes

import (
	"reflect"
	"statusbay/config"
	"testing"
)

func TestMarkApplicationDeploymentEvents(t *testing.T) {

	// events list
	searchPodEvents := []config.EventMarksConfig{
		{Pattern: "OOMKilled", Descriptions: []string{"pod event1"}},
		{Pattern: "CrashLoopBackOff", Descriptions: []string{"pod event2"}},
	}
	searchReplicasetEvents := []config.EventMarksConfig{
		{Pattern: "Created", Descriptions: []string{"rc event1"}},
	}
	searchDeploymentEvents := []config.EventMarksConfig{
		{Pattern: "Scaled", Descriptions: []string{"deployment event1"}},
	}
	searchPvcEvents := []config.EventMarksConfig{
		{Pattern: "not found", Descriptions: []string{"not found pvc1"}},
	}
	searchServiceEvents := []config.EventMarksConfig{
		{Pattern: "service error", Descriptions: []string{"not found pvc1"}},
	}
	searchDaemonsetEvents := []config.EventMarksConfig{
		{Pattern: "demonset error", Descriptions: []string{"error 1"}},
	}
	searchStatefulsetEvents := []config.EventMarksConfig{
		{Pattern: "statefulset error", Descriptions: []string{"error 1"}},
	}

	// messages content
	podEvent := map[string]ResponseDeploymenPod{
		"pod": {
			PVC: map[string][]ResponseEventMessages{
				"pvc": {
					{Message: "not found"},
					{Message: "not found"},
				},
			},
			Events: []ResponseEventMessages{
				{Message: "OOMKilled"},
				{Message: "CrashLoopBackOff"},
				{Message: "foo"},
			},
		},
	}

	service := map[string]ResponseServicesData{
		"service": {
			Events: []ResponseEventMessages{
				{Message: "service error"},
			},
		},
	}

	applyData := ResponseDeploymentData{
		Resources: ResponseResourcesData{
			Statefulsets: map[string]StatefulsetDataResponse{
				"statefulset": {
					Events: []ResponseEventMessages{
						{Message: "statefulset error"},
					},
					Pods:     podEvent,
					Services: service,
				},
			},
			Daemonsets: map[string]DaemonsetDataResponse{
				"daemonset": {
					Events: []ResponseEventMessages{
						{Message: "demonset error"},
					},
					Pods:     podEvent,
					Services: service,
				},
			},
			Deployments: map[string]DeploymentDataResponse{
				"deployment": {
					Events: []ResponseEventMessages{
						{Message: "Scaled"},
						{Message: "foo"},
					},
					Pods: podEvent,
					Replicaset: map[string]ResponseReplicaset{
						"rs": {
							Events: []ResponseEventMessages{
								{Message: "Created"},
								{Message: "foo"},
							},
						},
					},
					Services: service,
				},
			},
		},
	}

	eventsConfig := config.KubernetesMarksEvents{
		Pod:         searchPodEvents,
		Replicaset:  searchReplicasetEvents,
		Deployment:  searchDeploymentEvents,
		Pvc:         searchPvcEvents,
		Service:     searchServiceEvents,
		Demonset:    searchDaemonsetEvents,
		Statefulset: searchStatefulsetEvents,
	}
	MarkApplicationDeploymentEvents(&applyData, eventsConfig)

	testCases := []struct {
		test          string
		mutate        func(c ResponseDeploymentData) []ResponseEventMessages
		expectedCount int
	}{
		{
			"deployment",
			func(d ResponseDeploymentData) []ResponseEventMessages {
				return d.Resources.Deployments["deployment"].Events
			},
			1,
		},
		{
			"deployment pod",
			func(d ResponseDeploymentData) []ResponseEventMessages {
				return d.Resources.Deployments["deployment"].Pods["pod"].Events
			},
			2,
		},
		{
			"deployment replicaset",
			func(d ResponseDeploymentData) []ResponseEventMessages {
				return d.Resources.Deployments["deployment"].Replicaset["rs"].Events
			},
			1,
		},
		{
			"deployment pvc",
			func(d ResponseDeploymentData) []ResponseEventMessages {
				return d.Resources.Deployments["deployment"].Pods["pod"].PVC["pvc"]
			},
			2,
		},
		{
			"deployment service",
			func(d ResponseDeploymentData) []ResponseEventMessages {
				return d.Resources.Deployments["deployment"].Services["service"].Events
			},
			1,
		},
		{
			"daemonset",
			func(d ResponseDeploymentData) []ResponseEventMessages {
				return d.Resources.Daemonsets["daemonset"].Events
			},
			1,
		},
		{
			"daemonset pod",
			func(d ResponseDeploymentData) []ResponseEventMessages {
				return d.Resources.Daemonsets["daemonset"].Pods["pod"].Events
			},
			2,
		},
		{
			"daemonset service",
			func(d ResponseDeploymentData) []ResponseEventMessages {
				return d.Resources.Daemonsets["daemonset"].Services["service"].Events
			},
			1,
		},
		{
			"daemonset pvc",
			func(d ResponseDeploymentData) []ResponseEventMessages {
				return d.Resources.Daemonsets["daemonset"].Pods["pod"].PVC["pvc"]
			},
			2,
		},
		{
			"statefulset",
			func(d ResponseDeploymentData) []ResponseEventMessages {
				return d.Resources.Statefulsets["statefulset"].Events
			},
			1,
		},
		{
			"statefulset pod",
			func(d ResponseDeploymentData) []ResponseEventMessages {
				return d.Resources.Statefulsets["statefulset"].Pods["pod"].Events
			},
			2,
		},
		{
			"statefulset service",
			func(d ResponseDeploymentData) []ResponseEventMessages {
				return d.Resources.Statefulsets["statefulset"].Services["service"].Events
			},
			1,
		},
		{
			"statefulset pvc",
			func(d ResponseDeploymentData) []ResponseEventMessages {
				return d.Resources.Statefulsets["statefulset"].Pods["pod"].PVC["pvc"]
			},
			2,
		},
	}

	for _, ct := range testCases {
		t.Run(ct.test, func(t *testing.T) {
			resp := ct.mutate(applyData)
			markFound := 0
			for _, e := range resp {
				if len(e.MarkDescriptions) > 0 {
					markFound = markFound + 1
				}
			}
			if markFound != ct.expectedCount {
				t.Fatalf("unexpected mark message count, got %d expected %d", markFound, ct.expectedCount)
			}

		})
	}

}

func TestMarkApplicationDeploymentEventContent(t *testing.T) {

	expectedEventDescriptions := []string{"deployment event1"}

	appDeployment := ResponseDeploymentData{
		Resources: ResponseResourcesData{
			Deployments: map[string]DeploymentDataResponse{
				"deployment": {
					Events: []ResponseEventMessages{
						{Message: "Scaled"},
						{Message: "foo"},
					},
				},
			},
		},
	}

	eventsConfig := config.KubernetesMarksEvents{
		Deployment: []config.EventMarksConfig{
			{Pattern: "Scaled", Descriptions: expectedEventDescriptions},
		},
	}

	MarkApplicationDeploymentEvents(&appDeployment, eventsConfig)

	if !reflect.DeepEqual(appDeployment.Resources.Deployments["deployment"].Events[0].MarkDescriptions, expectedEventDescriptions) {
		t.Fatalf("unexpected mark message count, got %v expected %v", appDeployment.Resources.Deployments["deployment"].Events[0].MarkDescriptions, expectedEventDescriptions)
	}

}
