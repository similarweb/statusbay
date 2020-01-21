package testutil

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"

	slackApi "github.com/nlopes/slack"
)

type MockPostMessageRequest struct {
	ChannelID string
	Options   []slackApi.MsgOption
}

type MockSlack struct {
	PostMessageRequest []MockPostMessageRequest
}

func NewMockSlack() *MockSlack {
	return &MockSlack{
		PostMessageRequest: []MockPostMessageRequest{},
	}
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
	request := MockPostMessageRequest{
		ChannelID: channelID,
		Options:   options,
	}
	m.PostMessageRequest = append(m.PostMessageRequest, request)
	return "", "", nil
}
