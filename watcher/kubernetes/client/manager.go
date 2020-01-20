package client

import (
	"errors"
	log "github.com/sirupsen/logrus"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

// ClientManagerDescriber descrive client managet interface
type ClientManagerDescriber interface {
	GetInsecureClient() kubernetes.Interface
}

const (
	// DefaultQPS indicates the maximum QPS to the master from this client.
	DefaultQPS = 1e6

	//DefaultBurst indicates maximum burst for throttle.
	DefaultBurst = 1e6

	// DefaultContentType indicates the content type specifies the wire format used to communicate with the server.
	DefaultContentType = "application/json"

	// DefaultUserAgent specifies the caller of this request.
	DefaultUserAgent = "statusbay"
)

// Version of this binary
var Version = "UNKNOWN"

// clientManager create new kubernetes client manager
type clientManager struct {
	kubeConfigPath  string
	apiserverHost   string
	inClusterConfig *rest.Config
	insecureConfig  *rest.Config

	insecureClient kubernetes.Interface
}

// GetInsecureClient return the kubernetes client instance
func (cm *clientManager) GetInsecureClient() kubernetes.Interface {
	return cm.insecureClient
}

// initInClusterConfig create kubernetes rest config instance
func (cm *clientManager) initInClusterConfig() error {
	if len(cm.apiserverHost) > 0 || len(cm.kubeConfigPath) > 0 {
		return nil
	}

	cfg, err := rest.InClusterConfig()
	if err != nil {
		return err
	}

	cm.inClusterConfig = cfg
	return nil
}

// initInsecureConfig create kubernetes client config
func (cm *clientManager) initInsecureConfig() {
	cfg, err := cm.buildConfigFromFlags(cm.apiserverHost, cm.kubeConfigPath)
	if err != nil {
		panic(err)
	}

	cm.initConfig(cfg)
	cm.insecureConfig = cfg
}

// initConfig overide kubernetes rest config
func (cm *clientManager) initConfig(cfg *rest.Config) {
	cfg.QPS = DefaultQPS
	cfg.Burst = DefaultBurst
	cfg.ContentType = DefaultContentType
	cfg.UserAgent = DefaultUserAgent + "/" + Version
}

// buildConfigFromFlags return the rest kubernetes config.
func (cm *clientManager) buildConfigFromFlags(apiserverHost, kubeConfigPath string) (*rest.Config, error) {
	if len(kubeConfigPath) > 0 || len(apiserverHost) > 0 {
		return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
			&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeConfigPath},
			&clientcmd.ConfigOverrides{ClusterInfo: api.Cluster{Server: apiserverHost}}).ClientConfig()
	}

	if cm.inClusterConfig != nil {
		return cm.inClusterConfig, nil
	}

	return nil, errors.New("could not create client config")
}

// initInsecureClients creates Kubernetes client client.
func (cm *clientManager) initInsecureClients() error {
	cm.initInsecureConfig()
	k8sClient, err := kubernetes.NewForConfig(cm.insecureConfig)
	if err != nil {
		return err
	}

	cm.insecureClient = k8sClient
	return nil
}

// init kubernetes client
func (cm *clientManager) init() error {
	err := cm.initInClusterConfig()
	if err != nil {
		return err
	}

	err = cm.initInsecureClients()
	return err
}

// NewClientManager creates client manager with given kubeConfigPath and apiserverHost parameters.
func NewClientManager(kubeConfigPath, apiserverHost string) (ClientManagerDescriber, error) {

	log.WithFields(log.Fields{
		"config_path": kubeConfigPath,
		"server_host": apiserverHost,
	}).Info("Setting up kubernetes client manager")

	result := &clientManager{
		kubeConfigPath: kubeConfigPath,
		apiserverHost:  apiserverHost,
	}

	err := result.init()
	return result, err
}
