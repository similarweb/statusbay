package slack

import (
	"errors"
	"fmt"
	"github.com/nlopes/slack"
	"statusbay/notifiers/common"
	watcherCommon "statusbay/watcher/kubernetes/common"

	"testing"
)

type SentMessage struct {
	channelId string
}

type MockApiClient struct {
	sentMessages []SentMessage
	users        []slack.User
	err          error
	idx          int
}

func (m *MockApiClient) PostMessage(channelID string, _ ...slack.MsgOption) (string, string, error) {
	if m.err != nil {
		return "", "", m.err
	}

	m.sentMessages = append(m.sentMessages, SentMessage{
		channelId: channelID,
	})
	return "", "", nil
}

func (m *MockApiClient) GetUsers() ([]slack.User, error) {
	return m.users, m.err
}

func TestSlackLoadConfig(t *testing.T) {
	defaultConfig := map[ReportStage]*Message{
		started: {
			Title:   "initial_title",
			Pretext: "initial_pretext",
			Text:    "initial_text",
		},
		ended: {
			Title:   "initial_title",
			Pretext: "initial_pretext",
			Text:    "initial_text",
		},
		deleted: {
			Title:   "initial_title",
			Pretext: "initial_pretext",
			Text:    "initial_text",
		},
	}

	initialConfig := map[ReportStage]*Message{}
	for stage, message := range defaultConfig {
		initialConfig[stage] = &Message{
			Title:   message.Title,
			Pretext: message.Pretext,
			Text:    message.Text,
		}
	}

	t.Run("no custom settings, using defaults", func(t *testing.T) {
		slackManager := Manager{
			urlBase: "",
			config:  Config{MessageTemplates: initialConfig},
		}

		if err := slackManager.LoadConfig(common.NotifierConfig{"token": "test_token"}); err != nil {
			t.Error("Unexpected error")
		}

		for stage, message := range slackManager.config.MessageTemplates {
			if defaultConfig[stage].Title != message.Title {
				t.Errorf("Expected to still have %s (default) for Title in %s. instead got `%s`", defaultConfig[stage].Title, stage, message.Title)
			}

			if defaultConfig[stage].Pretext != message.Pretext {
				t.Errorf("Expected to still have %s (default) for Pretext in %s. instead got `%s`", defaultConfig[stage].Pretext, stage, message.Pretext)
			}

			if defaultConfig[stage].Text != message.Text {
				t.Errorf("Expected to still have %s (default) for Text in %s. instead got `%s`", defaultConfig[stage].Text, stage, message.Text)
			}
		}
	})

	t.Run("checking if config overrides defaults", func(t *testing.T) {
		for updatedStage := range defaultConfig {

			// new config to override the default one
			newConfig := common.NotifierConfig{
				"token": "test_token",
				string(updatedStage): map[string]string{
					"title":   "modified_title",
					"pretext": "modified_pretext",
					"text":    "modified_text",
				},
			}

			slackManager := Manager{
				urlBase: "",
				config:  Config{MessageTemplates: initialConfig},
			}

			t.Run(fmt.Sprintf("provided new value for %s only", updatedStage), func(t *testing.T) {
				if err := slackManager.LoadConfig(newConfig); err != nil {
					t.Error("Unexpected error")
				}

				for stage, message := range slackManager.config.MessageTemplates {

					if stage == updatedStage {
						expectedStageConfig := newConfig[string(stage)].(map[string]string)
						if expectedStageConfig["title"] != message.Title {
							t.Errorf("Expected to still have %s (updated) for Title in %s. instead got `%s`", expectedStageConfig["title"], stage, message.Title)
						}

						if expectedStageConfig["pretext"] != message.Pretext {
							t.Errorf("Expected to still have %s (updated) for Pretext in %s. instead got `%s`", expectedStageConfig["pretext"], stage, message.Pretext)
						}

						if expectedStageConfig["text"] != message.Text {
							t.Errorf("Expected to still have %s (updated) for Text in %s. instead got `%s`", expectedStageConfig["text"], stage, message.Text)
						}
					} else {
						if defaultConfig[stage].Title != message.Title {
							t.Errorf("Expected to still have %s (default) for Title in %s. instead got `%s`", defaultConfig[stage].Title, stage, message.Title)
						}

						if defaultConfig[stage].Pretext != message.Pretext {
							t.Errorf("Expected to still have %s (default) for Pretext in %s. instead got `%s`", defaultConfig[stage].Pretext, stage, message.Pretext)
						}

						if defaultConfig[stage].Text != message.Text {
							t.Errorf("Expected to still have %s (default) for Text in %s. instead got `%s`", defaultConfig[stage].Text, stage, message.Text)
						}
					}
				}
			})
		}
	})
}

