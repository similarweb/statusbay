package statuscake

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
	PageURL = "https://app.statuscake.com/UptimeStatus.php?tid=%d"

	// ParallelCheckRequests defind the number of request to Statuscake at the same time
	ParallelCheckRequests = 20
)

// Statuscake struct
type Statuscake struct {
	client ClientDescriber
	mutex  *sync.RWMutex
}

// NewStatuscakeManager create new Statuscake instance
func NewStatuscakeManager(client ClientDescriber) *Statuscake {
	log.Info("initializing StatusCake manager")

	return &Statuscake{
		client: client,
		mutex:  &sync.RWMutex{},
	}
}

// GetAlertByTags return the existing alerts from statuscake
func (sc *Statuscake) GetAlertByTags(tags string, from, to time.Time) ([]httpresponse.CheckResponse, error) {
	v := url.Values{}
	v.Set("tags", tags)

	checkResponses := []httpresponse.CheckResponse{}
	lg := log.WithFields(log.Fields{
		"tags": tags,
		"from": from,
		"to":   to,
	})
	tests, err := sc.client.GetTests(v)

	if err != nil {
		lg.WithError(err).Error("could not get tests from StatusCake")
		return checkResponses, err
	}

	lg.WithField("tests", len(tests)).Debug("received checks from StatusCake")

	var wg = sync.WaitGroup{}
	GoroutinesRequests := make(chan struct{}, ParallelCheckRequests)
	for _, test := range tests {
		GoroutinesRequests <- struct{}{}
		wg.Add(1)
		go func(test *Test) {
			defer wg.Done()
			alertEvents, err := sc.client.Periods(test.TestID)
			if err != nil {
				lg.WithError(err).WithField("test_id", test.TestID).Info("could not get uptime periods for check")
				return
			}

			check := httpresponse.CheckResponse{
				ID:      test.TestID,
				Name:    test.WebsiteName,
				URL:     fmt.Sprintf(PageURL, test.TestID),
				Periods: make([]httpresponse.PeriodsResponse, 0),
			}

			for _, event := range alertEvents {
				if event.StartUnix > from.Unix() && event.StartUnix < to.Unix() {
					check.Periods = append(check.Periods, httpresponse.PeriodsResponse{
						Status:    event.Status,
						StartUnix: event.StartUnix,
						EndUnix:   event.EndUnix,
					})
				}

			}
			sc.mutex.Lock()
			checkResponses = append(checkResponses, check)
			sc.mutex.Unlock()
			<-GoroutinesRequests
		}(test)

	}
	wg.Wait()
	return checkResponses, nil
}
