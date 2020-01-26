package testutil

import (
	"errors"
	"io"
	"statusbay/notifiers"
	"statusbay/notifiers/common"
	watcherCommon "statusbay/watcher/kubernetes/common"
	"strings"
)

var (
	emptyNotifierMaker notifiers.NotifierMaker = func(defaultConfigReader io.Reader, urlBase string) (notifier common.Notifier, err error) {
		return
	}
)

func GetNotifierMakerMock(makerType, errorMessage string) notifiers.NotifierMaker {
	switch makerType {
	case "error":
		return func(defaultConfigReader io.Reader, urlBase string) (notifier common.Notifier, err error) {
			return nil, errors.New(errorMessage)
		}
	case "mock":
		if errorMessage == "" {
			return func(defaultConfigReader io.Reader, urlBase string) (notifier common.Notifier, err error) {
				return &NotifierMock{}, nil
			}
		} else {
			return func(defaultConfigReader io.Reader, urlBase string) (notifier common.Notifier, err error) {
				return &NotifierMock{err: errors.New(errorMessage)}, nil
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

func (*NotifierMock) ReportStarted(message watcherCommon.DeploymentReporter) {
	panic("implement me")
}

func (*NotifierMock) ReportDeleted(message watcherCommon.DeploymentReporter) {
	panic("implement me")
}

func (*NotifierMock) ReportEnded(message watcherCommon.DeploymentReporter) {
	panic("implement me")
}

func NewMockReader(source string, err error) io.Reader {
	if err != nil {
		return &ReaderMock{err: err}
	} else {
		return &ReaderMock{source: strings.NewReader(source)}
	}
}

type ReaderMock struct {
	source io.Reader
	err    error
}

func (r *ReaderMock) Read(p []byte) (n int, err error) {
	if r.err != nil {
		err = r.err
		return
	}
	return r.source.Read(p)
}
