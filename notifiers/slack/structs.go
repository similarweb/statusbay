package slack

import (
	slackApi "github.com/nlopes/slack"
)

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

type Message struct {
	Title   string `yaml:"title" mapstructure:"title"`
	Pretext string `yaml:"pretext" mapstructure:"pretext"`
	Text    string `yaml:"text" mapstructure:"text"`
}

type Config struct {
	Token            string   `yaml:"token" mapstructure:"token"`
	DefaultChannels  []string `yaml:"default_channels" mapstructure:"default_channels"`
	BeginningMessage *Message `yaml:"beginning_message" mapstructure:"beginning_message"`
	EndMessage       *Message `yaml:"end_message" mapstructure:"end_message"`
	DeletedMessage   *Message `yaml:"deleted_message" mapstructure:"deleted_message"`
}

type Manager struct {
	client        *slackApi.Client
	emailToUser   map[string]string
	messageConfig map[ReportStage]*Message
	config        Config
	urlBase       string
}
