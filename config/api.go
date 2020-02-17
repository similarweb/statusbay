package config

import (
	"io/ioutil"
	"statusbay/state"
	"time"

	"gopkg.in/yaml.v2"
)

// MetricsProvider struct
type MetricsProvider struct {
	DataDog    *DatadogConfig    `yaml:"datadog"`
	Prometheus *PrometheusConfig `yaml:"prometheus"`
}

// AlertProvider struct
type AlertProvider struct {
	Statuscake *Statuscake `yaml:"statuscake"`
	Pingdom    *Pingdom    `yaml:"pingdom"`
}

// DatadogConfig configuration
type DatadogConfig struct {
	APIKey               string        `yaml:"api_key"`
	AppKey               string        `yaml:"app_key"`
	CacheCleanupInterval time.Duration `yaml:"cache_cleanup_interval"`
	CacheExpiration      time.Duration `yaml:"cache_expiration"`
}

// PrometheusConfig configuration
type PrometheusConfig struct {
	Address string `yaml:"address"`
}

// Pingdom configuration
type Pingdom struct {
	Endpoint string `yaml:"endpoint"`
	Token    string `yaml:"token"`
}

// Statuscake configuration
type Statuscake struct {
	Endpoint string `yaml:"endpoint"`
	Username string `yaml:"username"`
	APIKey   string `yaml:"api_key"`
}

// API is holds all application configuration
type API struct {
	Log             LogConfig          `yaml:"log"`
	MySQL           *state.MySQLConfig `yaml:"mysql"`
	MetricsProvider *MetricsProvider   `yaml:"metrics"`
	AlertProvider   *AlertProvider     `yaml:"alerts"`

	Telemetry MetricsConfig `yaml:"telemetry"`
}

// LoadConfigAPI will load all yaml configuration file to struct
func LoadConfigAPI(location string) (API, error) {
	config := API{}
	data, err := ioutil.ReadFile(location)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
