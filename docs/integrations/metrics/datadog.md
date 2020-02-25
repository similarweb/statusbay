# DataDog
This integration makes the following assumptions: 

* You have an active account in Datadog
* You have a dedicated User and Token for StatusBay to use Datadog's API.
  
StatusBay introduces the ability to show DataDog graphs for a specific service running on Kubernetes. 

On operator of a service in Kubernetes will have ability to see Datadog graph in StatusBay's UI based on Datadog query.


## How to enable this provider?

In order to enable this provider please proceed with the next steps:

* Configure DataDog provider via StatusBay [API configuration file](../../../examples/configuration/api.yaml#L13), you will find all the available configuration options in the example file.
* Add the [Available annotations](#available-annotations) for this provider

## Available annotations
| Name | Type | Associated Annotations | 
| ---- | ---- | ---------------------- | 
| Datadog | Metrics | `statusbay.io/metrics-datadog-<Metric-Name>: datadog-query` |



* Make sure you add these annotations to one of the following kinds: Deployment,Daemonset & Statefulset.
* The aforementioned annotation will need following: 
  * `<Metric-Name>` - Will be the name displayed  the name of the graph
  * `datadog-query` - Datadog query to get the graph
  
## The result
![Multiple Clusters](../../../ui/client/src/components/IntergationModals/MetricIntegrationModal/metrics.png)

* In the example above here are the comparison for the following
  * `<Metric-Name>` -> `2xx vs 4xx`
  * `datadog-query` -> `sw.haproxy.backend.response.2xx{*}.as_count(), sw.haproxy.backend.response.4xx{*}.as_count()`