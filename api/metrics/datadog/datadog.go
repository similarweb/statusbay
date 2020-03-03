package datadog

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/zorkian/go-datadog-api"

	"statusbay/api/httpresponse"
	"statusbay/cache"

	log "github.com/sirupsen/logrus"
)

// ClientDescriber is a interface for using function in DataDog package
type ClientDescriber interface {
	QueryMetrics(from int64, to int64, query string) ([]datadog.Series, error)
}

// Datadog is responsible for communicate with datadog and cache storage save/cleanup
type Datadog struct {
	client          ClientDescriber
	cache           *cache.CacheManager
	cacheExpiration time.Duration
	mu              *sync.RWMutex
	logger          *log.Entry
}

// NewDatadogManager creates a new NewDatadog
func NewDatadogManager(cache *cache.CacheManager, cacheExpiration time.Duration, apiKey, appKey string, client ClientDescriber) *Datadog {

	if client == nil {
		log.Info("initializing Datadog client")
		client = datadog.NewClient(apiKey, appKey)
	}
	return &Datadog{
		client:          client,
		cache:           cache,
		cacheExpiration: cacheExpiration,
		mu:              &sync.RWMutex{},
		logger:          log.WithField("metric_engine", "datadog"),
	}

}

// Serve will start listening metric request
func (dd *Datadog) Serve(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		for {
			select {
			case <-ctx.Done():
				dd.logger.Warn("Datatog has been shut down")
				wg.Done()
				return
			}
		}
	}()
}

// GetMetric comunicate with Datadog
// All the Datadog responses saved in (in-memory) cache (can be changed in dd configuration YAML file). the reason is that
// DD API has rate limiting, and we want to decrease the count of requests
func (dd *Datadog) GetMetric(query string, from, to time.Time) ([]httpresponse.MetricsQuery, error) {

	hashKey := dd.generateMetricHash(query, from, to)

	cacheMetrics, err := dd.cache.Client.Get(hashKey)
	if err == nil && cacheMetrics != "" {
		response := []httpresponse.MetricsQuery{}
		err = json.Unmarshal([]byte(cacheMetrics), &response)

		if err == nil {
			dd.logger.Info("found metric in cache")
			return response, nil
		}
		dd.logger.WithError(err).WithFields(log.Fields{
			"query": query,
			"from":  from,
			"to":    to,
		}).Error("cache Unmarshal failed")
	}

	metrics, err := dd.client.QueryMetrics(from.Unix(), to.Unix(), query)

	if err != nil {
		dd.logger.WithError(err).WithFields(log.Fields{
			"query": query,
			"from":  from,
			"to":    to,
		}).Warn("could not get metrics")
		return nil, err
	}

	var response []httpresponse.MetricsQuery
	for _, metric := range metrics {
		metricData := httpresponse.MetricsQuery{}
		metricData.Metric = metric.GetDisplayName()
		var points []httpresponse.DataPoint
		for _, point := range metric.Points {
			points = append(points, [2]float64{*point[0], *point[1]})
		}
		metricData.Points = points

		response = append(response, metricData)
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		dd.logger.WithError(err).Error("could not serialize response")
		return response, nil
	}

	err = dd.cache.Client.Set(hashKey, string(jsonResponse), dd.cacheExpiration)
	if err != nil {
		dd.logger.WithError(err).Error("could not save metric response")
	}

	return response, nil
}

//generateMetricHash return md5 of metric
func (dd *Datadog) generateMetricHash(query string, from, to time.Time) string {

	hasher := md5.New()
	key := fmt.Sprintf("%s-%d-%d", query, from.Unix(), to.Unix())
	hasher.Write([]byte(key))
	val := hex.EncodeToString(hasher.Sum(nil))
	return fmt.Sprintf("datadog-metrics-%s", val)
}
