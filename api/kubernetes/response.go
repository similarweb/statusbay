package kubernetes

import "time"

// START Kubernetes deployment response

type ResponseKubernetesApplicationsCount struct {
	Records []ResponseKubernetesApplications `json:"Records"`
	Count   int64                            `json:"Count"`
}

type ResponseKubernetesApplications struct {
	Name      string `json:"Name"`
	Status    string `json:"Status"`
	Cluster   string `json:"Cluster"`
	Namespace string `json:"Namespace"`
	DeployBy  string `json:"DeployBy"`
	Time      int64  `json:"Time"`
}
type ResponseMessageDeploy struct {
	UID               string            `json:"ID"`
	Kind              string            `json:"Kind"`
	Namespace         string            `json:"Namespace"`
	Name              string            `json:"Name"`
	CreationTimestamp time.Time         `json:"CreationTimestamp"`
	Labels            map[string]string `json:"Labels"`
}
type ResponseMetaData struct {
	Name         string            `json:"Name"`
	Namespace    string            `json:"Namespace"`
	ClusterName  string            `json:"ClusterName"`
	SpecHash     uint64            `json:"SpecHash"`
	Labels       map[string]string `json:"Labels"`
	DesiredState int32             `json:"DesiredState"`
}

type DeploymentDataResponse struct {
	Deployment       ResponseMetaData                `json:"MetaData"`
	DeploymentEvents []ResponseEventMessages         `json:"DeploymentEvents"`
	Metrics          []ResponseMetricsQuery          `json:"Metrics"`
	Pods             map[string]ResponseDeploymenPod `json:"Pods"`
	Replicaset       map[string]ResponseReplicaset   `json:"Replicaset"`
}

type ResponseDeploymenPod struct {
	Phase             *string                 `json:"Phase"`
	CreationTimestamp time.Time               `json:"CreationTimestamp"`
	Events            []ResponseEventMessages `json:"Events"`
}

type ResponseEventMessages struct {
	Message             string   `json:"Message"`
	Time                int64    `json:"Time"`
	Action              string   `json:"Action"`
	ReportingController string   `json:"ReportingController"`
	MarkDescriptions    []string `json:"MarkDescriptions"`
}

type ResponseReplicaset struct {
	Events []ResponseEventMessages `json:"Events"`
}

type ResponseMetricsQuery struct {
	Query    string `json:"Query"`
	Title    string `json:"Title"`
	SubTitle string `json:"SubTitle"`
}

type DaemonsetDataResponse struct {
	Metadata ResponseMetaData                `json:"MetaData"`
	Events   []ResponseEventMessages         `json:"DaemonsetEvents"`
	Pods     map[string]ResponseDeploymenPod `json:"Pods"`
}

type ResponseResourcesData struct {
	Deployment map[string]DeploymentDataResponse `json:"Deployments"`
	Daemonsets map[string]DaemonsetDataResponse  `json:"Daemonsets"`
}

type ResponseDeploymentData struct {
	Resources ResponseResourcesData `json:"Resources"`
}

type ResponseKubernetesDeployment struct {
	Name      string                 `json:"Name"`
	Cluster   string                 `json:"Cluster"`
	Namespace string                 `json:"Namespace"`
	Status    string                 `json:"Status"`
	Time      string                 `json:"Time"`
	Details   ResponseDeploymentData `json:"Details"`
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
