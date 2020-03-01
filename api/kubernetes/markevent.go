package kubernetes

import (
	"statusbay/api/eventmark"
	"statusbay/config"
)

// MarkApplicationDeploymentEvents returns list of markes events from configuration
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

		pods(deployment.Pods, eventMarksConfig)

		services(deployment.Services, eventMarksConfig)

	}

	for _, daemonset := range appDeployment.Resources.Daemonsets {

		for i, dep := range daemonset.Events {
			eventDescription := eventmark.MarkEvent(dep.Message, eventMarksConfig.Demonset)
			daemonset.Events[i].MarkDescriptions = eventDescription
		}

		pods(daemonset.Pods, eventMarksConfig)

		services(daemonset.Services, eventMarksConfig)

	}

	for _, statefulset := range appDeployment.Resources.Statefulsets {

		for i, dep := range statefulset.Events {
			eventDescription := eventmark.MarkEvent(dep.Message, eventMarksConfig.Statefulset)
			statefulset.Events[i].MarkDescriptions = eventDescription
		}

		pods(statefulset.Pods, eventMarksConfig)

		services(statefulset.Services, eventMarksConfig)

	}
}

// pod will mark pod messages from pod event section
func pods(pods map[string]ResponseDeploymenPod, events config.KubernetesMarksEvents) {

	for _, pod := range pods {
		for i, event := range pod.Events {
			eventDescription := eventmark.MarkEvent(event.Message, events.Pod)
			pod.Events[i].MarkDescriptions = eventDescription
		}

		for _, pvc := range pod.PVC {
			for i, event := range pvc {
				eventDescription := eventmark.MarkEvent(event.Message, events.Pvc)
				pvc[i].MarkDescriptions = eventDescription
			}
		}
	}
}

// services will mark pod messages from service event section
func services(service map[string]ResponseServicesData, events config.KubernetesMarksEvents) {

	for _, services := range service {
		for i, svc := range services.Events {
			eventDescription := eventmark.MarkEvent(svc.Message, events.Service)
			services.Events[i].MarkDescriptions = eventDescription
		}
	}

}
