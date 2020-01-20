package state

import (
	"context"
	"errors"
	"fmt"
	"statusbay/serverutil"
	"strings"
	"time"

	slackApi "github.com/nlopes/slack"
	log "github.com/sirupsen/logrus"
)

const (
	//UpdateSlackUserInterval interval for update slack users
	UpdateSlackUserInterval = time.Hour
)

type Slack interface {
	GetUsers() ([]slackApi.User, error)
	PostMessage(channelID string, options ...slackApi.MsgOption) (string, string, error)
}

type User struct {
	ID string
}

type SlackManager struct {
	Client Slack
	users  map[string]User
}

//NewSlack creates new slack manger
func NewSlack(client Slack) *SlackManager {

	slackManager := &SlackManager{
		Client: client,
	}
	slackManager.UpdateUsers()

	return slackManager

}

// Serve will loop of slack users
func (sl *SlackManager) Serve() serverutil.StopFunc {
	ctx, cancelFn := context.WithCancel(context.Background())
	stopped := make(chan bool)
	go func() {
		for {
			select {
			case <-time.After(UpdateSlackUserInterval):
				sl.UpdateUsers()
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
func (sl *SlackManager) UpdateUsers() bool {

	UsersData := map[string]User{}
	users, err := sl.Client.GetUsers()

	if err != nil {
		return false
	}

	for _, user := range users {
		if !user.Deleted && user.Profile.Email != "" {
			UsersData[user.Profile.Email] = User{
				ID: user.ID,
			}
		}
	}
	if len(UsersData) != len(sl.users) {
		sl.users = UsersData
		log.Info(fmt.Sprintf("Found %d slack users", len(sl.users)))
	}

	return true
}

// GetUserIDByEmail return slack user by email
func (sl *SlackManager) GetUserIDByEmail(email string) (string, error) {

	user, ok := sl.users[email]
	if ok {
		return user.ID, nil
	}

	log.WithFields(log.Fields{
		"email": email,
	}).Warn("Slack user by email not found")

	return "", errors.New("Slack user by email not found")
}

// Send will send slack notification to user
func (sl *SlackManager) Send(channelID string, attachment slackApi.Attachment) bool {

	asUserOption := slackApi.MsgOptionAsUser(true)

	_, _, err := sl.Client.PostMessage(channelID, slackApi.MsgOptionAttachments(attachment), asUserOption)
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
func (sl *SlackManager) GetChannelID(to string) (string, error) {

	if strings.HasPrefix(to, "#") {
		return to, nil
	}
	return sl.GetUserIDByEmail(to)
}
