package kuberneteswatcher

import (
	"context"
	notifierCommon "statusbay/notifiers/common"
	"statusbay/serverutil"
	"statusbay/watcher/kubernetes/common"

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
func (re *ReporterManager) Serve() serverutil.StopFunc {

	ctx, ctxCancel := context.WithCancel(context.Background())
	var notifierStoppers []serverutil.StopFunc
	for _, notifier := range re.availableNotifiers {
		notifierStoppers = append(notifierStoppers, notifier.Serve())
	}
	cancelFn := func() {
		ctxCancel()
		for _, notifierCancelFunc := range notifierStoppers {
			notifierCancelFunc()
		}
	}

	stopped := make(chan bool)
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
				log.Warn("Reporter has been shut down")
				stopped <- true
				return
			}
		}
	}()
	log.Info("Reporter started")

	return func() {
		cancelFn()
		<-stopped
	}
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
