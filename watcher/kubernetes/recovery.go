package kuberneteswatcher

import (
	"fmt"

	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
)

// RecoveryManager to recover restarts of the watcher
type RecoveryManager struct {

	// Registry manager will be owner to manage the running / new deployment
	registryManager *RegistryManager
	// Kubernetes client
	client kubernetes.Interface
}

func NewRecoveryManager(kubernetesClientset kubernetes.Interface, registryManager *RegistryManager) *RecoveryManager {
	return &RecoveryManager{
		registryManager: registryManager,
		client:          kubernetesClientset,
	}
}

// RecoverState updates the state before of the watcher starts
// 1. get all runnnig applies from the db
// 2. check each resource against remote k8s storage
// 3. update state on registryRow for each resource
// 4. END||run isFinished
func (rm *RecoveryManager) RecoverState() error {
	runningApplies := rm.registryManager.LoadRunningApplies()
	for _, app := range runningApplies {
		deployments := app.DBSchema.Resources.Deployments
		err := rm.recoverDeployments(app, deployments)
		if err != nil {
			fmt.Println("ERROR ISAN: error recovering ", err)
			return err
		}
		_ = app.DBSchema.Resources.Daemonsets
		_ = app.DBSchema.Resources.Statefulsets
	}
	return nil
}

func (rm *RecoveryManager) recoverDeployments(app *RegistryRow, runningDeployments map[string]*DeploymentData) error {
	for _, deploymentData := range runningDeployments {
		err := rm.recoverDeployment(app, deploymentData)
		if err != nil {
			return err
		}
	}
	return nil
}
func (rm *RecoveryManager) recoverDeployment(app *RegistryRow, runningDeployment *DeploymentData) error {
	ns := runningDeployment.Deployment.Namespace
	listOptions := metaV1.ListOptions{LabelSelector: labels.SelectorFromSet(runningDeployment.Deployment.Labels).String()}
	remoteDeployments, err := rm.client.AppsV1().Deployments(ns).List(listOptions)
	if err != nil {
		fmt.Println("ISAN: err ", err)
		return err
	}
	for _, remoteDep := range remoteDeployments.Items {
		fmt.Println("Recover: @@@@@@@@@@@@@@@@@@@@@@@a@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
		fmt.Println("Replicas = ", remoteDep.Status.Replicas)
		fmt.Println("UpdatedReplicas = ", remoteDep.Status.UpdatedReplicas)
		fmt.Println("ReadyReplicas = ", remoteDep.Status.ReadyReplicas)
		fmt.Println("AvailableReplicas = ", remoteDep.Status.AvailableReplicas)
		fmt.Println("UnavailableReplicas = ", remoteDep.Status.UnavailableReplicas)
		matchLabels := remoteDep.Spec.Selector.MatchLabels
		runningDeployment.UpdateDeploymentStatus(remoteDep.Status)
		err = rm.getReplicaSets(app, runningDeployment, matchLabels)
		fmt.Println("Recover: @@@@@@@@@@@@@@@@@@@@@@@a@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
	}
	return nil
}
func (rm *RecoveryManager) getReplicaSets(app *RegistryRow, deployment *DeploymentData, matchLabels map[string]string) error {
	listOptions := metaV1.ListOptions{LabelSelector: labels.SelectorFromSet(matchLabels).String()}
	replicas, _ := rm.client.AppsV1().ReplicaSets(deployment.Deployment.Namespace).List(listOptions)
	for _, rep := range replicas.Items {
		deployment.InitReplicaset(rep.GetName())
		err := deployment.UpdateReplicasetStatus(rep.GetName(), rep.Status)
		if err != nil {
			fmt.Println("isan :: err replicaset ", err)
			return err
		}
		fmt.Println("isan: maybe updated replicaset status")
	}
	return nil
}
func (rm *RecoveryManager) recoverDaemonsets() error {
	return nil
}
func (rm *RecoveryManager) recoverStatefulsets() error {
	return nil
}
