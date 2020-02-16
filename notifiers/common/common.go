package common

import (
	"context"
	"statusbay/watcher/kubernetes/common"
	"sync"
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
	ReportStarted(message common.DeploymentReport)
	ReportDeleted(message common.DeploymentReport)
	ReportEnded(message common.DeploymentReport)
	Serve(ctx context.Context, wg *sync.WaitGroup)
}
