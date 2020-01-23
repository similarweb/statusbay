package prometheus

import (
	"context"
	"strconv"
	"time"

	"statusbay/api/httpresponse"
	"statusbay/serverutil"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	log "github.com/sirupsen/logrus"
)

// ClientDescriber is a interface for using function in DataDog package
type ClientDescriber interface {
	QueryRange(ctx context.Context, query string, r v1.Range) (model.Value, error)
}

// Prometheus is responsible for communicate with datadog and cache storage save/cleanup
type Prometheus struct {
	api ClientDescriber
}

// NewPrometheusManager creates a new NewDatadog
func NewPrometheusManager(address string, v1api ClientDescriber) *Prometheus {

	if v1api == nil {
		log.Info("Creating Prometheus client")
		client, err := api.NewClient(api.Config{
			Address: address,
		})
		if err != nil {
			log.WithError(err).Fatal("Error creating client")
		}

		v1api = v1.NewAPI(client)
	}
	return &Prometheus{
		api: v1api,
	}

}

// Serve will start listening metric request
func (pm *Prometheus) Serve() serverutil.StopFunc {
	ctx, cancelFn := context.WithCancel(context.Background())
	stopped := make(chan bool)
	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Warn("Prometheus has been shut down")
				stopped <- true
				return
			}
		}
	}()
	return func() {
		cancelFn()
		<-stopped
	}
}

// GetMetric communicates with Prometheus
func (pm *Prometheus) GetMetric(query string, from, to time.Time) ([]httpresponse.MetricsQuery, error) {
	log.WithFields(log.Fields{
		"query": query,
		"from":  from,
		"to":    to,
	}).Debug("Fetch data from Prometheus")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	r := v1.Range{
		Start: from,
		End:   to,
		Step:  time.Duration(time.Second * 60),
	}

	val, err := pm.api.QueryRange(ctx, query, r)
	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"query": query,
			"from":  from,
			"to":    to,
		}).Warn("Error when trying to fetch data from Prometheus")
		return nil, err
	}

	if val == nil {
		return nil, nil
	}

	response := []httpresponse.MetricsQuery{}
	for _, metric := range val.(model.Matrix) {
		metricData := httpresponse.MetricsQuery{}

		metricName := metric.Metric.String()
		// If metric name is empty, fallback to query (requested metric)
		if metricName == "{}" {
			metricName = query
		}

		metricData.Metric = metricName

		points := []httpresponse.DataPoint{}
		for _, dp := range metric.Values {
			dpf, err := strconv.ParseFloat(dp.Value.String(), 64)
			if err != nil {
				log.WithError(err).Warn("could not convert datapoint to float")
				continue
			}

			points = append(points, [2]float64{float64(dp.Timestamp.Unix()), dpf})
		}

		metricData.Points = points

		response = append(response, metricData)
	}

	return response, nil
}
