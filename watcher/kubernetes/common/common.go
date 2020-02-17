package common

import (
	log "github.com/sirupsen/logrus"
	eventwatch "k8s.io/apimachinery/pkg/watch"
)

//DeploymentStatus defined the status of the deployment
type DeploymentStatus string

const (
	// DeploymentSuccessful when deployment finish successfully
	DeploymentSuccessful DeploymentStatus = "successful"

	// DeploymentStatusFailed when deployment failed
	DeploymentStatusFailed DeploymentStatus = "failed"

	// DeploymentStatusRunning when deployment still in progress
	DeploymentStatusRunning DeploymentStatus = "running"

	// DeploymentStatusDeleted when deployment deleted
	DeploymentStatusDeleted DeploymentStatus = "deleted"

	// DeploymentCanceled when statusbay stop watch
	DeploymentCanceled DeploymentStatus = "cancelled"
)

// DeploymentReport defined deployment reporter message
type DeploymentReport struct {
	// To is a  list of channels/username to send message to
	To []string

	// DeployBy owner
	DeployBy string

	// Name of the apply
	Name string

	// Status of the apply
	Status DeploymentStatus

	// Deployment URI
	URI string

	// LogEntry is the application logger
	LogEntry log.Entry

	// ClusterName of the apply
	ClusterName string
}

func IsSupportedEventType(eventType eventwatch.EventType) bool {
	return (eventType == eventwatch.Modified || eventType == eventwatch.Added || eventType == eventwatch.Deleted)
}
