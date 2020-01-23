package httpresponse

// MetricsQuery define the metrics response
type MetricsQuery struct {
	Metric string
	Points []DataPoint `json:"Points,omitempty"`
}

// DataPoint is a tuple of [UNIX timestamp, value]. This has to use floats
type DataPoint [2]float64