func TestReplacePlaceholders(t *testing.T) {
	statusValue := "status"
	linkValue := "link"
	deployedByValue := "deployed_by"

	t.Run("nothing to replace, value should stay the same", func(t *testing.T) {
		input := "nothing to replace here, {test}"

		if result := replacePlaceholders(input, statusValue, linkValue, deployedByValue); result != input {
			t.Errorf("Expected result to be `%s` but got `%s`", input, result)
		}
	})

	t.Run("expect all placeholders to be replaced", func(t *testing.T) {

		input := fmt.Sprintf("status: %s, link: %s, deployed_by: %s", common.StatusPlaceholder, common.LinkPlaceholder, common.DeployedByPlaceholder)
		expected := fmt.Sprintf("status: %s, link: %s, deployed_by: %s", statusValue, linkValue, deployedByValue)

		result := replacePlaceholders(input, "status", "link", "deployed_by")
		if result == input {
			t.Errorf("Expected result to be `%s` but got `%s`", expected, result)
		}

		if result != expected {
			t.Errorf("Expected result to be `%s` but got `%s`", expected, result)
		}

	})
}

func TestDistinct(t *testing.T) {
	t.Run("de-duplicates a slice", func(t *testing.T) {
		input := []string{"a", "b", "b", "c", "d", "d"}
		expected := []string{"a", "b", "c", "d"}

		result := distinct(input)

		equal := true

		for i := range result {
			if result[i] != expected[i] {
				equal = false
				break
			}
		}

		if !equal {
			t.Errorf("expected result `%v` to match `%v`", result, expected)
		}
	})

	t.Run("slice remains unchanged as there are no duplicates in it", func(t *testing.T) {
		input := []string{"a", "b22", "11c", "d"}
		expected := []string{"a", "b22", "11c", "d"}

		result := distinct(input)

		equal := true

		for i := range result {
			if result[i] != expected[i] {
				equal = false
				break
			}
		}

		if !equal {
			t.Errorf("expected result `%v` to match `%v`", result, expected)
		}
	})
}

func TestUpdateUsers(t *testing.T) {
	t.Run("unable to get users from slack api, existing list remains the same", func(t *testing.T) {
		mockClient := &MockApiClient{
			err: errors.New(""),
		}

		initialValues := []string{"does", "not", "change"}
		mockEmailToUser := map[string]string{}

		for _, val := range initialValues {
			mockEmailToUser[val] = ""
		}

		slackManager := Manager{
			client:      mockClient,
			emailToUser: mockEmailToUser,
		}

		slackManager.updateUsers()

		if len(slackManager.emailToUser) != 3 {
			t.Errorf("expected emailToUsers contain exactly 3 emails instead has %d", len(slackManager.emailToUser))
		}

		for _, val := range initialValues {
			if _, exists := slackManager.emailToUser[val]; !exists {
				t.Errorf("expected %s to remain in the emailToUser map", val)
			}
		}

	})
}

func TestGetUserIdByEmail(t *testing.T) {
	availableEmails := map[string]string{
		"email1": "id1",
		"email3": "id3",
		"email5": "id5",
	}
	unavailableEmails := []string{"email2", "email4"}

	mockEmailToUser := map[string]string{}
	for email, id := range availableEmails {
		mockEmailToUser[email] = id
	}

	slackManager := Manager{
		emailToUser: mockEmailToUser,
	}

	t.Run("check for available emails", func(t *testing.T) {
		for email, id := range availableEmails {
			resultId, err := slackManager.getUserIdByEmail(email)
			if err != nil {
				t.Errorf("unexpected error %s", err.Error())
			} else if resultId != id {
				t.Errorf("resultId (`%s`) does not match expected id (`%s`)", resultId, id)
			}

		}
	})

	t.Run("check for unavailable emails", func(t *testing.T) {
		for _, email := range unavailableEmails {
			_, err := slackManager.getUserIdByEmail(email)
			if err == nil {
				t.Errorf("expected error")
			}

		}
	})
}

func TestGetChannelId(t *testing.T) {
	availableEmails := map[string]string{
		"email1": "id1",
		"email2": "id2",
	}
	inputToExpected := map[string]string{
		"email1": "id1",
		"email2": "id2",
		"#chan":  "#chan",
	}

	mockEmailToUser := map[string]string{}
	for email, id := range availableEmails {
		mockEmailToUser[email] = id
	}

	slackManager := Manager{
		emailToUser: mockEmailToUser,
	}

	t.Run("check for different inputs", func(t *testing.T) {
		for input, expected := range inputToExpected {
			resultId, err := slackManager.GetChannelId(input)
			if err != nil {
				t.Errorf("unexpected error %s", err.Error())
			} else if resultId != expected {
				t.Errorf("resultId (`%s`) does not match expected id (`%s`)", resultId, expected)
			}

		}
	})

	t.Run("check for unavailable emails", func(t *testing.T) {
		if _, err := slackManager.GetChannelId("email3"); err == nil {
			t.Errorf("expected error")
		}

		if _, err := slackManager.GetChannelId("email6i5"); err == nil {
			t.Errorf("expected error")
		}
	})
}

