package kubernetes

import (
	"reflect"
	"statusbay/config"
	"testing"
)

func TestMarkApplicationDeploymentEvents(t *testing.T) {

	podEvents := []config.EventMarksConfig{
		{Pattern: "OOMKilled", Descriptions: []string{"pod event1"}},
		{Pattern: "CrashLoopBackOff", Descriptions: []string{"pod event2"}},
	}
	replicasetEvents := []config.EventMarksConfig{
		{Pattern: "Created", Descriptions: []string{"rc event1"}},
	}
	deploymentEvents := []config.EventMarksConfig{
		{Pattern: "Scaled", Descriptions: []string{"deployment event1"}},
	}

	appDeployment := ResponseDeploymentData{
		Resources: ResponseResourcesData{
			Deployments: map[string]DeploymentDataResponse{
				"deployment": {
					DeploymentEvents: []ResponseEventMessages{
						{Message: "Scaled"},
						{Message: "foo"},
					},
					Pods: map[string]ResponseDeploymenPod{
						"pod": {
							Events: []ResponseEventMessages{
								{Message: "OOMKilled"},
								{Message: "CrashLoopBackOff"},
								{Message: "foo"},
							},
						},
					},
					Replicaset: map[string]ResponseReplicaset{
						"rs": {
							Events: []ResponseEventMessages{
								{Message: "Created"},
								{Message: "foo"},
							},
						},
					},
				},
			},
		},
	}

	eventsConfig := config.KubernetesMarksEvents{
		Pod:        podEvents,
		Replicaset: replicasetEvents,
		Deployment: deploymentEvents,
	}
	MarkApplicationDeploymentEvents(&appDeployment, eventsConfig)

	testCases := []struct {
		test          string
		mutate        func(c ResponseDeploymentData) []ResponseEventMessages
		expectedCount int
	}{
		{
			"deployment",
			func(d ResponseDeploymentData) []ResponseEventMessages {
				return d.Resources.Deployments["deployment"].DeploymentEvents
			},
			1,
		},
		{
			"deployment",
			func(d ResponseDeploymentData) []ResponseEventMessages {
				return d.Resources.Deployments["deployment"].Pods["pod"].Events
			},
			2,
		},
		{
			"deployment",
			func(d ResponseDeploymentData) []ResponseEventMessages {
				return d.Resources.Deployments["deployment"].Replicaset["rs"].Events
			},
			1,
		},
	}

	for _, ct := range testCases {
		t.Run(ct.test, func(t *testing.T) {
			resp := ct.mutate(appDeployment)
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
					DeploymentEvents: []ResponseEventMessages{
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

	if !reflect.DeepEqual(appDeployment.Resources.Deployments["deployment"].DeploymentEvents[0].MarkDescriptions, expectedEventDescriptions) {
		t.Fatalf("unexpected mark message count, got %v expected %v", appDeployment.Resources.Deployments["deployment"].DeploymentEvents[0].MarkDescriptions, expectedEventDescriptions)
	}

}
