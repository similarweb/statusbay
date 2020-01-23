package testutil

import (
	"statusbay/api/alerts"
	"statusbay/api/httpresponse"
	"time"
)

type MockAlerts struct {
}

func NewMockAlerts() *MockAlerts {
	return &MockAlerts{}
}

func NewMultipleMockAlerts() map[string]alerts.AlertsManagerDescriber {

	providers := map[string]alerts.AlertsManagerDescriber{}

	providers["foo"] = NewMockAlerts()
	providers["foo2"] = NewMockAlerts()

	return providers

}

func (m *MockAlerts) GetAlertByTags(tags string, from, to time.Time) ([]httpresponse.CheckResponse, error) {
	response := []httpresponse.CheckResponse{
		{
			ID: 1, URL: "foo.com", Name: "foo", Periods: []httpresponse.PeriodsResponse{
				{Status: "up", StartUnix: 1, EndUnix: 1},
			},
		},
	}
	return response, nil
}
