package metrics

import (
	"context"
	"statusbay/api/httpresponse"
	"statusbay/api/metrics/datadog"
	"statusbay/api/metrics/prometheus"
	"statusbay/cache"
	"statusbay/config"
	"sync"
	"time"
)

// MetricManagerDescriber descrive the metric interface
type MetricManagerDescriber interface {
	Serve(ctx context.Context, wg *sync.WaitGroup)
	GetMetric(query string, from, to time.Time) ([]httpresponse.MetricsQuery, error)
}

// Load sets all metrics providers
func Load(metricsProviders *config.MetricsProvider, cache *cache.CacheManager) map[string]MetricManagerDescriber {

	providers := map[string]MetricManagerDescriber{}

	if metricsProviders == nil {
		return providers
	}
	if metricsProviders.DataDog != nil {
		providers["datadog"] = datadog.NewDatadogManager(cache, metricsProviders.DataDog.CacheExpiration, metricsProviders.DataDog.APIKey, metricsProviders.DataDog.AppKey, nil)
	}

	if metricsProviders.Prometheus != nil {
		providers["prometheus"] = prometheus.NewPrometheusManager(metricsProviders.Prometheus.Address, nil)
	}

	return providers

}
