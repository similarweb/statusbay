// +build !race

package slack

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/nlopes/slack"
)

func TestServe(t *testing.T) {
	var wg *sync.WaitGroup
	ctx := context.Background()
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
		slackManager.Serve(ctx, wg)

		mockClient.users = nil
		time.Sleep(time.Millisecond)

		if slackManager.emailToUser == nil {
			t.Error("expected the goroutine to stop running and stop updating the users")
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

		slackManager.Serve(ctx, wg)

		time.Sleep(time.Millisecond)
		if len(slackManager.emailToUser) != 0 {
			t.Error("expected to have no users available")
		}

		// add one user to the returned from the api
		mockClient.users = append(mockClient.users, slack.User{
			Profile: slack.UserProfile{Email: "user1"},
		})
		time.Sleep(time.Millisecond)
		if len(slackManager.emailToUser) != 1 {
			t.Error("expected to get exactly one user")
		}

		// add 2 more users
		mockClient.users = append(mockClient.users,
			slack.User{Profile: slack.UserProfile{Email: "user2"}},
			slack.User{Profile: slack.UserProfile{Email: "user3"}})
		time.Sleep(time.Millisecond)
		if len(slackManager.emailToUser) != 3 {
			t.Error("expected to get exactly three users")
		}

		// return no users
		mockClient.users = nil
		time.Sleep(time.Millisecond)
		if len(slackManager.emailToUser) != 0 {
			t.Error("expected to have no users available")
		}
	})
}
