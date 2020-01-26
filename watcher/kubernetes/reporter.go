package kuberneteswatcher

import (
	"context"
	"fmt"
	"statusbay/serverutil"
	"statusbay/state"
	"strings"

	slackApi "github.com/nlopes/slack"
	log "github.com/sirupsen/logrus"
)

// DeploymentReporter defined deployment reporter message
type DeploymentReporter struct {
	// List of channels/username to send message to
	To []string

	// Deployment owner
	DeployBy string

	// Name of the deployment
	Name string

	// Status of the deployment
	Status DeploymentStatus

	//Deployment URI
	URI string
}

// ReporterManager defined reporter metadata
type ReporterManager struct {
	// Received channel when deployment started
	DeploymentStarted chan DeploymentReporter

	// Received channel when deployment deleted
	DeploymentDeleted chan DeploymentReporter

	// Received channel when deployment finish
	DeploymentFinished chan DeploymentReporter

	// Slack manger instance, will be owner to comunicate with slack
	slack *state.SlackManager

	// List of default channel report the messages
	defaultChannels []string

	// Base UI URL. will be concat to URI message from DeploymentReporter
	basestatusbayUI string
}

// NewReporter creates new reporter
func NewReporter(slack *state.SlackManager, defaultChannels []string, basestatusbayUI string) *ReporterManager {
	return &ReporterManager{
		slack:           slack,
		defaultChannels: defaultChannels,
		basestatusbayUI: basestatusbayUI,

		DeploymentStarted:  make(chan DeploymentReporter),
		DeploymentDeleted:  make(chan DeploymentReporter),
		DeploymentFinished: make(chan DeploymentReporter),
	}
}

// Serve will start to reporter listener
func (re *ReporterManager) Serve() serverutil.StopFunc {

	ctx, cancelFn := context.WithCancel(context.Background())
	stopped := make(chan bool)
	go func() {
		for {
			select {
			case request := <-re.DeploymentStarted:
				re.deploymentStarted(request)
			case request := <-re.DeploymentDeleted:
				re.deploymentDeleted(request)
			case request := <-re.DeploymentFinished:
				re.deploymentFinish(request)
			case <-ctx.Done():
				log.Warn("Reporter has been shut down")
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

// deploymentStarted will send slack message to channel/user when the deployment is started
func (re *ReporterManager) deploymentStarted(message DeploymentReporter) {

	reporters := unique(append(message.To, re.defaultChannels...))
	deployBy, err := re.slack.GetUserIDByEmail(message.DeployBy)

	if err != nil {
		deployBy = message.DeployBy
	} else {
		deployBy = fmt.Sprintf("by <@%s>", deployBy)
	}

	for _, to := range reporters {
		if to == "" {
			continue
		}
		toChannel, err := re.slack.GetChannelID(to)
		if err == nil {
			attachment := slackApi.Attachment{
				Title:   fmt.Sprintf("Latency, Hits, Events and more can be found in statusbay report. <%s|Click here> to view statusbay report", fmt.Sprintf("%s/%s", re.basestatusbayUI, message.URI)),
				Pretext: fmt.Sprintf("Kubernetes deployment started %s", deployBy),
				Color:   "#3aa3e3",
				// TODO:: add cluster + namespace name
				Fields: []slackApi.AttachmentField{
					{
						Title: "Application Name:",
						Value: message.Name,
						Short: false,
					},
				},
			}
			re.slack.Send(toChannel, attachment)

		} else {
			log.WithFields(log.Fields{
				"to": to,
			}).Debug("Slack id not found")
		}

	}

}

// deploymentDeleted will send slack message to channel/user when the deployment is deleted
func (re *ReporterManager) deploymentDeleted(message DeploymentReporter) {

	reporters := unique(append(message.To, re.defaultChannels...))
	deployBy, err := re.slack.GetUserIDByEmail(message.DeployBy)

	if err != nil {
		deployBy = message.DeployBy
	} else {
		deployBy = fmt.Sprintf("by <@%s>", deployBy)
	}

	for _, to := range reporters {
		if to == "" {
			continue
		}
		toChannel, err := re.slack.GetChannelID(to)
		if err == nil {
			attachment := slackApi.Attachment{
				Title: fmt.Sprintf("Deployment deleted %s. <%s|Click here> to view statusbay report", deployBy, fmt.Sprintf("%s/%s", re.basestatusbayUI, message.URI)),
				Color: "#3aa3e3",
				// TODO:: add cluster + namespace name
				Fields: []slackApi.AttachmentField{
					{
						Title: "Application Name:",
						Value: message.Name,
						Short: false,
					},
				},
			}
			re.slack.Send(toChannel, attachment)

		} else {
			log.WithFields(log.Fields{
				"to": to,
			}).Debug("Slack id not found")
		}

	}

}

// deploymentFinish will send slack message to channel/user when the deployment is finished
func (re *ReporterManager) deploymentFinish(message DeploymentReporter) {

	reporters := unique(append(message.To, re.defaultChannels...))
	deployBy, err := re.slack.GetUserIDByEmail(message.DeployBy)

	if err != nil {
		deployBy = message.DeployBy
	} else {
		deployBy = fmt.Sprintf("by <@%s>", deployBy)
	}

	for _, to := range reporters {
		if to == "" {
			continue
		}
		toChannel, err := re.slack.GetChannelID(to)
		if err == nil {

			color := "#25ba81"
			if message.Status == "successful" {
				color = "#25ba81"
			} else if message.Status == "cancelled" {
				color = "#ffeeba"
			} else if message.Status == "failed" {
				color = "#c84034"
			}

			attachment := slackApi.Attachment{
				Text:    fmt.Sprintf("<%s|Click here> to view statusbay report", fmt.Sprintf("%s/%s", re.basestatusbayUI, message.URI)),
				Pretext: fmt.Sprintf("Kubernetes deployment finished with status %s", strings.ToUpper(string(message.Status))),
				Color:   color,
				// TODO:: add cluster + namespace name
				Fields: []slackApi.AttachmentField{
					{
						Title: "Application Name:",
						Value: message.Name,
						Short: false,
					},
				},
			}
			re.slack.Send(toChannel, attachment)

		} else {
			log.WithFields(log.Fields{
				"to": to,
			}).Debug("Slack id not found")
		}

	}
}

// GetDeploymentURL returns ui url
func (re *ReporterManager) getDeploymentURL(name string, deployTime int64) string {
	return fmt.Sprintf("%s/deployments/%s/%d", re.basestatusbayUI, name, deployTime)
}

// unique return unique values in slice
func unique(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
