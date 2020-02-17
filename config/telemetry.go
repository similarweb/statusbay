package config

import (
	"time"

	"github.com/armon/go-metrics"
	"github.com/armon/go-metrics/datadog"
	"github.com/armon/go-metrics/prometheus"
)

const (
	_minFlushIntervalSec  int64 = 10
	_defaultMetricsPrefix       = "statusbay."
)

type MetricsConfig struct {

	// FlushIntervalSec is the number of seconds we want to wait between metric flushes to sinks
	// optional. defaults to the value in _minFlushIntervalSec  if not provided
	FlushIntervalSec int64 `yaml:"flush_interval,omitempty"`

	// AllowedPrefixes is a list of filter rules to apply for allowing metrics by prefix
	AllowedPrefixes []string `yaml:"allowed_prefixes,omitempty"`

	// BlockedPrefixes is a list of filter rules to apply for blocking metrics by prefix
	BlockedPrefixes []string `yaml:"blocked_prefixes,omitempty"`

	// DisableHostname will disable hostname prefixing for all metrics.
	DisableHostname bool `yaml:"disable_hostname,omitempty"`

	// FilterDefault indicates whether we want to allow metrics by default
	FilterDefault bool `yaml:"filter_default,omitempty"`

	// MetricsPrefix is the prefix used to write stats values to.
	// optional. defaults to the value in _defaultMetricsPrefix if not provided
	MetricsPrefix string `yaml:"metrics_prefix,omitempty"`

	// DogStatsdAddr is the address of a dogstatsd instance
	DogstatsdAddr string `yaml:"dogstatsd_addr,omitempty"`

	// DatadogTags are tags that should be sent with each packet to dogstatsd
	// A list of strings where each string looks like this "my_tag_name:my_tag_value"
	DatadogTags []string `yaml:"datadog_tags,omitempty"`

	// PrometheusRetentionTime is the retention time for prometheus metrics in seconds.
	// a non-positive value disable Prometheus support. It is highly advised to put large values
	// days/hours/at least the interval between prometheus requests.
	PrometheusRetentionTimeSeconds int64 `yaml:"prometheus_retention_time_sec,omitempty"`

	// StatsdAddr is the address of a statsd instance
	StatsdAddr string `yaml:"statsd_address,omitempty"`

	// StatsiteAddr is the address of a statsite instance
	StatsiteAddr string `yaml:"statsite_address,omitempty"`
}

func dogstatdSink(cfg MetricsConfig, hostname string) (metrics.MetricSink, error) {
	if cfg.DogstatsdAddr == "" {
		return nil, nil
	}

	if sink, err := datadog.NewDogStatsdSink(cfg.DogstatsdAddr, hostname); err != nil {
		return nil, err
	} else {
		sink.SetTags(cfg.DatadogTags)
		return sink, nil
	}
}

func prometheusSink(cfg MetricsConfig, _ string) (metrics.MetricSink, error) {
	if cfg.PrometheusRetentionTimeSeconds < 1 {
		return nil, nil
	}
	prometheusOpts := prometheus.PrometheusOpts{
		Expiration: time.Duration(cfg.PrometheusRetentionTimeSeconds) * time.Second,
	}

	return prometheus.NewPrometheusSinkFrom(prometheusOpts)
}

func statsiteSink(cfg MetricsConfig, _ string) (metrics.MetricSink, error) {
	if cfg.StatsiteAddr == "" {
		return nil, nil
	}
	return metrics.NewStatsiteSink(cfg.StatsiteAddr)
}

func statsdSink(cfg MetricsConfig, _ string) (metrics.MetricSink, error) {
	if cfg.StatsdAddr == "" {
		return nil, nil
	}
	return metrics.NewStatsdSink(cfg.StatsdAddr)
}

// InitMetricAggregator sets up go-metrics using the provided MetricsConfig
func InitMetricAggregator(cfg MetricsConfig) (err error) {
	flushIntervalSec := _minFlushIntervalSec
	if cfg.FlushIntervalSec > flushIntervalSec { // overwrite if lower than minimum allowed value
		flushIntervalSec = _minFlushIntervalSec
	}

	memSink := metrics.NewInmemSink(time.Duration(flushIntervalSec)*time.Second, time.Minute)
	metrics.DefaultInmemSignal(memSink)

	metricsPrefix := _defaultMetricsPrefix
	if cfg.MetricsPrefix != "" {
		metricsPrefix = cfg.MetricsPrefix
	}

	metricsConf := metrics.DefaultConfig(metricsPrefix)
	metricsConf.AllowedPrefixes = cfg.AllowedPrefixes
	metricsConf.BlockedPrefixes = cfg.BlockedPrefixes
	metricsConf.EnableHostname = !cfg.DisableHostname
	metricsConf.FilterDefault = cfg.FilterDefault

	var sinks metrics.FanoutSink

	addSink := func(name string, sinkFunc func(MetricsConfig, string) (metrics.MetricSink, error)) error {
		if sink, err := sinkFunc(cfg, metricsConf.HostName); err != nil {
			return err
		} else if sink != nil {
			sinks = append(sinks, sink)
		}
		return nil
	}
	if err = addSink("dogstatd", dogstatdSink); err != nil {
		return err
	}
	if err = addSink("prometheus", prometheusSink); err != nil {
		return err
	}
	if err = addSink("statsite", statsiteSink); err != nil {
		return err
	}
	if err = addSink("statsd", statsdSink); err != nil {
		return err
	}
	if len(sinks) > 0 {
		sinks = append(sinks, memSink)
		_, err = metrics.NewGlobal(metricsConf, sinks)
	} else {
		metricsConf.EnableHostname = false
		_, err = metrics.NewGlobal(metricsConf, memSink)
	}

	return
}
