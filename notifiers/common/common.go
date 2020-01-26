package common

import (
	"statusbay/serverutil"
	"statusbay/watcher/kubernetes/common"
)

const (
	StatusPlaceholder     = "{status}"
	LinkPlaceholder       = "{link}"
	DeployedByPlaceholder = "{deployed_by}"
)

type NotifierName string
type NotifierConfig map[string]interface{}
type ConfigByName map[NotifierName]NotifierConfig

type Notifier interface {
	LoadConfig(notifierConfig NotifierConfig) (err error)
	ReportStarted(message common.DeploymentReporter)
	ReportDeleted(message common.DeploymentReporter)
	ReportEnded(message common.DeploymentReporter)
	Serve() serverutil.StopFunc
}
