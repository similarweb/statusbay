package kuberneteswatcher

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	appsV1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"

	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type RegistryData interface {
	UpdatePodEvents(podName string, pvcName string, event EventMessages) error
	NewPod(pod *v1.Pod) error
	UpdatePod(pod *v1.Pod, status string) error
	NewService(pod *v1.Service) error
	UpdateServiceEvents(name string, event EventMessages) error
	GetName() string
}

// WatchData struct
type WatchData struct {
	Ctx          context.Context
	ListOptions  metaV1.ListOptions
	RegistryData RegistryData
	Namespace    string
	LogEntry     log.Entry
	ApplyID      string
}

type MessageDeploy struct {
	UID               string            `json:"ID"`
	Kind              string            `json:"Kind"`
	Namespace         string            `json:"Namespace"`
	Name              string            `json:"Name"`
	CreationTimestamp time.Time         `json:"CreationTimestamp"`
	Labels            map[string]string `json:"Labels"`
}

// MetaData struct  TODO ::
type MetaData struct {
	Name         string            `json:"Name"`
	Namespace    string            `json:"Namespace"`
	ClusterName  string            `json:"ClusterName"`
	Labels       map[string]string `json:"Labels"`
	Annotations  map[string]string `json:"Annotations"`
	Metrics      []Metrics         `json:"Metrics"`
	Alerts       []Alerts          `json:"Alerts"`
	DesiredState int32             `json:"DesiredState"`
}

// DeploymenPod struct  TODO ::
type DeploymenPod struct {
	Phase             *string                    `json:"Phase"`
	CreationTimestamp time.Time                  `json:"CreationTimestamp"`
	Events            *[]EventMessages           `json:"Events"`
	Pvcs              map[string][]EventMessages `json:"Pvcs"`
}

// EventMessages struct  TODO ::
type EventMessages struct {
	Message             string `json:"Message"`
	Time                int64  `json:"Time"`
	Action              string `json:"Action"`
	ReportingController string `json:"ReportingController"`
}

// Replicaset struct  TODO ::
type Replicaset struct {
	Events *[]EventMessages         `json:"Events"`
	Status *appsV1.ReplicaSetStatus `json:"Status"`
}

// DeploymentData struct TODO
type DeploymentData struct {
	Deployment              MetaData                `json:"MetaData"`
	Status                  appsV1.DeploymentStatus `json:"Status"`
	Events                  []EventMessages         `json:"Events"`
	Replicaset              map[string]Replicaset   `json:"Replicaset"`
	Pods                    map[string]DeploymenPod `json:"Pods"`
	Services                map[string]ServicesData `json:"Services"`
	ProgressDeadlineSeconds int64
}

// DaemonsetData
type DaemonsetData struct {
	Metadata                MetaData                `json:"MetaData"`
	Status                  appsV1.DaemonSetStatus  `json:"Status"`
	Events                  []EventMessages         `json:"Events"`
	Pods                    map[string]DeploymenPod `json:"Pods"`
	Services                map[string]ServicesData `json:"Services"`
	ProgressDeadlineSeconds int64
}

// StatefulsetData holds the data of Statefulset for the registry
type StatefulsetData struct {
	Statefulset             MetaData                 `json:"MetaData"`
	Status                  appsV1.StatefulSetStatus `json:"Status"`
	Events                  []EventMessages          `json:"Events"`
	Pods                    map[string]DeploymenPod  `json:"Pods"`
	Services                map[string]ServicesData  `json:"Services"`
	ProgressDeadlineSeconds int64
}

// ServicesData holds the data of services
type ServicesData struct {
	Events *[]EventMessages `json:"Events"`
}

//Metrics describe the metrics data integration
type Metrics struct {
	Name     string `json:"Name"`
	Provider string `json:"Provider"`
	Query    string `json:"Query"`
}

type Alerts struct {
	Provider string `json:"Provider"`
	Tags     string `json:"Tags"`
}
