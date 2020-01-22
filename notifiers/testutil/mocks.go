package testutil

import (
	"errors"
	"io"
	"statusbay/notifiers"
	"statusbay/notifiers/common"
	watcherCommon "statusbay/watcher/kubernetes/common"
)

var (
	emptyNotifierMaker notifiers.NotifierMaker = func(defaultConfigReader io.Reader, config common.NotifierConfig, urlBase string) (notifier common.Notifier, err error) {
		return
	}
)

func GetNotifierMakerMock(makerType, errorMessage string) notifiers.NotifierMaker {
	switch makerType {
	case "error":
		return func(defaultConfigReader io.Reader, config common.NotifierConfig, urlBase string) (notifier common.Notifier, err error) {
			return nil, errors.New(errorMessage)
		}
	case "mock":
		return func(defaultConfigReader io.Reader, config common.NotifierConfig, urlBase string) (notifier common.Notifier, err error) {
			return &NotifierMock{}, nil
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

func (*NotifierMock) ReportStarted(message watcherCommon.DeploymentReporter) {
	panic("implement me")
}

func (*NotifierMock) ReportDeleted(message watcherCommon.DeploymentReporter) {
	panic("implement me")
}

func (*NotifierMock) ReportEnded(message watcherCommon.DeploymentReporter) {
	panic("implement me")
}
