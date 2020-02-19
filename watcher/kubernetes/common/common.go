package common

import (
	log "github.com/sirupsen/logrus"
	eventwatch "k8s.io/apimachinery/pkg/watch"
)

//DeploymentStatus defined the status of the deployment
type DeploymentStatus string

const (
	// ApplySuccessful when deployment finish successfully
	ApplySuccessful DeploymentStatus = "successful"

	// ApplyStatusFailed when deployment failed
	ApplyStatusFailed DeploymentStatus = "failed"

	// ApplyStatusRunning when deployment still in progress
	ApplyStatusRunning DeploymentStatus = "running"

	// ApplyStatusDeleted when deployment deleted
	ApplyStatusDeleted DeploymentStatus = "deleted"

	// ApplyCanceled when statusbay stop watch
	ApplyCanceled DeploymentStatus = "cancelled"
)

// DeploymentStatusDescription are the various descriptions of the states a deployment can be in.
type DeploymentStatusDescription string

const (

	// ApplyStatusDescriptionRunning running deployment
	ApplyStatusDescriptionRunning DeploymentStatusDescription = "Deployment is running"

	// ApplyStatusDescriptionSuccessful successfully deployment
	ApplyStatusDescriptionSuccessful DeploymentStatusDescription = "Deployment completed successfully"

	// ApplyStatusDescriptionProgressDeadline progress deadline ended
	ApplyStatusDescriptionProgressDeadline DeploymentStatusDescription = "Failed due to progress deadline"

	// ApplyStatusDescriptionCanceled description when apply canceld
	ApplyStatusDescriptionCanceled DeploymentStatusDescription = "Deployment canceld"
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
