// +build !race

package slack

import (
	"github.com/nlopes/slack"
	"testing"
	"time"
)

func TestServe(t *testing.T) {
	t.Run("checking if the cancel function works", func(t *testing.T) {
		mockClient := &MockApiClient{
			users: []slack.User{
				{
					Profile: slack.UserProfile{Email: "user1"},
				},
			},
		}
		slackManager := Manager{
			client:      mockClient,
			emailToUser: map[string]string{},
		}

		UpdateSlackUserInterval = time.Nanosecond
		cancelFunc := slackManager.Serve()

		mockClient.users = nil
		time.Sleep(time.Millisecond)

		cancelFunc()

		if slackManager.emailToUser == nil {
			t.Errorf("expected the goroutine to stop running and stop updating the users")
		}
	})

	t.Run("check if user list is continuously being updated", func(t *testing.T) {
		mockClient := &MockApiClient{}
		slackManager := Manager{
			client:      mockClient,
			emailToUser: map[string]string{},
		}

		UpdateSlackUserInterval = time.Nanosecond

		// start with no users
		mockClient.users = nil

		cancelFunc := slackManager.Serve()
		defer cancelFunc()

		time.Sleep(time.Millisecond)
		if len(slackManager.emailToUser) != 0 {
			t.Errorf("expected to have no users available")
		}

		// add one user to the returned from the api
		mockClient.users = append(mockClient.users, slack.User{
			Profile: slack.UserProfile{Email: "user1"},
		})
		time.Sleep(time.Millisecond)
		if len(slackManager.emailToUser) != 1 {
			t.Errorf("expected to get exactly one user")
		}

		// add 2 more users
		mockClient.users = append(mockClient.users,
			slack.User{Profile: slack.UserProfile{Email: "user2"}},
			slack.User{Profile: slack.UserProfile{Email: "user3"}})
		time.Sleep(time.Millisecond)
		if len(slackManager.emailToUser) != 3 {
			t.Errorf("expected to get exactly three users")
		}

		// return no users
		mockClient.users = nil
		time.Sleep(time.Millisecond)
		if len(slackManager.emailToUser) != 0 {
			t.Errorf("expected to have no users available")
		}
	})
}
