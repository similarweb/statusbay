package slack_test

import (
	"errors"
	"statusbay/notifiers/common"
	"statusbay/notifiers/slack"
	"statusbay/notifiers/testutil"
	"strings"
	"testing"
)

func TestSlackCtor(t *testing.T) {
	t.Run("Failed to read default config", func(t *testing.T) {
		errorMessage := "read failed"

		slackNotifier, err := slack.NewSlack(testutil.NewMockReader("", errors.New(errorMessage)), "")

		if err == nil {
			t.Error("Error expected")
		}

		if slackNotifier != nil {
			t.Error("the slack notifier should not have been initialized")
		}

	})

	t.Run("Failed to parse config yaml", func(t *testing.T) {
		expectedError := "line 1: cannot unmarshal"

		slackNotifier, err := slack.NewSlack(testutil.NewMockReader("123", nil), "")

		if err == nil {
			t.Error("Error expected")
		}

		if err != nil && !strings.Contains(err.Error(), expectedError) {
			t.Errorf("Unexpected error, expected error to contain %s", expectedError)
		}

		if slackNotifier != nil {
			t.Error("the slack notifier should not have been initialized")
		}

	})

	t.Run("returns a slack notifier instance", func(t *testing.T) {

		slackNotifier, err := slack.NewSlack(testutil.NewMockReader("", nil), "")

		if err != nil {
			t.Error("Unexpected error")
		}

		if slackNotifier == nil {
			t.Error("the slack notifier should have been initialized")
		}

	})
}

func TestSlackLoadConfig(t *testing.T) {
	slackNotifier, _ := slack.NewSlack(testutil.NewMockReader("", nil), "url-base.com")

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
