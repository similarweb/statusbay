package alerts

import (
	"statusbay/api/alerts/pingdom"
	"statusbay/api/alerts/statuscake"
	"statusbay/api/httpresponse"
	"statusbay/config"
	"statusbay/request"
	"time"
)

// AlertsManagerDescriber descrive the alerts providers integrations
type AlertsManagerDescriber interface {
	GetAlertByTags(tags string, from, to time.Time) ([]httpresponse.CheckResponse, error)
}

// Load all given alerts providers
func Load(alertProviders *config.AlertProvider) map[string]AlertsManagerDescriber {

	providers := map[string]AlertsManagerDescriber{}

	if alertProviders == nil {
		return providers
	}

	if alertProviders.Statuscake != nil {
		HTTPClient := request.NewHTTPClient()
		client := statuscake.NewClient(alertProviders.Statuscake.Endpoint, alertProviders.Statuscake.Username, alertProviders.Statuscake.APIKey, HTTPClient)
		providers["statuscake"] = statuscake.NewStatuscakeManager(client)
	}

	if alertProviders.Pingdom != nil {
		HTTPClient := request.NewHTTPClient()
		client := pingdom.NewClient(alertProviders.Pingdom.Endpoint, alertProviders.Pingdom.Token, HTTPClient)
		providers["pingdom"] = pingdom.NewPingdomManager(client)
	}

	return providers

}
