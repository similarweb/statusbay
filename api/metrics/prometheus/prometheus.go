package prometheus

import (
	"context"
	"strconv"
	"sync"
	"time"

	"statusbay/api/httpresponse"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	log "github.com/sirupsen/logrus"
)

// ClientDescriber is a interface for using function in DataDog package
type ClientDescriber interface {
	QueryRange(ctx context.Context, query string, r v1.Range) (model.Value, v1.Warnings, error)
}

// Prometheus is responsible for communicate with datadog and cache storage save/cleanup
type Prometheus struct {
	api    ClientDescriber
	logger *log.Entry
}

// NewPrometheusManager creates a new NewDatadog
func NewPrometheusManager(address string, v1api ClientDescriber) *Prometheus {

	if v1api == nil {
		log.Info("initializing Prometheus client")
		client, err := api.NewClient(api.Config{
			Address: address,
		})
		if err != nil {
			log.WithError(err).Fatal("could not create Prometheus client")
		}

		v1api = v1.NewAPI(client)
	}
	return &Prometheus{
		api:    v1api,
		logger: log.WithField("metric_engine", "prometheus"),
	}

}

// Serve will start listening metric request
func (pm *Prometheus) Serve(ctx context.Context, wg *sync.WaitGroup) {

	go func() {
		for {
			select {
			case <-ctx.Done():
				pm.logger.Warn("Prometheus has been shut down")
				wg.Done()
				return
			}
		}
	}()

}

// GetMetric communicates with Prometheus
func (pm *Prometheus) GetMetric(query string, from, to time.Time) ([]httpresponse.MetricsQuery, error) {
	pm.logger.WithFields(log.Fields{
		"query": query,
		"from":  from,
		"to":    to,
	}).Debug("fetching metrics")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	r := v1.Range{
		Start: from,
		End:   to,
		Step:  time.Duration(time.Second * 60),
	}

	val, _, err := pm.api.QueryRange(ctx, query, r)
	if err != nil {
		pm.logger.WithError(err).WithFields(log.Fields{
			"query": query,
			"from":  from,
			"to":    to,
		}).Warn("could not get metrics")
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
				pm.logger.WithError(err).Warn("could not convert datapoint to float")
				continue
			}

			points = append(points, [2]float64{float64(dp.Timestamp.Unix()), dpf})
		}

		metricData.Points = points

		response = append(response, metricData)
	}

	return response, nil
}
