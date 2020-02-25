# Pingdom
This integration makes the following assumptions: 

* You have an active account in Pingdom
* You have a dedicated User and Token for StatusBay to use Pingdom's API.
* You are using tags for all your alerts in Pingdom
  
StatusBay introduces the ability to show Pingdom alerts for a specific service running on Kubernetes. 

On operator of a service in Kubernetes will have ability to see alerts graph in StatusBay's UI based on Pingdom tags.


## How to enable this provider?

In order to enable this provider please proceed with the next steps:

* Configure Pingdom provider via StatusBay [API configuration file](../../../examples/configuration/api.yaml#L25), you will find all the available configuration options in the example file.
* Add the [Available annotations](#available-annotations) for this provider

## Available annotations
| Name | Type | Associated Annotations | 
| ---- | ---- | ---------------------- | 
| Pingdom | Endpoints Alerts | `statusbay.io/alerts-pingdom: tag1,tag2` |

* Make sure you add these annotations to one of the following kinds: Deployment,Daemonset & Statefulset.
* The annotations support comma separated list of tags of Alerts which we want to display. 

## The result
![Multiple Clusters](../../../ui/client/src/components/IntergationModals/AlertsIntegrationModal/alerts.png)