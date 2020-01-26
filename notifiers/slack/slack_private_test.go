package slack

import (
	"fmt"
	"statusbay/notifiers/common"
	"testing"
)

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
