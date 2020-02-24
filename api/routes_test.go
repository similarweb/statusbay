package api_test

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"statusbay/api"
	"statusbay/api/alerts"
	"statusbay/api/httpresponse"
	"statusbay/api/metrics"
	"statusbay/api/testutil"
	"sync"
	"testing"
)

type testServer struct {
	api *api.Server
}

func MockServer(t *testing.T, storageMockFile string, metrics map[string]metrics.MetricManagerDescriber, alertsClient map[string]alerts.AlertsManagerDescriber) testServer {

	version := testutil.NewMockVersion()
	storage := testutil.NewMockStorage()
	return testServer{
		api: api.NewServer(storage, "8080", "./testutil/events.yaml", metrics, alertsClient, version),
	}
}

func TestApplicationMetricsEndpoint(t *testing.T) {
	var wg *sync.WaitGroup
	ctx := context.Background()

	metrics := make(map[string]metrics.MetricManagerDescriber)
	metrics["dummy"] = testutil.NewMockMetrics()
	ms := MockServer(t, "", metrics, nil)
	ms.api.BindEndpoints()
	ms.api.Serve(ctx, wg)

	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/v1/application/metric?provider=dummy&query=foo.2xx&from=1&to=1", nil)
	if err != nil {
		t.Fatalf("http request returned with error")
	}

	ms.api.Router().ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("handler returned unexpected status code: got %v want %v", rr.Code, http.StatusOK)
	}

	if err != nil {
		t.Fatal(err)
	}

	response := []httpresponse.MetricsQuery{}
	body, err := ioutil.ReadAll(rr.Body)
	json.Unmarshal(body, &response)

	if err != nil {
		t.Fatal(err)
	}

	if len(response) != 1 {
		t.Fatalf("unexpected length for metrics endpoint response, got %d expected %d", len(response), 1)
	}

}

func TestApplicationMetricsEndpointWithInvalidQueryParameters(t *testing.T) {
	var wg *sync.WaitGroup
	ctx := context.Background()

	metrics := make(map[string]metrics.MetricManagerDescriber)
	metrics["dummy"] = testutil.NewMockMetrics()
	ms := MockServer(t, "", metrics, nil)
	ms.api.BindEndpoints()
	ms.api.Serve(ctx, wg)

	testCases := []struct {
		endpoint                string
		expectedStatusCode      int
		expectedValidationCount int
	}{
		{"/api/v1/application/metric", http.StatusBadRequest, 4},
		{"/api/v1/application/metric?query=2xx", http.StatusBadRequest, 3},
		{"/api/v1/application/metric?query=2xx&from=1", http.StatusBadRequest, 2},
		{"/api/v1/application/metric?query=2xx&from=2&to=1", http.StatusBadRequest, 2},
		{"/api/v1/application/metric?query=2xx&from=a&to=b", http.StatusBadRequest, 3},
		{"/api/v1/application/metric?query=2xx&from=1&to=123", http.StatusBadRequest, 1},
		{"/api/v1/application/metric?query=2xx&from=1&to=123&provider=dummy1", http.StatusBadRequest, 1},
		{"/api/v1/application/metric?query=2xx&from=1&to=123&provider=dummy", http.StatusOK, 0},
	}

	for _, test := range testCases {
		t.Run(test.endpoint, func(t *testing.T) {
			rr := httptest.NewRecorder()
			req, err := http.NewRequest("GET", test.endpoint, nil)
			if err != nil {
				t.Fatalf("http request returned with error")
			}

			ms.api.Router().ServeHTTP(rr, req)
			if rr.Code != test.expectedStatusCode {
				t.Fatalf("handler returned unexpected status code: got %d want %d", rr.Code, test.expectedStatusCode)
			}

			if err != nil {
				t.Fatal(err)
			}

			response := map[string]map[string][]string{}
			body, err := ioutil.ReadAll(rr.Body)
			json.Unmarshal(body, &response)

			if len(response["Validation"]) != test.expectedValidationCount {
				t.Fatalf("unexpected validation message length response, got %d, expected %d", len(response["Validation"]), test.expectedValidationCount)
			}
		})
	}

}

func TestApplicationAlertsEndpoint(t *testing.T) {
	var wg *sync.WaitGroup
	ctx := context.Background()

	alerts := testutil.NewMultipleMockAlerts()
	ms := MockServer(t, "", nil, alerts)
	ms.api.BindEndpoints()
	ms.api.Serve(ctx, wg)

	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/v1/application/alerts?tags=foo,foo2&from=1&to=1&provider=foo", nil)
	if err != nil {
		t.Fatalf("http request returned with error")
	}

	ms.api.Router().ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("handler returned unexpected status code: got %d want %d", rr.Code, http.StatusOK)
	}

	if err != nil {
		t.Fatal(err)
	}

	response := []httpresponse.CheckResponse{}
	body, err := ioutil.ReadAll(rr.Body)
	json.Unmarshal(body, &response)

	if len(response) != 1 {
		t.Fatalf("unexpected length for metrics endpoint response, got %d expected %d", len(response), 1)
	}

}

func TestApplicationAlertsEndpointWithInvalidQueryParameters(t *testing.T) {
	var wg *sync.WaitGroup
	ctx := context.Background()

	alerts := testutil.NewMultipleMockAlerts()
	ms := MockServer(t, "", nil, alerts)
	ms.api.BindEndpoints()
	ms.api.Serve(ctx, wg)

	testCases := []struct {
		endpoint                string
		expectedStatusCode      int
		expectedValidationCount int
	}{
		{"/api/v1/application/alerts", http.StatusBadRequest, 4},
		{"/api/v1/application/alerts?tags=foo", http.StatusBadRequest, 3},
		{"/api/v1/application/alerts?tags=foo&from=1", http.StatusBadRequest, 2},
		{"/api/v1/application/alerts?tags=foo&from=1&to=123", http.StatusBadRequest, 1},
		{"/api/v1/application/alerts?tags=foo&from=1&to=123&provider=foo", http.StatusOK, 0},
	}

	for _, test := range testCases {
		t.Run(test.endpoint, func(t *testing.T) {
			rr := httptest.NewRecorder()
			req, err := http.NewRequest("GET", test.endpoint, nil)
			if err != nil {
				t.Fatalf("http request returned with error")
			}

			ms.api.Router().ServeHTTP(rr, req)
			if rr.Code != test.expectedStatusCode {
				t.Fatalf("handler returned unexpected status code: got %d want %d", rr.Code, test.expectedStatusCode)
			}

			if err != nil {
				t.Fatal(err)
			}

			response := map[string]map[string][]string{}
			body, err := ioutil.ReadAll(rr.Body)
			err = json.Unmarshal(body, &response)

			if len(response["Validation"]) != test.expectedValidationCount {
				t.Fatalf("unexpected validation message length response, got %d, expected %d", len(response["Validation"]), test.expectedValidationCount)
			}
		})
	}

}
