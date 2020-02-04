package load

import (
	"fmt"
	"os"
	"statusbay/notifiers"
	"statusbay/notifiers/common"
	"statusbay/notifiers/slack"
)

var (
	GetDefaultConfigReaderFunc = getDefaultConfigReader
)

// RegisterNotifiers registers existing notifier ctor to the ctor map we use to initiate all notifiers
func RegisterNotifiers() {
	notifiers.Register("slack", slack.NewSlack)
}

// getDefaultConfigReader returns the io.reader for the default config file for a specific notifier
func getDefaultConfigReader(basePath string, notifierName common.NotifierName) (*os.File, error) {
	return os.Open(fmt.Sprintf("%s/notifiers/%s/defaults.yaml", basePath, notifierName))
}

// Load returns a list of notifiers that were provided in the config and are implemented
func Load(rawNotifiersConfig common.ConfigByName, basePath, baseKubernetesUrl string) (notifierInstances []common.Notifier, err error) {
	var (
		notifierMaker         notifiers.NotifierMaker
		defaultNotifierConfig *os.File
		notifier              common.Notifier
	)

	for notifierName, notifierConfig := range rawNotifiersConfig {
		if notifierMaker, err = notifiers.GetNotifierMaker(notifierName); err != nil {
			return
		}

		if defaultNotifierConfig, err = GetDefaultConfigReaderFunc(basePath, notifierName); err != nil {
			return
		}

		if notifier, err = notifierMaker(defaultNotifierConfig, baseKubernetesUrl); err != nil {
			return
		}

		if err = notifier.LoadConfig(notifierConfig); err != nil {
			return
		}

		notifierInstances = append(notifierInstances, notifier)
	}
	return
}
