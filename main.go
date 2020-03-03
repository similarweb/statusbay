package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"statusbay/api"
	"statusbay/api/alerts"
	apiKubernetes "statusbay/api/kubernetes"
	"statusbay/api/metrics"
	"statusbay/cache"
	"statusbay/config"
	"statusbay/serverutil"
	"statusbay/state"
	"statusbay/version"
	"statusbay/visibility"
	kuberneteswatcher "statusbay/watcher/kubernetes"
	"statusbay/watcher/kubernetes/client"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	// DefaultGracefulShutDown is the default graceful shot down the server
	DefaultGracefulShutDown = time.Second * 10

	// DefaultConfigPath is the default configuration file path
	DefaultConfigPath = "/etc/statusbay/config.yaml"

	// DefaultEventsPath is the default events configuration file path
	DefaultEventsPath = "/etc/statusbay/events.yaml"

	//ModeAPI will upload a server for client Website UI
	ModeAPI = "api"

	//KubernetesWatcher start watch on kubernetes deployments
	KubernetesWatcher = "kubernetes"
)

func main() {

	var configPath, eventsPath, mode string
	// parsing flags
	flag.StringVar(&mode, "mode", "", fmt.Sprintf("Server mode to start. Must be either \"%s\" or \"%s\".", ModeAPI, KubernetesWatcher))
	flag.StringVar(&configPath, "config", DefaultConfigPath, "Path to configuration file")
	flag.StringVar(&eventsPath, "events", DefaultEventsPath, "Path to events configuration file")

	// Only for kubernetes
	var kubeconfig, apiserverHost string
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to kubeconfig file with authorization and master location information.")
	flag.StringVar(&apiserverHost, "apiserverhost", "", "Path to kubeconfig file with authorization and master location information.")
	flag.Parse()

	ctx, cancelFn := context.WithCancel(context.Background())
	var runner *serverutil.Runner

	switch mode {
	case ModeAPI:
		runner = startAPIServer(ctx, configPath, eventsPath)
	case KubernetesWatcher:
		runner = startKubernetesWatcher(ctx, configPath, kubeconfig, apiserverHost)
	default:
		flag.Usage()
		os.Exit(1)
	}

	//register signal handler
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop // block until we are requested to stop
	runner.StopFunc(cancelFn)

}

