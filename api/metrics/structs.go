package metrics

import (
	"statusbay/serverutil"
	"statusbay/api/httpresponse"
	"time"
)

// MetricManagerDescriber descrive the metric interface
type MetricManagerDescriber interface {
	Serve() serverutil.StopFunc
	GetMetric(query string, from, to time.Time) ([]httpresponse.MetricsQuery, error)
}
