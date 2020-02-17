package config

import (
	"io/ioutil"
	notifierCommon "statusbay/notifiers/common"
	notifierLoader "statusbay/notifiers/load"
	"statusbay/state"
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

// EventMarksConfig is defined how the mark event will look
type EventMarksConfig struct {
	Pattern      string   `yaml:"pattern"`
	Descriptions []string `yaml:"descriptions"`
}

// Kubernetes is holds all application configuration
type Kubernetes struct {
	ClusterName     string                      `yaml:"cluster_name"`
	Log             LogConfig                   `yaml:"log"`
	MySQL           *state.MySQLConfig          `yaml:"mysql"`
	NotifierConfigs notifierCommon.ConfigByName `yaml:"notifiers"`
	UI              *UIConfig                   `yaml:"ui"`
	Applies         *KubernetesApplies          `yaml:"applies"`

	Telemetry MetricsConfig `yaml:"telemetry"`

	registeredNotifiers []notifierCommon.Notifier
}

func (k *Kubernetes) BuildNotifiers() (registeredNotifiers []notifierCommon.Notifier, err error) {
	if k.NotifierConfigs != nil {
		notifierLoader.RegisterNotifiers()

		if registeredNotifiers, err = notifierLoader.Load(k.NotifierConfigs, k.UI.BaseURL); err != nil {
			return
		}
		k.registeredNotifiers = registeredNotifiers
	} else {
		registeredNotifiers = []notifierCommon.Notifier{}
	}
	return
}

type KubernetesMarksEvents struct {
	Pod         []EventMarksConfig `yaml:"pod"`
	Replicaset  []EventMarksConfig `yaml:"replicaset"`
	Deployment  []EventMarksConfig `yaml:"deployment"`
	Demonset    []EventMarksConfig `yaml:"demonset"`
	Statefulset []EventMarksConfig `yaml:"statefulset"`
}

// LoadKubernetesConfig will load all yaml configuration file to struct
func LoadKubernetesConfig(path string) (config Kubernetes, err error) {
	var (
		data []byte
	)
	if data, err = ioutil.ReadFile(path); err != nil {
		return
	}

	if err = yaml.Unmarshal(data, &config); err != nil {
		return
	}

	return
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
