package pingdom

import (
	"fmt"
	"net/url"
	"statusbay/api/httpresponse"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	// PageURL define the direct link to check
	PageURL = "https://my.pingdom.com/app/reports/uptime#check=%d"

	// ParallelCheckRequests defind the number of request to Pingdom at the same time
	ParallelCheckRequests = 20
)

// Pingdom struct
type Pingdom struct {
	client ClientDescriber
}

// NewPingdomManager create new Pingdom instance
func NewPingdomManager(client ClientDescriber) *Pingdom {
	log.Info("Init Pingdom manager")
	return &Pingdom{
		client: client,
	}
}

// GetAlertByTags return the existing alerts ... from tags and time.. todo::....
func (pi *Pingdom) GetAlertByTags(tags string, from, to time.Time) ([]httpresponse.CheckResponse, error) {
	v := url.Values{}
	v.Set("tags", tags)

	lg := log.WithFields(log.Fields{
		"tags": tags,
		"from": from,
		"to":   to,
	})

	checkResponses := []httpresponse.CheckResponse{}

	checks, err := pi.client.GetChecks(v)
	if err != nil {
		lg.WithError(err).Error("Error when trying to get Pingdom checks")
		return checkResponses, nil
	}

	lg.WithField("checks", len(checks.Checks)).Debug("Alerts was found")

	// Prepare query string filter
	resultQueryString := url.Values{}
	resultQueryString.Set("from", fmt.Sprintf("%d", from.Unix()))
	resultQueryString.Set("to", fmt.Sprintf("%d", to.Unix()))

	var wg = sync.WaitGroup{}
	GoroutinesRequests := make(chan struct{}, ParallelCheckRequests)

	for _, check := range checks.Checks {

		GoroutinesRequests <- struct{}{}
		wg.Add(1)
		go func(check Check) {
			defer wg.Done()
			checkStatus, err := pi.client.GetCheckSummaryOutage(check.ID, resultQueryString)
			if err != nil {
				if err != nil {
					lg.WithError(err).WithField("check_id", check.ID).Info("Failed to call summary outage check")
					return
				}
			}

			checkData := httpresponse.CheckResponse{
				ID:      check.ID,
				Name:    check.Name,
				URL:     fmt.Sprintf(PageURL, check.ID),
				Periods: make([]httpresponse.PeriodsResponse, 0),
			}

			for _, su := range checkStatus.Summary.States {
				checkData.Periods = append(checkData.Periods, httpresponse.PeriodsResponse{
					Status:    su.Status,
					StartUnix: su.TimeFrom,
					EndUnix:   su.TimeTo,
				})
			}

			checkResponses = append(checkResponses, checkData)
			<-GoroutinesRequests
		}(check)

	}
	wg.Wait()
	return checkResponses, nil
}
