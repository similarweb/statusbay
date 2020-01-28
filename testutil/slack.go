package testutil

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"

	slackApi "github.com/nlopes/slack"
)

type MockSlack struct {
	messagesSent []string
}

func NewMockSlack() *MockSlack {
	return &MockSlack{}
}

func (m *MockSlack) GetUsers() ([]slackApi.User, error) {
	_, filename, _, _ := runtime.Caller(0)
	currentFolderPath := filepath.Dir(filename)
	reader, err := ioutil.ReadFile(fmt.Sprintf("%s/data/slack/users.json", currentFolderPath))

	if err != nil {
		return []slackApi.User{}, err
	}
	var usersResponse []slackApi.User
	var users []slackApi.User
	err = json.Unmarshal(reader, &users)
	if err != nil {
		return nil, err
	}

	for idx := range users {
		usersResponse = append(usersResponse, users[idx])
	}
	return usersResponse, nil

}

func (m *MockSlack) PostMessage(channelID string, options ...slackApi.MsgOption) (string, string, error) {
	m.messagesSent = append(m.messagesSent, channelID)
	return "", "", nil
}
