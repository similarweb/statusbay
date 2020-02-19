package slack_test

import (
	"statusbay/notifiers/common"
	"statusbay/notifiers/slack"
	"testing"
)

func TestSlackCtor(t *testing.T) {
	t.Run("returns a slack notifier instance", func(t *testing.T) {

		slackNotifier := slack.NewSlack("")

		if slackNotifier == nil {
			t.Error("the slack notifier should have been initialized")
		}

	})
}

func TestSlackLoadConfig(t *testing.T) {
	slackNotifier := slack.NewSlack("url-base.com")

	t.Run("failed to decode common.NotifierConfig to `Config` struct", func(t *testing.T) {
		if err := slackNotifier.LoadConfig(common.NotifierConfig{"default_channels": "fail"}); err == nil {
			t.Error("error expected")
		}
	})

	t.Run("failed to decode common.NotifierConfig to `Config` struct", func(t *testing.T) {
		if err := slackNotifier.LoadConfig(nil); err == nil {
			t.Error("error expected")
		} else if err.Error() != slack.NoTokenErr.Error() {
			t.Errorf("expected error to be %s, instead got `%s`", slack.NoTokenErr.Error(), err.Error())
		}
	})
}
