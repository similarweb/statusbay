package testutil

import (
	"statusbay/api/httpresponse"
	"statusbay/serverutil"
	"time"
)

type MockMetrics struct {
}

func NewMockMetrics() *MockMetrics {
	return &MockMetrics{}
}

func (m *MockMetrics) GetMetric(query string, from, to time.Time) ([]httpresponse.MetricsQuery, error) {
	response := []httpresponse.MetricsQuery{
		{Metric: "foo", Points: []httpresponse.DataPoint{
			{1, 2},
		}},
	}
	return response, nil
}

func (dd *MockMetrics) Serve() serverutil.StopFunc {

	return func() {}
}
