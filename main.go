package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"statusbay/config"
	"statusbay/serverutil"
	"statusbay/state"
	"statusbay/visibility"

	"statusbay/api"
	"statusbay/api/alerts"
	apiKubernetes "statusbay/api/kubernetes"
	"statusbay/api/metrics"
	kuberneteswatcher "statusbay/watcher/kubernetes"
	"statusbay/watcher/kubernetes/client"
	"time"

	"github.com/nlopes/slack"

	log "github.com/sirupsen/logrus"
)

const (
	// DefaultGracefulShutDown is the default gracefull shot down the server
	DefaultGracefulShutDown = time.Second * 10

	// DefaultConfigPath is the default configuration file path
	DefaultConfigPath = "/etc/statusbay/config.yaml"

	//ModeAPI will upload a server for client Website UI
	ModeAPI = "api"

	//KubernetesWatcher start watch on kubernetes deployments
	KubernetesWatcher = "kubernetes"
)

func main() {

	var configPath, mode string
	// parsing flags
	flag.StringVar(&mode, "mode", "", fmt.Sprintf("Server mode to start. Must be either \"%s\" or \"%s\".", ModeAPI, KubernetesWatcher))
	flag.StringVar(&configPath, "config", DefaultConfigPath, "Path to configuration file")

	// Only for kubernetes
	var kubeconfig, apiserverHost string
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to kubeconfig file with authorization and master location information.")
	flag.StringVar(&apiserverHost, "apiserverhost", "", "Path to kubeconfig file with authorization and master location information.")
	flag.Parse()

	var stopper serverutil.StopFunc
	switch mode {
	case ModeAPI:
		stopper = startAPIServer(configPath, "./events.yaml")
	case KubernetesWatcher:
		stopper = startKubernetesWatcher(configPath, kubeconfig, apiserverHost)
	default:
		flag.Usage()
		os.Exit(1)
	}

	//register signal handler
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop // block until we are requested to stop
	stopper()

}

func startKubernetesWatcher(configPath, kubeconfig, apiserverHost string) serverutil.StopFunc {

	watcherConfig, err := config.LoadKubernetesConfig(configPath)
	if err != nil {
		log.WithError(err).Panic("could not load Kubernetes configuration file")
		os.Exit(1)
	}

	// Init kubernetes client
	kubernetesClientManager, err := client.NewClientManager(kubeconfig, apiserverHost)
	if err != nil {
		log.WithError(err).Panic("could not init kubernetes client")
		os.Exit(1)
	}
	kubernetesClientset := kubernetesClientManager.GetInsecureClient()

	// Init mysql storage
	mysqlManager := state.NewMysqlClient(watcherConfig.MySQL)
	mysqlManager.Migration()
	mysql := kuberneteswatcher.NewMysql(mysqlManager)

	//Slack
	api := slack.New(watcherConfig.Slack.Token)
	slack := state.NewSlack(api)
	reporter := kuberneteswatcher.NewReporter(slack, watcherConfig.Slack.DefaultChannels, watcherConfig.UI.BaseURL)

	//Replicaset manager
	registryManager := kuberneteswatcher.NewRegistryManager(watcherConfig.KubernetesConfig.Deployment.SaveInterval, watcherConfig.KubernetesConfig.Deployment.SaveDeploymentHistoryDuration, watcherConfig.KubernetesConfig.Deployment.CheckFinishDelay, watcherConfig.KubernetesConfig.Deployment.CollectDataAfterDeploymentFinish, mysql, reporter)

	//Event manager
	eventManager := kuberneteswatcher.NewEventsManager(kubernetesClientset)

	//Service manager
	serviceManager := kuberneteswatcher.NewServiceManager(kubernetesClientset)

	//Pods manager
	podsManager := kuberneteswatcher.NewPodsManager(kubernetesClientset, eventManager)

	//Replicaset manager
	replicasetManager := kuberneteswatcher.NewReplicasetManager(kubernetesClientset, eventManager, podsManager)

	//Deployment manager
	deploymentManager := kuberneteswatcher.NewDeploymentManager(kubernetesClientset, eventManager, registryManager, replicasetManager, serviceManager, watcherConfig.KubernetesConfig.Deployment.MaxDeploymentTime)

	// ControllerRevision Manager
	controllerRevisionManager := kuberneteswatcher.NewControllerReisionManager(kubernetesClientset, podsManager)
	// Daemonset manager
	daemonsetManager := kuberneteswatcher.NewDaemonsetManager(kubernetesClientset, eventManager, registryManager, serviceManager, podsManager, controllerRevisionManager, watcherConfig.KubernetesConfig.Deployment.MaxDeploymentTime)

	//run lis of backround proccess for the server
	return serverutil.RunAll(eventManager, podsManager, deploymentManager, daemonsetManager, replicasetManager, registryManager, serviceManager, reporter).StopFunc

}

func startAPIServer(configPath, eventConfigPath string) serverutil.StopFunc {

	apiConfig, err := config.LoadConfigAPI(configPath)
	if err != nil {
		log.WithError(err).Panic("could not load configuration file")
		os.Exit(1)
	}

	//Setup logging
	visibility.SetupLogging(apiConfig.LogLevel, "api")

	mysqlManager := state.NewMysqlClient(apiConfig.MySQL)
	mysqlManager.Migration()

	// TODO:: should be more generic solution, we can start with this solution when we use only one orchestration
	kubernetesStorage := apiKubernetes.NewMysql(mysqlManager)

	var metricClient metrics.MetricManagerDescriber

	metricsProviders := metrics.Load(apiConfig.MetricsProvider)

	alertsProviders := alerts.Load(apiConfig.AlertProvider)

	//Start the server
	server := api.NewServer(kubernetesStorage, "8080", eventConfigPath, metricsProviders, alertsProviders)

	//run lis of backround proccess for the server
	return serverutil.RunAll(server, metricClient).StopFunc

}
