package common

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

	// DeploymentStatusTimeout when statusbay stop watch
	DeploymentCanceled DeploymentStatus = "cancelled"
)

// DeploymentReporter defined deployment reporter message
type DeploymentReporter struct {
	// List of channels/username to send message to
	To []string

	// Deployment owner
	DeployBy string

	// Name of the deployment
	Name string

	// Status of the deployment
	Status DeploymentStatus

	//Deployment URI
	URI string
}