func TestSend(t *testing.T) {
	messagesToSend := []SentMessage{
		{
			channelId: "#chan1",
		},
		{
			channelId: "email1",
		},
	}

	t.Run("fails to send messages", func(t *testing.T) {
		mockClient := &MockApiClient{err: errors.New("")}
		slackManager := Manager{client: mockClient}

		for _, message := range messagesToSend {
			slackManager.send(message.channelId, slack.Attachment{})
		}

		if len(mockClient.sentMessages) != 0 {
			t.Errorf("expected to successfully send 0 messages, sent %d", len(mockClient.sentMessages))
		}
	})

	t.Run("sent all messages we expected to send", func(t *testing.T) {
		mockClient := &MockApiClient{}
		slackManager := Manager{client: mockClient}

		for _, message := range messagesToSend {
			slackManager.send(message.channelId, slack.Attachment{})
		}

		if len(messagesToSend) != len(mockClient.sentMessages) {
			t.Errorf("expected the number of sent messages `%d` to match the number of messages we tried to send `%d`", len(mockClient.sentMessages), len(messagesToSend))
		}

		for idx := range messagesToSend {
			if messagesToSend[idx].channelId != mockClient.sentMessages[idx].channelId {
				t.Errorf("channelId mismatch between between the message we wanted to sent and the one we sent %s!=%s", messagesToSend[idx].channelId, mockClient.sentMessages[idx].channelId)
			}
		}
	})

}

func TestReportFuncs(t *testing.T) {
	t.Run("sending success message from an unknown user", func(t *testing.T) {
		mockClient := &MockApiClient{}
		slackManager := Manager{
			client: mockClient,
			emailToUser: map[string]string{
				"test1": "test1",
			},
			config: Config{
				DefaultChannels: []string{"#default_test"},
				MessageTemplates: map[ReportStage]*Message{
					started: {},
					ended:   {},
					deleted: {},
				},
			},
		}

		expectedMessageTargets := map[string]*struct{}{
			"test1":         nil,
			"#default_test": nil,
		}

		slackManager.ReportStarted(watcherCommon.DeploymentReporter{
			To:       []string{"test1", "test1", "test2", ""},
			DeployBy: "unknown",
		})

		if mockClient.sentMessages == nil {
			t.Error("expected to send messages, sent none")
		} else {

			if len(mockClient.sentMessages) != 2 {
				t.Errorf("expected to send 2 messages, sent %d", len(mockClient.sentMessages))
			}

			for _, target := range mockClient.sentMessages {
				if _, exists := expectedMessageTargets[target.channelId]; !exists {
					t.Errorf("expected %s to be a target of a sent message", target.channelId)
				}
			}
		}

	})

	t.Run("sending end message from a known user", func(t *testing.T) {
		mockClient := &MockApiClient{}
		slackManager := Manager{
			client:      mockClient,
			emailToUser: map[string]string{"email1": "id2"},
			config: Config{
				DefaultChannels: []string{"#default_test"},
				MessageTemplates: map[ReportStage]*Message{
					started: {},
					ended:   {},
					deleted: {},
				},
			},
		}

		expectedMessageTargets := map[string]*struct{}{
			"test1":         nil,
			"#default_test": nil,
		}

		slackManager.ReportEnded(watcherCommon.DeploymentReporter{
			To:       []string{"test1", "test1", ""},
			DeployBy: "email1",
		})

		if mockClient.sentMessages == nil {
			t.Error("expected to send messages, sent none")
		} else {

			if len(mockClient.sentMessages) != 1 {
				t.Errorf("expected to send 1 message, sent %d", len(mockClient.sentMessages))
			}

			for _, target := range mockClient.sentMessages {
				if _, exists := expectedMessageTargets[target.channelId]; !exists {
					t.Errorf("expected %s to be a target of a sent message", target.channelId)
				}
			}
		}

	})

	t.Run("sending delete message from a known user", func(t *testing.T) {
		mockClient := &MockApiClient{}
		slackManager := Manager{
			client:      mockClient,
			emailToUser: map[string]string{"email1": "id2"},
			config: Config{
				DefaultChannels: []string{"#default_test"},
				MessageTemplates: map[ReportStage]*Message{
					started: {},
					ended:   {},
					deleted: {},
				},
			},
		}

		expectedMessageTargets := map[string]*struct{}{
			"test1":         nil,
			"#default_test": nil,
		}

		slackManager.ReportDeleted(watcherCommon.DeploymentReporter{
			To:       []string{"test1", "test1", ""},
			DeployBy: "email1",
		})

		if mockClient.sentMessages == nil {
			t.Error("expected to send messages, sent none")
		} else {

			if len(mockClient.sentMessages) != 1 {
				t.Errorf("expected to send 1 message, sent %d", len(mockClient.sentMessages))
			}

			for _, target := range mockClient.sentMessages {
				if _, exists := expectedMessageTargets[target.channelId]; !exists {
					t.Errorf("expected %s to be a target of a sent message", target.channelId)
				}
			}
		}

	})

}
