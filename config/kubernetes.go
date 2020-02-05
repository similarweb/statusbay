package config

import (
	"io/ioutil"
	"time"

	"gopkg.in/yaml.v2"
)

type DashboardConfig struct {
	Service string `yaml:"service"`
}

// UIConfig configuration
type UIConfig struct {
	BaseURL string `yaml:"base_url"`
}

// KubernetesApplies configuration
type KubernetesApplies struct {
	// SaveInterval storage save interval
	SaveInterval time.Duration `yaml:"save_interval"`

	// MaxApplyTime default watch duration time. Can be override by `ProgressDeadlineSeconds`
	MaxApplyTime time.Duration `yaml:"max_apply_time"`

	// CheckFinishDelay defind time to wait until starting check the status of the apply
	CheckFinishDelay time.Duration `yaml:"check_finish_delay"`

	// CollectDataAfterApplyFinish defind how many time to continue collect apply events
	CollectDataAfterApplyFinish time.Duration `yaml:"collect_data_after_apply_finish"`
}

// EventMarksConfig is defind how the mark event will look
type EventMarksConfig struct {
	Pattern      string   `yaml:"pattern"`
	Descriptions []string `yaml:"descriptions"`
}

// Kubernetes is holds all application configuration
type Kubernetes struct {
	Log     LogConfig          `yaml:"log"`
	MySQL   *MySQLConfig       `yaml:"mysql"`
	Slack   *SlackConfig       `yaml:"slack"`
	UI      *UIConfig          `yaml:"ui"`
	Applies *KubernetesApplies `yaml:"applies"`
}

type KubernetesMarksEvents struct {
	Pod        []EventMarksConfig `yaml:"pod"`
	Replicaset []EventMarksConfig `yaml:"replicaset"`
	Deployment []EventMarksConfig `yaml:"deployment"`
}

// LoadKubernetesConfig will load all yaml configuration file to struct
func LoadKubernetesConfig(location string) (Kubernetes, error) {
	config := Kubernetes{}
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

func LoadKubernetesMarksConfig(location string) (KubernetesMarksEvents, error) {
	config := KubernetesMarksEvents{}
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
