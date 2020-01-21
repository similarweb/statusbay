package slack

import (
	"context"
	"fmt"
	"github.com/apex/log"
	"github.com/mitchellh/mapstructure"
	slackApi "github.com/nlopes/slack"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"statusbay/notifiers/common"
	"statusbay/serverutil"
	watcherCommon "statusbay/watcher/kubernetes/common"
	"strings"
	"time"
)

var (
	noTokenErr = errors.New("slack token is required")
)

const (
	//UpdateSlackUserInterval interval for update slack user list
	UpdateSlackUserInterval = time.Hour
)

func (sl *Manager) LoadConfig(notifierConfig common.NotifierConfig) (err error) {
	sl.config = Config{}
	if err = mapstructure.Decode(notifierConfig, &sl.config); err != nil {
		return
	}

	// validate config
	if sl.config.Token == "" {
		return noTokenErr
	}

	// overwrite defaults
	if sl.config.BeginningMessage != nil {
		sl.messageConfig[started] = sl.config.BeginningMessage
	}

	if sl.config.EndMessage != nil {
		sl.messageConfig[ended] = sl.config.EndMessage
	}

	if sl.config.DeletedMessage != nil {
		sl.messageConfig[deleted] = sl.config.DeletedMessage
	}

	return
}

func (sl *Manager) sendToAll(stage ReportStage, message watcherCommon.DeploymentReporter, color MessageColor) {
	var (
		deployBy string
		err      error
	)

	if deployBy, err = sl.getUserIDByEmail(message.DeployBy); err != nil {
		deployBy = message.DeployBy
	} else {
		deployBy = fmt.Sprintf("by <@%s>", deployBy)
	}

	status := strings.ToUpper(string(message.Status))
	link := fmt.Sprintf("%s/%s", sl.urlBase, message.URI)

	for _, to := range distinct(append(message.To, sl.config.DefaultChannels...)) {
		if to == "" {
			continue
		}
		toChannel, err := sl.GetChannelID(to)
		if err == nil {
			attachment := slackApi.Attachment{
				Title:   replacePlaceholders(sl.messageConfig[stage].Title, status, link, deployBy),
				Pretext: replacePlaceholders(sl.messageConfig[stage].Pretext, status, link, deployBy),
				Text:    replacePlaceholders(sl.messageConfig[stage].Text, status, link, deployBy),
				Color:   string(color),
				// TODO:: add cluster + namespace name
				Fields: []slackApi.AttachmentField{
					{
						Title: "Application Name:",
						Value: message.Name,
						Short: false,
					},
				},
			}
			sl.send(toChannel, attachment)

		} else {
			log.WithFields(log.Fields{
				"to": to,
			}).Debug("Slack id not found")
		}

	}
}

func (sl *Manager) ReportStarted(message watcherCommon.DeploymentReporter) {
	sl.sendToAll(started, message, blue)
}

func (sl *Manager) ReportDeleted(message watcherCommon.DeploymentReporter) {
	sl.sendToAll(deleted, message, red)
}

func (sl *Manager) ReportEnded(message watcherCommon.DeploymentReporter) {
	color := green

	switch message.Status {
	case watcherCommon.DeploymentSuccessful:
		color = green
	case watcherCommon.DeploymentCanceled:
		color = yellow
	case watcherCommon.DeploymentStatusFailed:
		color = red
	}

	sl.sendToAll(ended, message, color)
}

func NewSlack(defaultConfigReader io.Reader, config common.NotifierConfig, urlBase string) (notifier common.Notifier, err error) {
	var data []byte
	defaultMessageConfig := map[ReportStage]*Message{}

	if data, err = ioutil.ReadAll(defaultConfigReader); err != nil {
		return
	} else if err = yaml.Unmarshal(data, &defaultMessageConfig); err != nil {
		return
	}

	slackManager := &Manager{
		messageConfig: defaultMessageConfig,
		urlBase:       urlBase,
	}

	if err = slackManager.LoadConfig(config); err != nil {
		return
	} else {
		slackManager.updateUsers()
		notifier = slackManager
		return
	}
}

// Serve will loop of slack users
func (sl *Manager) Serve() serverutil.StopFunc {
	ctx, cancelFn := context.WithCancel(context.Background())
	stopped := make(chan bool)
	go func() {
		for {
			select {
			case <-time.After(UpdateSlackUserInterval):
				sl.updateUsers()
			case <-ctx.Done():
				log.Warn("Slack Loop has been shut down")
				stopped <- true
				return
			}
		}
	}()
	return func() {
		cancelFn()
		<-stopped
	}
}

// UpdateUsers all users from slack
func (sl *Manager) updateUsers() bool {
	currentUsers := map[string]string{}

	users, err := sl.client.GetUsers()
	if err != nil {
		return false
	}

	for _, user := range users {
		if !user.Deleted && user.Profile.Email != "" {
			currentUsers[user.Profile.Email] = user.ID
		}
	}
	if len(currentUsers) != len(sl.emailToUser) {
		sl.emailToUser = currentUsers
		log.Info(fmt.Sprintf("Found %d slack users", len(sl.emailToUser)))
	}

	return true
}

// getUserIDByEmail return slack user by email
func (sl *Manager) getUserIDByEmail(email string) (string, error) {
	if userId, ok := sl.emailToUser[email]; !ok {
		//log.WithField("email", email).Warn("Slack user by email not found")
		return "", errors.New("slack user by email not found")
	} else {
		return userId, nil
	}
}

// Send will send slack notification to user
func (sl *Manager) send(channelID string, attachment slackApi.Attachment) bool {
	_, _, err := sl.client.PostMessage(channelID, slackApi.MsgOptionAttachments(attachment), slackApi.MsgOptionAsUser(true))
	if err != nil {

		log.WithError(err).WithFields(log.Fields{
			"channel_id": channelID,
		}).Error("Error when trying to send post message")
		return false
	}
	log.WithFields(log.Fields{
		"channel_id": channelID,
	}).Debug("Slack message was sent")
	return true
}

// GetChannelID returns the channel id. if is it email, search the user channel id by his email
func (sl *Manager) GetChannelID(to string) (string, error) {
	if strings.HasPrefix(to, "#") {
		return to, nil
	}
	return sl.getUserIDByEmail(to)
}

// return distinct values in slice
func distinct(inputSlice []string) []string {
	keys := make(map[string]struct{})
	var list []string
	for _, entry := range inputSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = struct{}{}
			list = append(list, entry)
		}
	}
	return list
}

func replacePlaceholders(input, status, link, deployedBy string) string {
	return strings.ReplaceAll(
		strings.ReplaceAll(
			strings.ReplaceAll(input, common.StatusPlaceholder, status),
			common.LinkPlaceholder, link),
		common.DeployedByPlaceholder, deployedBy)
}
