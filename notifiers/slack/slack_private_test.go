package slack

import (
	"statusbay/notifiers/common"
	"testing"
)

func TestSlackLoadConfig(t *testing.T) {
	slackManager := Manager{
		urlBase: "",
		messageConfig: map[ReportStage]*Message{

		}
	}

	t.Run("failed to decode common.NotifierConfig to `Config` struct", func(t *testing.T) {
		if err := slackNotifier.LoadConfig(common.NotifierConfig{"default_channels": "fail"}); err == nil {
			t.Error("Error expected")
		}
	})

	t.Run("failed to decode common.NotifierConfig to `Config` struct", func(t *testing.T) {
		expectedError := "slack token is required"
		if err := slackNotifier.LoadConfig(nil); err == nil {
			t.Error("Error expected")
		} else if err.Error() != expectedError {
			t.Errorf("Expected error to be %s, instead got `%s`", expectedError, err.Error())
		}
	})
}