func startKubernetesWatcher(ctx context.Context, configPath, kubeconfig, apiserverHost string) *serverutil.Runner {

	version.NewVersion(ctx, "wacher_kubernetes", 12*time.Hour)

	watcherConfig, err := config.LoadKubernetesConfig(configPath)
	if err != nil {
		log.WithError(err).Panic("could not load Kubernetes configuration file")
		os.Exit(1)
	}

	err = config.InitMetricAggregator(watcherConfig.Telemetry)
	if err != nil {
		log.WithError(err).Panic("failed to initialize telemetry")
		os.Exit(1)
	}

	//Setup logging
	visibility.SetupLogging(watcherConfig.Log.Level, watcherConfig.Log.GelfAddress, "wacher_kubernetes")

	// Init kubernetes client
	kubernetesClientManager, err := client.NewClientManager(kubeconfig, apiserverHost)
	if err != nil {
		log.WithError(err).Panic("failed to initialize Kubernetes client")
		os.Exit(1)
	}
	kubernetesClientset := kubernetesClientManager.GetInsecureClient()

	// Init mysql storage
	mysqlManager := state.NewMysqlClient(watcherConfig.MySQL)
	mysqlManager.Migration()
	mysql := kuberneteswatcher.NewMysql(mysqlManager)

	notifiers, err := watcherConfig.BuildNotifiers()
	if err != nil {
		log.WithError(err).Panic("failed to initialize notifiers")
		os.Exit(1)
	}

	// Init Reporter
	reporter := kuberneteswatcher.NewReporter(notifiers)

	//Registry manager
	registryManager := kuberneteswatcher.NewRegistryManager(watcherConfig.Applies.SaveInterval, watcherConfig.Applies.CheckFinishDelay, watcherConfig.Applies.CollectDataAfterApplyFinish, mysql, reporter, watcherConfig.ClusterName)
	runningApplies := registryManager.LoadRunningApplies()
	//Event manager
	eventManager := kuberneteswatcher.NewEventsManager(kubernetesClientset)

	//Service manager
	serviceManager := kuberneteswatcher.NewServiceManager(kubernetesClientset, eventManager)

	//Pvc manager
	pvcManager := kuberneteswatcher.NewPvcManager(kubernetesClientset, eventManager)

	//Pods manager
	podsManager := kuberneteswatcher.NewPodsManager(kubernetesClientset, eventManager, pvcManager)

	//Replicaset manager
	replicasetManager := kuberneteswatcher.NewReplicasetManager(kubernetesClientset, eventManager, podsManager)

	//Deployment manager
	deploymentManager := kuberneteswatcher.NewDeploymentManager(kubernetesClientset, eventManager, registryManager, replicasetManager, serviceManager, runningApplies, watcherConfig.Applies.MaxApplyTime)

	// ControllerRevision Manager
	controllerRevisionManager := kuberneteswatcher.NewControllerRevisionManager(kubernetesClientset, podsManager)

	// Daemonset manager
	daemonsetManager := kuberneteswatcher.NewDaemonsetManager(kubernetesClientset, eventManager, registryManager, serviceManager, controllerRevisionManager, runningApplies, watcherConfig.Applies.MaxApplyTime)

	//Statefulset manager
	statefulsetManager := kuberneteswatcher.NewStatefulsetManager(kubernetesClientset, eventManager, registryManager, serviceManager, controllerRevisionManager, runningApplies, watcherConfig.Applies.MaxApplyTime)

	servers := []serverutil.Server{
		eventManager, podsManager, pvcManager, deploymentManager, daemonsetManager, statefulsetManager, replicasetManager, registryManager, serviceManager, reporter,
	}

	// Run a list of backround process for the server
	return serverutil.RunAll(ctx, servers)
}

func startAPIServer(ctx context.Context, configPath string, eventsPath string) *serverutil.Runner {

	version := version.NewVersion(ctx, "webserver", 12*time.Hour)

	apiConfig, err := config.LoadConfigAPI(configPath)
	if err != nil {
		log.WithError(err).Panic("could not load API configuration file")
		os.Exit(1)
	}

	eventsConfig, err := config.LoadEvents(eventsPath)
	if err != nil {
		log.WithError(err).Panic("could not load events configuration file")
		os.Exit(1)
	}

	err = config.InitMetricAggregator(apiConfig.Telemetry)
	if err != nil {
		log.WithError(err).Panic("failed to initialize telemetry")
		os.Exit(1)
	}

	cacheManager := cache.NewRedisClient(apiConfig.Redis)

	//Setup logging
	visibility.SetupLogging(apiConfig.Log.Level, apiConfig.Log.GelfAddress, "api")

	mysqlManager := state.NewMysqlClient(apiConfig.MySQL)
	mysqlManager.Migration()

	// TODO:: should be more generic solution, we can start with this solution when we use only one orchestration
	kubernetesStorage := apiKubernetes.NewMysql(mysqlManager)

	metricsProviders := metrics.Load(apiConfig.MetricsProvider, cacheManager)

	alertsProviders := alerts.Load(apiConfig.AlertProvider)

	//Start the server
	server := api.NewServer(kubernetesStorage, "8080", eventsConfig, metricsProviders, alertsProviders, version)

	servers := []serverutil.Server{
		server,
	}

	for _, metric := range metricsProviders {
		servers = append(servers, metric)
	}
	//run lis of backround process for the server
	return serverutil.RunAll(ctx, servers)

}
