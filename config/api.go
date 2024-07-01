package config

import (
	"os"
	"statusbay/cache"
	"statusbay/state"
	"time"

	"gopkg.in/yaml.v2"
)

// MetricsProvider struct
type MetricsProvider struct {
	DataDog    *DatadogConfig    `yaml:"datadog" env:", prefix=DATADOG_"`
	Prometheus *PrometheusConfig `yaml:"prometheus" env:", prefix=PROMETHEUS_"`
}

// AlertProvider struct
type AlertProvider struct {
	Statuscake *Statuscake `yaml:"statuscake" env:", prefix=STATUSCAKE_"`
	Pingdom    *Pingdom    `yaml:"pingdom"  env:", prefix=PINGDOM_"`
}

// DatadogConfig configuration
type DatadogConfig struct {
	APIKey          string        `yaml:"api_key" env:"API_KEY, overwrite"`
	AppKey          string        `yaml:"app_key" env:"APP_KEY, overwrite"`
	CacheExpiration time.Duration `yaml:"cache_expiration"`
}

// PrometheusConfig configuration
type PrometheusConfig struct {
	Address string `yaml:"address" env:"ADDRESS, overwrite"`
}

// Pingdom configuration
type Pingdom struct {
	Endpoint string `yaml:"endpoint" env:"ENDPOINT, overwrite"`
	Token    string `yaml:"token" env:"TOKEN, overwrite"`
}

// Statuscake configuration
type Statuscake struct {
	Endpoint string `yaml:"endpoint" env:"ENDPOINT, overwrite"`
	Username string `yaml:"username" env:"USERNAME, overwrite"`
	APIKey   string `yaml:"api_key" env:"API_KEY, overwrite"`
}

// KubernetesMarksEvents is the struct representing the events StatusBay marks
type KubernetesMarksEvents struct {
	Pod         []EventMarksConfig `yaml:"pod"`
	Replicaset  []EventMarksConfig `yaml:"replicaset"`
	Deployment  []EventMarksConfig `yaml:"deployment"`
	Demonset    []EventMarksConfig `yaml:"demonset"`
	Statefulset []EventMarksConfig `yaml:"statefulset"`
	Service     []EventMarksConfig `yaml:"service"`
	Pvc         []EventMarksConfig `yaml:"pvc"`
}

// API is holds all application configuration
type API struct {
	Log             LogConfig          `yaml:"log"`
	MySQL           *state.MySQLConfig `yaml:"mysql" env:", prefix=MYSQL_"`
	Redis           *cache.RedisConfig `yaml:"redis"`
	MetricsProvider *MetricsProvider   `yaml:"metrics"  env:", prefix=METRICS_"`
	AlertProvider   *AlertProvider     `yaml:"alerts"  env:", prefix=ALERTS_"`
	Telemetry       MetricsConfig      `yaml:"telemetry"`
}

// LoadConfigAPI will load all yaml configuration file to struct
func LoadConfigAPI(location string) (API, error) {
	config := API{}
	data, err := os.ReadFile(location)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func LoadEvents(location string) (KubernetesMarksEvents, error) {
	config := KubernetesMarksEvents{}
	data, err := os.ReadFile(location)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
