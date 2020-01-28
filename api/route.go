package api

import (
	"errors"
	"net/http"
	"net/url"
	"statusbay/api/alerts"
	"strconv"
	"time"

	"statusbay/api/httpparameters"
	"statusbay/api/httpresponse"
	"statusbay/api/metrics"
)

// MetricHandler returns query metric data points for a specific time window
func (server *Server) MetricHandler(resp http.ResponseWriter, req *http.Request) {
	var metricsProvider metrics.MetricManagerDescriber
	var found bool
	errs := url.Values{}

	// Parse query parameters
	provider := httpparameters.QueryParamWithDefault(req, "provider", "")
	query := httpparameters.QueryParamWithDefault(req, "query", "")
	from := httpparameters.QueryParamWithDefault(req, "from", "")
	to := httpparameters.QueryParamWithDefault(req, "to", "")

	// Convert `from` and `to` query params to integers
	fromInt, _ := strconv.ParseInt(from, 10, 64)
	toInt, _ := strconv.ParseInt(to, 10, 64)

	// Query parameters validation
	if provider == "" {
		errs.Add("provider", "the provider field is mandatory")
	} else {
		metricsProvider, found = server.metricClientProviders[provider]
		if !found {
			errs.Add("provider-existance", "provider not configured")
		}
	}

	if query == "" {
		errs.Add("query", "the `query` field is mandatory")
	}

	if fromInt == 0 {
		errs.Add("from", "the `from` field is mandatory")
	}

	if toInt == 0 {
		errs.Add("to", "the `to` field is mandatory")
	}

	if fromInt != 0 && toInt != 0 && fromInt > toInt {
		errs.Add("range", "`to` field has to be bigger than `from` field")
	}

	if len(errs) > 0 {
		httpresponse.JSONErrorParameters(resp, http.StatusBadRequest, errs)
		return
	}

	metrics, err := metricsProvider.GetMetric(query, time.Unix(fromInt, 0), time.Unix(toInt, 0))
	if err != nil {
		httpresponse.JSONError(resp, http.StatusInternalServerError, err)
		return
	}

	httpresponse.JSONWrite(resp, http.StatusOK, metrics)
}

// AlertsHandler returns triggered data points for alerts filtered by tags
func (server *Server) AlertsHandler(resp http.ResponseWriter, req *http.Request) {
	var alertsProvider alerts.AlertsManagerDescriber
	var found bool
	errs := url.Values{}

	// Parse query parameters
	provider := httpparameters.QueryParamWithDefault(req, "provider", "")
	tags := httpparameters.QueryParamWithDefault(req, "tags", "")
	from := httpparameters.QueryParamWithDefault(req, "from", "")
	to := httpparameters.QueryParamWithDefault(req, "to", "")

	// Convert `from` and `to` query params to integers
	fromInt, _ := strconv.ParseInt(from, 10, 64)
	toInt, _ := strconv.ParseInt(to, 10, 64)

	// Query parameters validation
	if provider == "" {
		errs.Add("provider", "the `provider` field is mandatory")
	} else {
		alertsProvider, found = server.alertClientProviders[provider]
		if !found {
			errs.Add("provider-existance", "provider not configured")
		}
	}

	if tags == "" {
		errs.Add("tags", "the `tags` field is mandatory")
	}

	if fromInt == 0 {
		errs.Add("from", "the `from` field is mandatory")
	}

	if toInt == 0 {
		errs.Add("to", "the `to` field is mandatory")
	}

	if fromInt != 0 && toInt != 0 && fromInt > toInt {
		errs.Add("range", "`to` field has to be bigger than `from` field")
	}

	if len(errs) > 0 {
		httpresponse.JSONErrorParameters(resp, http.StatusBadRequest, errs)
		return
	}

	alerts, err := alertsProvider.GetAlertByTags(tags, time.Unix(fromInt, 0), time.Unix(toInt, 0))
	if err != nil {
		httpresponse.JSONError(resp, http.StatusInternalServerError, err)
		return
	}

	httpresponse.JSONWrite(resp, http.StatusOK, alerts)
}

// NotFoundRoute a 404 handler
func (server *Server) NotFoundRoute(resp http.ResponseWriter, req *http.Request) {
	httpresponse.JSONError(resp, http.StatusNotFound, errors.New("path not found"))
}

// HealthCheckHandler application health check
func (server *Server) HealthCheckHandler(resp http.ResponseWriter, req *http.Request) {
	httpresponse.JSONWrite(resp, http.StatusOK, httpresponse.HealthResponse{Status: true})
}
