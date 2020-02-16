package kubernetes

import (
	"statusbay/api/eventmark"
	"statusbay/config"
)

// MarkApplicationDeploymentEvents returns list of markes events from given configuration (See file /events.yaml)
// This is a ugly/bad and fast implementation, we must refactor this code. I have open a github issue for that https://github.com/similarweb/statusbay/issues/72
func MarkApplicationDeploymentEvents(appDeployment *ResponseDeploymentData, eventMarksConfig config.KubernetesMarksEvents) {

	for _, deployment := range appDeployment.Resources.Deployments {

		for i, dep := range deployment.Events {
			eventDescription := eventmark.MarkEvent(dep.Message, eventMarksConfig.Deployment)
			deployment.Events[i].MarkDescriptions = eventDescription
		}

		for _, rs := range deployment.Replicaset {
			for i, event := range rs.Events {
				eventDescription := eventmark.MarkEvent(event.Message, eventMarksConfig.Replicaset)
				rs.Events[i].MarkDescriptions = eventDescription
			}
		}

		for _, pod := range deployment.Pods {
			for i, event := range pod.Events {
				eventDescription := eventmark.MarkEvent(event.Message, eventMarksConfig.Pod)
				pod.Events[i].MarkDescriptions = eventDescription
			}
		}
	}

	for _, daemonset := range appDeployment.Resources.Daemonsets {

		for i, dep := range daemonset.Events {
			eventDescription := eventmark.MarkEvent(dep.Message, eventMarksConfig.Demonset)
			daemonset.Events[i].MarkDescriptions = eventDescription
		}

		for _, pod := range daemonset.Pods {
			for i, event := range pod.Events {
				eventDescription := eventmark.MarkEvent(event.Message, eventMarksConfig.Pod)
				pod.Events[i].MarkDescriptions = eventDescription
			}
		}

	}

	for _, statefulset := range appDeployment.Resources.Statefulsets {

		for i, dep := range statefulset.Events {
			eventDescription := eventmark.MarkEvent(dep.Message, eventMarksConfig.Statefulset)
			statefulset.Events[i].MarkDescriptions = eventDescription
		}

		for _, pod := range statefulset.Pods {
			for i, event := range pod.Events {
				eventDescription := eventmark.MarkEvent(event.Message, eventMarksConfig.Pod)
				pod.Events[i].MarkDescriptions = eventDescription
			}
		}
	}
}
