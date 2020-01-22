package kuberneteswatcher

import (
	"context"
	"time"

	appsV1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"

	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type RegistryData interface {
	UpdatePodEvents(podName string, event EventMessages) error
	UpdatePod(pod *v1.Pod, status string) error
	NewPod(pod *v1.Pod) error
	GetName() string
}

// WatchData struct
type WatchData struct {
	ListOptions  metaV1.ListOptions
	RegistryData RegistryData
	Namespace    string
	Ctx          context.Context
}

type MessageDeploy struct {
	UID               string            `json:"ID"`
	Kind              string            `json:"Kind"`
	Namespace         string            `json:"Namespace"`
	Name              string            `json:"Name"`
	CreationTimestamp time.Time         `json:"CreationTimestamp"`
	Labels            map[string]string `json:"Labels"`
}

// Deployment struct  TODO ::
type MetaData struct {
	Name         string            `json:"Name"`
	Namespace    string            `json:"Namespace"`
	ClusterName  string            `json:"ClusterName"`
	SpecHash     uint64            `json:"SpecHash"`
	Labels       map[string]string `json:"Labels"`
	DesiredState int32             `json:"DesiredState"`
}

// DeploymenPod struct  TODO ::
type DeploymenPod struct {
	Phase             *string          `json:"Phase"`
	CreationTimestamp time.Time        `json:"CreationTimestamp"`
	Events            *[]EventMessages `json:"Events"`
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
	DeploymentEvents        []EventMessages         `json:"DeploymentEvents"`
	Replicaset              map[string]Replicaset   `json:"Replicaset"`
	Pods                    map[string]DeploymenPod `json:"Pods"`
	ProgressDeadlineSeconds int64
}

// Daemonset Data
type DaemonsetData struct {
	Metadata                MetaData                `json:"MetaData"`
	Status                  appsV1.DaemonSetStatus  `json:"Status"`
	DaemonsetEvents         []EventMessages         `json:"DaemonsetEvents"`
	Pods                    map[string]DeploymenPod `json:"Pods"`
	ProgressDeadlineSeconds int64
}
