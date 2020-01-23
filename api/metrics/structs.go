package metrics

import (
	"statusbay/api/httpresponse"
	"statusbay/api/metrics/datadog"
	"statusbay/api/metrics/prometheus"
	"statusbay/config"
	"statusbay/serverutil"
	"time"
)

// MetricManagerDescriber descrive the metric interface
type MetricManagerDescriber interface {
	Serve() serverutil.StopFunc
	GetMetric(query string, from, to time.Time) ([]httpresponse.MetricsQuery, error)
}

// Load sets all metrics providers
func Load(metricsProviders *config.MetricsProvider) map[string]MetricManagerDescriber {

	providers := map[string]MetricManagerDescriber{}

	if metricsProviders.DataDog != nil {
		providers["datadog"] = datadog.NewDatadogManager(metricsProviders.DataDog.CacheCleanupInterval, metricsProviders.DataDog.CacheExpiration, metricsProviders.DataDog.APIKey, metricsProviders.DataDog.AppKey, nil)
	}

	if metricsProviders.Prometheus != nil {
		providers["prometheus"] = prometheus.NewPrometheusManager(metricsProviders.Prometheus.Address, nil)
	}

	return providers

}
