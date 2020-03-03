package testutil

import (
	"context"
	"statusbay/api/httpresponse"
	"sync"
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

func (dd *MockMetrics) Serve(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
}
