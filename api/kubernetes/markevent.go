package kubernetes

import (
	"statusbay/api/eventmark"
	"statusbay/config"
)

func MarkApplicationDeploymentEvents(appDeployment *ResponseDeploymentData, eventMarksConfig config.KubernetesMarksEvents) {

	for _, application := range appDeployment.Resources.Deployments {

		for i, dep := range application.DeploymentEvents {
			eventDescription := eventmark.MarkEvent(dep.Message, eventMarksConfig.Deployment)
			application.DeploymentEvents[i].MarkDescriptions = eventDescription
		}

		for _, rs := range application.Replicaset {
			for i, event := range rs.Events {
				eventDescription := eventmark.MarkEvent(event.Message, eventMarksConfig.Replicaset)
				rs.Events[i].MarkDescriptions = eventDescription
			}
		}

		for _, pod := range application.Pods {
			for i, event := range pod.Events {
				eventDescription := eventmark.MarkEvent(event.Message, eventMarksConfig.Pod)
				pod.Events[i].MarkDescriptions = eventDescription
			}
		}
	}
}
