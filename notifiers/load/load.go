package load

import (
	"statusbay/notifiers"
	"statusbay/notifiers/common"
	"statusbay/notifiers/slack"
)

// RegisterNotifiers registers existing notifier ctor to the ctor map we use to initiate all notifiers
func RegisterNotifiers() {
	notifiers.Register("slack", slack.NewSlack)
}

// Load returns a list of notifiers that were provided in the config and are implemented
func Load(rawNotifiersConfig common.ConfigByName, baseKubernetesUrl string) (notifierInstances []common.Notifier, err error) {
	var (
		notifierMaker notifiers.NotifierMaker
		notifier      common.Notifier
	)

	for notifierName, notifierConfig := range rawNotifiersConfig {
		if notifierMaker, err = notifiers.GetNotifierMaker(notifierName); err != nil {
			return
		}

		notifier = notifierMaker(baseKubernetesUrl)
		if err = notifier.LoadConfig(notifierConfig); err != nil {
			return
		}

		notifierInstances = append(notifierInstances, notifier)
	}
	return
}
