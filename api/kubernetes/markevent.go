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

		pod(deployment.Pods, eventMarksConfig)

		service(deployment.Service, eventMarksConfig)

	}

	for _, daemonset := range appDeployment.Resources.Daemonsets {

		for i, dep := range daemonset.Events {
			eventDescription := eventmark.MarkEvent(dep.Message, eventMarksConfig.Demonset)
			daemonset.Events[i].MarkDescriptions = eventDescription
		}

		pod(daemonset.Pods, eventMarksConfig)

		service(daemonset.Service, eventMarksConfig)

	}

	for _, statefulset := range appDeployment.Resources.Statefulsets {

		for i, dep := range statefulset.Events {
			eventDescription := eventmark.MarkEvent(dep.Message, eventMarksConfig.Statefulset)
			statefulset.Events[i].MarkDescriptions = eventDescription
		}

		pod(statefulset.Pods, eventMarksConfig)

		service(statefulset.Service, eventMarksConfig)

	}
}

func pod(pods map[string]ResponseDeploymenPod, events config.KubernetesMarksEvents) {

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

func service(service map[string]ResponseServicesData, events config.KubernetesMarksEvents) {

	for _, services := range service {
		for i, svc := range services.Events {
			eventDescription := eventmark.MarkEvent(svc.Message, events.Service)
			services.Events[i].MarkDescriptions = eventDescription
		}
	}

}
