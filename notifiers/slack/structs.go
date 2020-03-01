package slack

import (
	slackApi "github.com/nlopes/slack"
)

var defaultMessageConfig = map[ReportStage]*Message{
	started: {
		Pretext: "Kubernetes deployment started {deployed_by}",
		Text:    "Metrics, events, links and more are available through the <{link}|StatusBay report>",
	},
	ended: {
		Pretext: "Kubernetes deployment finished with status {status}",
		Text:    "<{link}|Click here> to view the StatusBay report",
	},
	deleted: {
		Pretext: "Deployment deleted {deployed_by}",
		Text:    "<{link}|Click here> to view the StatusBay report",
	},
}

type ReportStage string

const (
	started ReportStage = "beginning_message"
	ended   ReportStage = "end_message"
	deleted ReportStage = "deleted_message"
)

type MessageColor string

const (
	yellow MessageColor = "#ffeeba"
	red    MessageColor = "#c84034"
	blue   MessageColor = "#3aa3e3"
	green  MessageColor = "#25ba81"
)

type ApiClient interface {
	PostMessage(channelID string, options ...slackApi.MsgOption) (string, string, error)
	GetUsers() ([]slackApi.User, error)
}

type Message struct {
	Title   string `yaml:"title" mapstructure:"title"`
	Pretext string `yaml:"pretext" mapstructure:"pretext"`
	Text    string `yaml:"text" mapstructure:"text"`
}

type Config struct {
	Token            string                   `yaml:"token" mapstructure:"token"`
	DefaultChannels  []string                 `yaml:"default_channels" mapstructure:"default_channels"`
	MessageTemplates map[ReportStage]*Message `yaml:"message_templates" mapstructure:"message_templates"`
}

type Manager struct {
	client      ApiClient
	emailToUser map[string]string
	config      Config
	urlBase     string
}
