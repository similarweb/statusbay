package testutil

import (
	"errors"
	"statusbay/notifiers"
	"statusbay/notifiers/common"
	"statusbay/serverutil"
	watcherCommon "statusbay/watcher/kubernetes/common"
)

var (
	emptyNotifierMaker notifiers.NotifierMaker = func(urlBase string) common.Notifier {
		return nil
	}
)

func GetNotifierMakerMock(makerType, errorMessage string) notifiers.NotifierMaker {
	switch makerType {
	case "mock":
		if errorMessage == "" {
			return func(urlBase string) common.Notifier {
				return &NotifierMock{}
			}
		} else {
			return func(urlBase string) common.Notifier {
				return &NotifierMock{err: errors.New(errorMessage)}
			}
		}
	default:
		return emptyNotifierMaker
	}
}

type NotifierMock struct {
	err error
}

func (n *NotifierMock) LoadConfig(common.NotifierConfig) (err error) {
	if n.err != nil {
		err = n.err
	}
	return
}

func (*NotifierMock) ReportStarted(watcherCommon.DeploymentReport) {
	panic("implement me")
}

func (*NotifierMock) ReportDeleted(watcherCommon.DeploymentReport) {
	panic("implement me")
}

func (*NotifierMock) ReportEnded(watcherCommon.DeploymentReport) {
	panic("implement me")
}

func (*NotifierMock) Serve() (sf serverutil.StopFunc) {
	return
}
