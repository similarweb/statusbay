package kuberneteswatcher

import (
	"context"
	notifierCommon "statusbay/notifiers/common"
	"statusbay/watcher/kubernetes/common"
	"sync"

	log "github.com/sirupsen/logrus"
)

// ReporterManager defined reporter metadata
type ReporterManager struct {
	// Received channel when deployment started
	DeploymentStarted chan common.DeploymentReport

	// Received channel when deployment deleted
	DeploymentDeleted chan common.DeploymentReport

	// Received channel when deployment finish
	DeploymentFinished chan common.DeploymentReport

	// available ways to notify about changes in the deployment stages
	availableNotifiers []notifierCommon.Notifier
}

// NewReporter creates new reporter
func NewReporter(availableNotifiers []notifierCommon.Notifier) *ReporterManager {
	return &ReporterManager{
		availableNotifiers: availableNotifiers,

		DeploymentStarted:  make(chan common.DeploymentReport),
		DeploymentDeleted:  make(chan common.DeploymentReport),
		DeploymentFinished: make(chan common.DeploymentReport),
	}
}

// Serve will start to reporter listener
func (re *ReporterManager) Serve(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)

	for _, notifier := range re.availableNotifiers {
		notifier.Serve(ctx, wg)
	}

	go func() {
		for {
			select {
			case request := <-re.DeploymentStarted:
				re.deploymentStarted(request)
			case request := <-re.DeploymentDeleted:
				re.deploymentDeleted(request)
			case request := <-re.DeploymentFinished:
				re.deploymentFinish(request)
			case <-ctx.Done():
				log.Warn("reporter has been shut down")
				wg.Done()
				return
			}
		}
	}()
	log.Info("reporter started")

}

// deploymentStarted will send slack message to channel/user when the deployment is started
func (re *ReporterManager) deploymentStarted(message common.DeploymentReport) {
	for _, notifier := range re.availableNotifiers {
		notifier.ReportStarted(message)
	}
}

// deploymentDeleted will send slack message to channel/user when the deployment is deleted
func (re *ReporterManager) deploymentDeleted(message common.DeploymentReport) {
	for _, notifier := range re.availableNotifiers {
		notifier.ReportDeleted(message)
	}
}

// deploymentFinish will send slack message to channel/user when the deployment is finished
func (re *ReporterManager) deploymentFinish(message common.DeploymentReport) {
	for _, notifier := range re.availableNotifiers {
		notifier.ReportEnded(message)
	}
}
