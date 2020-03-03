package datadog

import (
	"context"
	testutil "statusbay/api/metrics/datadog/testutils"
	"sync"
	"testing"
	"time"
)

func MockDatadog(cacheExpiration, cacheCleanupInterval time.Duration) *Datadog {
	ddMockClient := testutil.NewMockDatadog()

	dd := NewDatadogManager(cacheCleanupInterval, cacheExpiration, "", "", ddMockClient)

	return dd
}

func TestGetMetric(t *testing.T) {
	var wg sync.WaitGroup
	ctx := context.Background()

	cacheExpiration := time.Second * 2
	cacheCleanupInterval := time.Second * 3
	datadog := MockDatadog(cacheExpiration, cacheCleanupInterval)
	datadog.Serve(ctx, &wg)

	from := time.Unix(1557942490, 0)
	to := time.Unix(1557942490, 0)

	testsMetrics := []struct {
		query               string
		expectedMetricCount int
		expectedMetricName  string
		expectedDatapoints  int
	}{
		{"single-metric", 1, "foo.metric.response.2xx", 2},
		{"multiple-metric", 2, "foo.metric.response.4xx", 3},
	}

	for _, test := range testsMetrics {
		t.Run(test.query, func(t *testing.T) {
			response, _ := datadog.GetMetric(test.query, from, to)

			if len(response) != test.expectedMetricCount {
				t.Fatalf("unexpected metrics count, got %d, expected %d", len(response), test.expectedMetricCount)
			}

			if response[0].Metric != test.expectedMetricName {
				t.Fatalf("unexpected metric name, got %s, expected %s", response[0].Metric, test.expectedMetricName)
			}
			if len(response[0].Points) != test.expectedDatapoints {
				t.Fatalf("unexpected metric datapoints count, got %d, expected %d", len(response[0].Points), test.expectedDatapoints)
			}
		})
	}
}

func TestCache(t *testing.T) {
	var wg sync.WaitGroup
	ctx := context.Background()

	cacheExpiration := time.Second * 2
	cacheCleanupInterval := time.Second * 3
	datadog := MockDatadog(cacheExpiration, cacheCleanupInterval)
	datadog.Serve(ctx, &wg)

	from := time.Unix(1557942490, 0)
	to := time.Unix(1557942490, 0)

	datadog.GetMetric("single-metric", from, to)
	datadog.GetMetric("multiple-metric", from, to)

	if datadog.cacheResponses.ItemCount() != 2 {
		t.Fatalf("unexpected metric cache, got %d, expected %d", datadog.cacheResponses.ItemCount(), 2)
	}

	time.Sleep(time.Second * 5)
	if datadog.cacheResponses.ItemCount() != 0 {
		t.Fatalf("unexpected clear cache, got %d, expected %d", datadog.cacheResponses.ItemCount(), 0)
	}

}
