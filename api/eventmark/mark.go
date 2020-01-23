package eventmark

import (
	"statusbay/config"
	"strings"
)

// MarkEvent will mark message event if the string contain in the given mark list configuration
func MarkEvent(messageEvent string, marks []config.EventMarksConfig) []string {

	markDescriptions := []string{}
	for _, mark := range marks {
		if strings.Contains(strings.ToLower(messageEvent), strings.ToLower(mark.Pattern)) {
			markDescriptions = append(markDescriptions, mark.Descriptions...)
		}
	}
	return markDescriptions
}
