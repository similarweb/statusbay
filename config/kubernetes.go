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

	// MaxDeploymentTime default watch deployment time. Can be override by `ProgressDeadlineSeconds`
	MaxDeploymentTime time.Duration `yaml:"max_deployment_time"`

	// SaveDeploymentHistoryDuration defind period on time to save deployment history in memory
	SaveDeploymentHistoryDuration time.Duration `yaml:"deployment_history"`

	// CheckFinishDelay defind time to wait until starting check the status of the deployment
	CheckFinishDelay time.Duration `yaml:"check_finish_delay"`

	// CollectDataAfterDeploymentFinish defind how many time to continue collect deployment event after deployment finished
	CollectDataAfterDeploymentFinish time.Duration `yaml:"collect_data_after_deployment_finish"`
}

// EventMarksConfig is defind how the mark event will look
type EventMarksConfig struct {
	Pattern      string   `yaml:"pattern"`
	Descriptions []string `yaml:"descriptions"`
}

// Kubernetes is holds all application configuration
type Kubernetes struct {
	LogLevel string             `yaml:"log_level"`
	MySQL    *MySQLConfig       `yaml:"mysql"`
	Slack    *SlackConfig       `yaml:"slack"`
	UI       *UIConfig          `yaml:"ui"`
	Applies  *KubernetesApplies `yaml:"Applies"`
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
