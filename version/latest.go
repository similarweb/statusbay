package version

import (
	"context"
	"errors"
	"fmt"
	"time"

	notifier "github.com/similarweb/client-notifier"
	log "github.com/sirupsen/logrus"
)

var (
	// MainVersion is the main version number that is being run at the moment.
	MainVersion = "0.0.2"
)

// VersionDescriptor describe the version struct
type VersionDescriptor interface {
	Get() (*notifier.Response, error)
}

// Version struct
type Version struct {
	duration        time.Duration
	params          *notifier.UpdaterParams
	requestSettings notifier.RequestSetting
	response        *notifier.Response
}

// NewVersion create new instance of version
func NewVersion(ctx context.Context, component string, duration time.Duration) *Version {

	params := &notifier.UpdaterParams{
		Application: "statusbay",
		Component:   component,
		Version:     MainVersion,
	}

	version := &Version{
		params:   params,
		duration: duration,
	}

	response, err := notifier.Get(version.params, version.requestSettings)
	version.printResults(response, err)
	version.interval(ctx)

	return version

}

// interval is a periodic version check
func (v *Version) interval(ctx context.Context) {
	notifier.GetInterval(ctx, v.params, v.duration, v.printResults, v.requestSettings)
}

// printResults print the notifier response to the logger
func (v *Version) printResults(n *notifier.Response, err error) {

	if err != nil {
		log.WithError(err).Error(fmt.Sprintf("failed to get Statusbat latest version"))
		return
	}
	v.response = n
	if n.Outdated {
		log.Error(fmt.Sprintf("Newer Statusbay version available. latest version %s, current version %s", n.LatestVersion, v.params.Version))
	}

	for _, notification := range n.Notifications {
		log.Error(notification.Message)
	}

}

// Get returns the notifier response
func (v *Version) Get() (*notifier.Response, error) {

	if v.response == nil {
		return nil, errors.New("Version response not found")
	}
	return v.response, nil
}
