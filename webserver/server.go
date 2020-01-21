package webserver

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"statusbay/serverutil"
	"statusbay/webserver/alerts"
	"statusbay/webserver/kubernetes"
	"statusbay/webserver/metrics"
)

const (
	// DrainTimeout is how long to wait until the server is drained before closing it
	DrainTimeout = time.Second * 30
)

// Server is the API server struct
type Server struct {
	router     *mux.Router
	httpserver *http.Server

	kubernetesStorage        kubernetes.Storage
	kubernetesMarkEventsPath string
	metricClientProviders    map[string]metrics.MetricManagerDescriber
	alertClientProviders     map[string]alerts.AlertsManagerDescriber
}

// NewServer returns a new Server
func NewServer(kubernetesStorage kubernetes.Storage, port string, kubernetesMarkEventsPath string, metricClientProviders map[string]metrics.MetricManagerDescriber, alertClientProviders map[string]alerts.AlertsManagerDescriber) *Server {

	router := mux.NewRouter()
	corsObj := handlers.AllowedOrigins([]string{"*"})
	return &Server{
		router:                   router,
		kubernetesStorage:        kubernetesStorage,
		kubernetesMarkEventsPath: kubernetesMarkEventsPath,
		metricClientProviders:    metricClientProviders,
		alertClientProviders:     alertClientProviders,
		httpserver: &http.Server{
			Handler: handlers.CORS(corsObj)(router),
			Addr:    fmt.Sprintf("0.0.0.0:%s", port),
		},
	}
}

// Serve starts the HTTP server and listens until StopFunc is called
func (server *Server) Serve() serverutil.StopFunc {
	ctx, cancelFn := context.WithCancel(context.Background())
	server.BindEndpoints()
	log.WithField("bind_address", server.httpserver.Addr).Info("Starting statusbay server")
	stopped := make(chan bool)
	go func() {
		<-ctx.Done()
		serverCtx, serverCancelFn := context.WithTimeout(context.Background(), DrainTimeout)
		err := server.httpserver.Shutdown(serverCtx)
		if err != nil {
			log.WithError(err).Error("error occured while shutting down manager HTTP server")
		}
		serverCancelFn()
		stopped <- true
	}()
	go func() {
		server.httpserver.ListenAndServe()
	}()
	return func() {
		cancelFn()
		<-stopped
		log.Warn("HTTP server has been drained and shut down")
	}
}

// BindEndpoints sets up the router to handle API endpoints
func (server *Server) BindEndpoints() {

	// KUBERNETES ROUTES
	kubernetes.NewKubernetesRoutes(server.kubernetesStorage, server.router, server.kubernetesMarkEventsPath)

	// Genetic routes
	server.router.HandleFunc("/api/v1/health", server.HealthCheckHandler).Methods("GET")
	server.router.HandleFunc("/api/v1/application/metric", server.MetricHandler).Methods("GET")
	server.router.HandleFunc("/api/v1/application/alerts", server.AlertsHandler).Methods("GET")
	server.router.NotFoundHandler = http.HandlerFunc(server.NotFoundRoute)

}

// Router returns the Gorilla Mux HTTP router defined for this server
func (server *Server) Router() *mux.Router {
	return server.router
}
