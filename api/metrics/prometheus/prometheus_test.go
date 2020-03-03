package prometheus

import (
	"context"
	testutil "statusbay/api/metrics/prometheus/testutils"
	"sync"
	"testing"
	"time"
)

func MockPrometheus() *Prometheus {
	pmMockClient := testutil.NewMockPrometheus()

	pm := NewPrometheusManager("", pmMockClient)

	return pm
}

func TestGetMetric(t *testing.T) {
	var wg sync.WaitGroup
	ctx := context.Background()

	prometheus := MockPrometheus()
	prometheus.Serve(ctx, &wg)

	from := time.Unix(1557942490, 0)
	to := time.Unix(1557942490, 0)

	testsMetrics := []struct {
		query               string
		expectedMetricCount int
		expectedMetricName  string
		expectedDatapoints  int
	}{
		{"single-metric", 1, "prometheus_http_requests_total{code=\"200\", handler=\"/alerts\", instance=\"localhost:9090\", job=\"prometheus\"}", 4},
		{"summed-metric", 1, "summed-metric", 4},
		{"multiple-metrics", 9, "prometheus_http_requests_total{code=\"200\", handler=\"/alerts\", instance=\"localhost:9090\", job=\"prometheus\"}", 4},
	}

	for _, test := range testsMetrics {
		t.Run(test.query, func(t *testing.T) {
			response, _ := prometheus.GetMetric(test.query, from, to)

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
