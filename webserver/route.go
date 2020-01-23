package webserver

import (
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"statusbay/webserver/httpparameters"
	"statusbay/webserver/httpresponse"
)

// MetricHandler return metrics providers data
func (server *Server) MetricHandler(resp http.ResponseWriter, req *http.Request) {
	errs := url.Values{}

	// Parse query parameters
	provider := httpparameters.QueryParamWithDefault(req, "provider", "")
	query := httpparameters.QueryParamWithDefault(req, "query", "")
	from := httpparameters.QueryParamWithDefault(req, "from", "")
	to := httpparameters.QueryParamWithDefault(req, "to", "")
	fromInt, _ := strconv.ParseInt(from, 10, 64)
	toInt, _ := strconv.ParseInt(to, 10, 64)

	// Validate query parameters
	if query == "" {
		errs.Add("query", "The query field is required")
	}

	if provider == "" {
		errs.Add("provider", "The provider field is required")
	}

	if fromInt == 0 {
		errs.Add("from", "The from field is required")
	}

	if toInt == 0 {
		errs.Add("to", "The to field is required")
	}

	if fromInt != 0 && toInt != 0 && fromInt > toInt {
		errs.Add("range", "To field should be bigger the from")
	}

	if len(errs) > 0 {
		httpresponse.JSONErrorParameters(resp, http.StatusBadRequest, errs)
		return
	}

	if server.metricClientProviders[provider] == nil {
		httpresponse.JSONError(resp, http.StatusBadRequest, errors.New("Client metric not enabled"))
		return
	}

	metrics, err := server.metricClientProviders[provider].GetMetric(query, time.Unix(fromInt, 0), time.Unix(toInt, 0))
	if err != nil {
		httpresponse.JSONError(resp, http.StatusInternalServerError, err)
		return
	}
	httpresponse.JSONWrite(resp, http.StatusOK, metrics)
}

func (server *Server) AlertsHandler(resp http.ResponseWriter, req *http.Request) {

	errs := url.Values{}
	tags := httpparameters.QueryParamWithDefault(req, "tags", "")
	provider := httpparameters.QueryParamWithDefault(req, "provider", "")
	from := httpparameters.QueryParamWithDefault(req, "from", "")
	to := httpparameters.QueryParamWithDefault(req, "to", "")
	fromInt, _ := strconv.ParseInt(from, 10, 64)
	toInt, _ := strconv.ParseInt(to, 10, 64)

	// Validate query parameters
	if tags == "" {
		errs.Add("query", "The title field is required")
	}

	if provider == "" {
		errs.Add("provider", "The title field is required")
	}

	if fromInt == 0 {
		errs.Add("from", "The title field is required")
	}

	if toInt == 0 {
		errs.Add("to", "The title field is required")
	}

	if len(errs) > 0 {
		httpresponse.JSONErrorParameters(resp, http.StatusBadRequest, errs)
		return
	}

	clientProvider, found := server.alertClientProviders[provider]

	if !found {
		httpresponse.JSONError(resp, http.StatusBadRequest, errors.New("Provider not supported"))
		return
	}

	alerts, err := clientProvider.GetAlertByTags(tags, time.Unix(fromInt, 0), time.Unix(toInt, 0))

	if err != nil {
		httpresponse.JSONError(resp, http.StatusInternalServerError, err)
		return
	}

	httpresponse.JSONWrite(resp, http.StatusOK, alerts)

}

//NotFoundRoute return when route not found
func (server *Server) NotFoundRoute(resp http.ResponseWriter, req *http.Request) {
	httpresponse.JSONError(resp, http.StatusNotFound, errors.New("Path not found"))
}

//HealthCheckHandler return ok if server is up
func (server *Server) HealthCheckHandler(resp http.ResponseWriter, req *http.Request) {
	httpresponse.JSONWrite(resp, http.StatusOK, httpresponse.HealthResponse{Status: true})
}
