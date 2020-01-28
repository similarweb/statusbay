package datadog

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/zorkian/go-datadog-api"
	"sync"
	"time"

	"statusbay/api/httpresponse"
	"statusbay/serverutil"

	log "github.com/sirupsen/logrus"
)

// ClientDescriber is a interface for using function in DataDog package
type ClientDescriber interface {
	QueryMetrics(from int64, to int64, query string) ([]datadog.Series, error)
}

// cacheResponse is a cache structure
type cacheResponse struct {
	ttl  time.Time
	data []httpresponse.MetricsQuery
}

// Datadog is responsible for communicate with datadog and cache storage save/cleanup
type Datadog struct {
	client               ClientDescriber
	cacheCleanupInterval time.Duration
	cacheExpiration      time.Duration
	cacheResponses       map[string]cacheResponse
	mu                   *sync.RWMutex
}

// NewDatadogManager creates a new NewDatadog
func NewDatadogManager(cacheCleanupInterval, cacheExpiration time.Duration, apiKey, appKey string, client ClientDescriber) *Datadog {

	if client == nil {
		log.Info("Creating Datadog client")
		client = datadog.NewClient(apiKey, appKey)
	}
	return &Datadog{

		client:               client,
		cacheCleanupInterval: cacheCleanupInterval,
		cacheExpiration:      cacheExpiration,
		cacheResponses:       make(map[string]cacheResponse, 0),
		mu:                   &sync.RWMutex{},
	}

}

// Serve will start listening metric request
func (dd *Datadog) Serve() serverutil.StopFunc {
	ctx, cancelFn := context.WithCancel(context.Background())
	stopped := make(chan bool)
	go func() {
		for {
			select {
			case <-time.After(dd.cacheCleanupInterval):
				dd.DeleteCacheExpired()
			case <-ctx.Done():
				log.Warn("Datatog has been shut down")
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

// GetMetric comunicate with Datadog
// All the Datadog responses saved in (in-memory) cache (can be changed in dd configuration YAML file). the reason is that
// DD API has rate limiting, and we want to decrease the count of requests
func (dd *Datadog) GetMetric(query string, from, to time.Time) ([]httpresponse.MetricsQuery, error) {

	hashKey := dd.generateMetricHash(query, from, to)
	if metrics, ok := dd.cacheResponses[hashKey]; ok {
		log.Debug("Return Datadog metric from cache")
		return metrics.data, nil
	}

	log.WithFields(log.Fields{
		"query": query,
		"from":  from,
		"to":    to,
	}).Debug("Fetch data from Datadog")

	metrics, err := dd.client.QueryMetrics(from.Unix(), to.Unix(), query)

	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"query": query,
			"from":  from,
			"to":    to,
		}).Warn("Error when trying to fetch data from datadog")
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

	dd.cacheResponses[hashKey] = cacheResponse{data: response, ttl: time.Now().Add(dd.cacheExpiration)}

	return response, nil
}

//generateMetricHash return md5 of metric
func (dd *Datadog) generateMetricHash(query string, from, to time.Time) string {

	hasher := md5.New()
	key := fmt.Sprintf("%s-%d-%d", query, from.Unix(), to.Unix())
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

//DeleteCacheExpired will check if cached metric data should be deleted
func (dd *Datadog) DeleteCacheExpired() {
	dd.mu.Lock()
	defer dd.mu.Unlock()
	data := dd.cacheResponses

	now := time.Now()

	for index, cache := range data {
		if now.After(cache.ttl) {
			log.Info("Metrics cache was deleted")
			delete(data, index)
			continue
		}
	}

	dd.cacheResponses = data

}
