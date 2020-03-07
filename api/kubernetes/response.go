package kubernetes

import (
	"time"
)

// START Kubernetes deployment response

type ResponseKubernetesApplicationsCount struct {
	Records []ResponseKubernetesApplications `json:"Records"`
	Count   int64                            `json:"Count"`
}

type ResponseKubernetesApplications struct {
	ApplyID   string `json:"ApplyID"`
	Name      string `json:"Name"`
	Status    string `json:"Status"`
	Cluster   string `json:"Cluster"`
	Namespace string `json:"Namespace"`
	DeployBy  string `json:"DeployBy"`
	Time      int64  `json:"Time"`
}

type ResponseMetaData struct {
	Name         string            `json:"Name"`
	Namespace    string            `json:"Namespace"`
	ClusterName  string            `json:"ClusterName"`
	Labels       map[string]string `json:"Labels"`
	DesiredState int32             `json:"DesiredState"`
	Alerts       []ResponseAlerts  `json:"Alerts"`
	Metrics      []ResponseMetrics `json:"Metrics"`
}

type ResponseDeploymenPod struct {
	Phase             *string                            `json:"Phase"`
	CreationTimestamp time.Time                          `json:"CreationTimestamp"`
	Events            []ResponseEventMessages            `json:"Events"`
	PVC               map[string][]ResponseEventMessages `json:"Pvcs"`
	Containers        map[string]Container               `json:"Containers"`
}

type ResponseEventMessages struct {
	Message             string   `json:"Message"`
	Time                int64    `json:"Time"`
	Action              string   `json:"Action"`
	ReportingController string   `json:"ReportingController"`
	MarkDescriptions    []string `json:"MarkDescriptions"`
}

type ResponseReplicaset struct {
	Events []ResponseEventMessages  `json:"Events"`
	Status ResponseDeploymentStatus `json:"Status"`
}

type ResponseServicesData struct {
	Events []ResponseEventMessages `json:"Events"`
}
type ResponseMetricsQuery struct {
	Query    string `json:"Query"`
	Title    string `json:"Title"`
	SubTitle string `json:"SubTitle"`
}

type ResponseCondition struct {
	Type               string    `json:"Type"`
	Status             string    `json:"Status"`
	LastTransitionTime time.Time `json:"LastTransitionTime"`
	Reason             string    `json:"Reason"`
	Message            string    `json:"Message"`
}
type ResponseDeploymentStatus struct {
	ObservedGeneration  int64 `json:"ObservedGeneration"`
	Replicas            int32 `json:"Replicas"`
	UpdatedReplicas     int32 `json:"UpdatedReplicas"`
	ReadyReplicas       int32 `json:"ReadyReplicas"`
	AvailableReplicas   int32 `json:"AvailableReplicas"`
	UnavailableReplicas int32 `json:"UnavailableReplicas"`
	Conditions          []ResponseCondition
}

type DeploymentDataResponse struct {
	Deployment ResponseMetaData                `json:"MetaData"`
	Events     []ResponseEventMessages         `json:"Events"`
	Metrics    []ResponseMetricsQuery          `json:"Metrics"`
	Pods       map[string]ResponseDeploymenPod `json:"Pods"`
	Replicaset map[string]ResponseReplicaset   `json:"Replicaset"`
	Status     ResponseDeploymentStatus        `json:"Status"`
	Services   map[string]ResponseServicesData `json:"Services"`
}

type DaemonsetDataResponse struct {
	Metadata ResponseMetaData                `json:"MetaData"`
	Events   []ResponseEventMessages         `json:"Events"`
	Pods     map[string]ResponseDeploymenPod `json:"Pods"`
	Status   ResponseDeploymentStatus        `json:"Status"`
	Services map[string]ResponseServicesData `json:"Services"`
}

type StatefulsetDataResponse struct {
	Statefulset ResponseMetaData                `json:"MetaData"`
	Events      []ResponseEventMessages         `json:"Events"`
	Pods        map[string]ResponseDeploymenPod `json:"Pods"`
	Status      ResponseDeploymentStatus        `json:"Status"`
	Services    map[string]ResponseServicesData `json:"Services"`
}

type ResponseResourcesData struct {
	Deployments  map[string]DeploymentDataResponse  `json:"Deployments"`
	Daemonsets   map[string]DaemonsetDataResponse   `json:"Daemonsets"`
	Statefulsets map[string]StatefulsetDataResponse `json:"Statefulsets"`
}

type ResponseDeploymentData struct {
	Resources ResponseResourcesData `json:"Resources"`
}

type ResponseKubernetesDeployment struct {
	Name      string                 `json:"Name"`
	DeployBy  string                 `json:"DeployBy"`
	Cluster   string                 `json:"Cluster"`
	Namespace string                 `json:"Namespace"`
	Status    string                 `json:"Status"`
	Time      int64                  `json:"Time"`
	Details   ResponseDeploymentData `json:"Details"`
}

type ResponseAlerts struct {
	Tags     string `json:"Tags"`
	Provider string `json:"Provider"`
}

type ResponseMetrics struct {
	Name     string `json:"Name"`
	Query    string `json:"Query"`
	Provider string `json:"Provider"`
}

// END Kubernetes deployment response

type PeriodsResponse struct {
	Status      string
	Start       string
	StartUnix   int64
	End         string
	EndUnix     int64
	Description string
}

//DeploymentDatadogLink represents datadog section in html
type DeploymentDatadogLink struct {
	Group   string `json:"Group"`
	Task    string `json:"Task"`
	Service string `json:"Service"`
	URL     string `json:"Url"`
}

type DeploymentsResponse struct {
	Name       string `json:"Name"`
	DeployTime int64  `json:"DeployTime"`
	Status     string `json:"Status"`
}

type DeploymentListResponse struct {
	Job string `json:"Job"`
}

// Container holds the data of container
type Container struct {
	Logs *[]string `json:"Logs"`
}
