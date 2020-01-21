package load

import (
	"fmt"
	"os"
	"statusbay/notifiers"
	"statusbay/notifiers/common"
	"statusbay/notifiers/slack"
)

func RegisterNotifiers() {
	notifiers.Register("slack", slack.NewSlack)
}

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

		if defaultNotifierConfig, err = os.Open(fmt.Sprintf("%s/notifiers/%s/defaults.yaml", basePath, notifierName)); err != nil {
			return
		}

		if notifier, err = notifierMaker(defaultNotifierConfig, notifierConfig, baseKubernetesUrl); err != nil {
			return
		}

		notifierInstances = append(notifierInstances, notifier)
	}
	return
}
