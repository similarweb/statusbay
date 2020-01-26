package kuberneteswatcher

import (
	"context"
	log "github.com/sirupsen/logrus"
	notifierCommon "statusbay/notifiers/common"
	"statusbay/serverutil"
	"statusbay/watcher/kubernetes/common"
)

// ReporterManager defined reporter metadata
type ReporterManager struct {
	// Received channel when deployment started
	DeploymentStarted chan common.DeploymentReporter

	// Received channel when deployment deleted
	DeploymentDeleted chan common.DeploymentReporter

	// Received channel when deployment finish
	DeploymentFinished chan common.DeploymentReporter

	// available ways to notify about changes in the deployment stages
	availableNotifiers []notifierCommon.Notifier
}

// NewReporter creates new reporter
func NewReporter(availableNotifiers []notifierCommon.Notifier) *ReporterManager {
	return &ReporterManager{
		availableNotifiers: availableNotifiers,

		DeploymentStarted:  make(chan common.DeploymentReporter),
		DeploymentDeleted:  make(chan common.DeploymentReporter),
		DeploymentFinished: make(chan common.DeploymentReporter),
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

	return func() {
		cancelFn()
		<-stopped
	}
}

// deploymentStarted will send slack message to channel/user when the deployment is started
func (re *ReporterManager) deploymentStarted(message common.DeploymentReporter) {
	for _, notifier := range re.availableNotifiers {
		notifier.ReportStarted(message)
	}
}

// deploymentDeleted will send slack message to channel/user when the deployment is deleted
func (re *ReporterManager) deploymentDeleted(message common.DeploymentReporter) {
	for _, notifier := range re.availableNotifiers {
		notifier.ReportDeleted(message)
	}
}

// deploymentFinish will send slack message to channel/user when the deployment is finished
func (re *ReporterManager) deploymentFinish(message common.DeploymentReporter) {
	for _, notifier := range re.availableNotifiers {
		notifier.ReportDeleted(message)
	}
}
