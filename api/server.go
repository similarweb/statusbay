package api

import (
	"context"
	"fmt"
	"net/http"
	"statusbay/config"
	"sync"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"statusbay/api/alerts"
	"statusbay/api/kubernetes"
	"statusbay/api/metrics"
	"statusbay/version"
)

const (
	// DrainTimeout is how long to wait until the server is drained before closing it
	DrainTimeout = time.Second * 30
)

// Server is the API server struct
type Server struct {
	router     *mux.Router
	httpserver *http.Server

	kubernetesStorage     kubernetes.Storage
	kubernetesMarkEvents  config.KubernetesMarksEvents
	metricClientProviders map[string]metrics.MetricManagerDescriber
	alertClientProviders  map[string]alerts.AlertsManagerDescriber
	version               version.VersionDescriptor
	absoluteLogsPodPath   string
}

// NewServer returns a new Server
func NewServer(kubernetesStorage kubernetes.Storage, port string, kubernetesMarkEvents config.KubernetesMarksEvents, metricClientProviders map[string]metrics.MetricManagerDescriber, alertClientProviders map[string]alerts.AlertsManagerDescriber, version version.VersionDescriptor, absoluteLogsPodPath string) *Server {

	router := mux.NewRouter()
	corsObj := handlers.AllowedOrigins([]string{"*"})
	return &Server{
		router:                router,
		kubernetesStorage:     kubernetesStorage,
		kubernetesMarkEvents:  kubernetesMarkEvents,
		metricClientProviders: metricClientProviders,
		alertClientProviders:  alertClientProviders,
		version:               version,
		absoluteLogsPodPath:   absoluteLogsPodPath,
		httpserver: &http.Server{
			Handler: handlers.CORS(corsObj)(router),
			Addr:    fmt.Sprintf("0.0.0.0:%s", port),
		},
	}
}

// Serve starts the HTTP server and listens until StopFunc is called
func (server *Server) Serve(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	server.BindEndpoints()
	log.WithField("bind_address", server.httpserver.Addr).Info("starting StatusBay server")
	go func() {
		<-ctx.Done()
		err := server.httpserver.Shutdown(ctx)
		if err != nil {
			log.WithError(err).Error("failed to shutdown manager HTTP server")
		}
		log.Warn("HTTP server has been shut down")
		wg.Done()
	}()
	go func() {
		server.httpserver.ListenAndServe()
	}()

}

// BindEndpoints sets up the router to handle API endpoints
func (server *Server) BindEndpoints() {

	// KUBERNETES ROUTES
	kubernetes.NewKubernetesRoutes(server.kubernetesStorage, server.router, server.kubernetesMarkEvents, server.absoluteLogsPodPath)

	// Genetic routes
	server.router.HandleFunc("/api/v1/health", server.HealthCheckHandler).Methods("GET")
	server.router.HandleFunc("/api/v1/version", server.VersionHandler).Methods("GET")
	server.router.HandleFunc("/api/v1/application/metric", server.MetricHandler).Methods("GET")
	server.router.HandleFunc("/api/v1/application/alerts", server.AlertsHandler).Methods("GET")
	server.router.NotFoundHandler = http.HandlerFunc(server.NotFoundRoute)

}

// Router returns the Gorilla Mux HTTP router defined for this server
func (server *Server) Router() *mux.Router {
	return server.router
}
