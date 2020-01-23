package api_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"statusbay/api"
	"statusbay/api/alerts"
	"statusbay/api/httpresponse"
	"statusbay/api/metrics"
	"statusbay/api/testutil"
	"testing"
)

type testServer struct {
	api *api.Server
}

func MockServer(t *testing.T, storageMockFile string, metrics metrics.MetricManagerDescriber, alertsClient map[string]alerts.AlertsManagerDescriber) testServer {

	storage := testutil.NewMockStorage()
	return testServer{
		api: api.NewServer(storage, "8080", "./kubernetes/testutil/events.yaml", metrics, alertsClient),
	}
}

func TestApplicationMetricsEndpoint(t *testing.T) {
	metrics := testutil.NewMockMetrics()
	ms := MockServer(t, "", metrics, nil)
	ms.api.BindEndpoints()
	ms.api.Serve()

	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/v1/application/metric?query=foo.2xx&from=1&to=1", nil)
	if err != nil {
		t.Fatalf("Http request returned with error")
	}

	ms.api.Router().ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
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
		t.Fatalf("unexpected metrics endpont response, got %d expected %d", len(response), 1)
	}

}

func TestApplicationMetricsEndpointWithInvalidQueryParameters(t *testing.T) {
	metrics := testutil.NewMockMetrics()
	ms := MockServer(t, "", metrics, nil)
	ms.api.BindEndpoints()
	ms.api.Serve()

	testCases := []struct {
		endpoint                string
		expectedStatusCode      int
		expectedValidationCount int
	}{
		{"/api/v1/application/metric", http.StatusBadRequest, 3},
		{"/api/v1/application/metric?query=2xx", http.StatusBadRequest, 2},
		{"/api/v1/application/metric?query=2xx&from=1", http.StatusBadRequest, 1},
		{"/api/v1/application/metric?query=2xx&from=2&to=1", http.StatusBadRequest, 1},
		{"/api/v1/application/metric?query=2xx&from=a&to=b", http.StatusBadRequest, 2},
		{"/api/v1/application/metric?query=2xx&from=1&to=123", http.StatusOK, 0},
	}

	for _, test := range testCases {
		t.Run(test.endpoint, func(t *testing.T) {
			rr := httptest.NewRecorder()
			req, err := http.NewRequest("GET", test.endpoint, nil)
			if err != nil {
				t.Fatalf("Http request returned with error")
			}

			ms.api.Router().ServeHTTP(rr, req)
			if rr.Code != test.expectedStatusCode {
				t.Fatalf("handler returned wrong status code: got %v want %v", rr.Code, test.expectedStatusCode)
			}

			if err != nil {
				t.Fatal(err)
			}

			response := map[string]map[string][]string{}
			body, err := ioutil.ReadAll(rr.Body)
			json.Unmarshal(body, &response)

			if len(response["Validation"]) != test.expectedValidationCount {
				t.Fatalf("unexpected validation message count response, got %d, expected %d", len(response["Validation"]), test.expectedValidationCount)
			}
		})
	}

}

func TestApplicationAlertsEndpoint(t *testing.T) {
	alerts := testutil.NewMultipleMockAlerts()
	ms := MockServer(t, "", nil, alerts)
	ms.api.BindEndpoints()
	ms.api.Serve()

	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/v1/application/alerts?tags=foo,foo2&from=1&to=1&provider=foo", nil)
	if err != nil {
		t.Fatalf("Http request returned with error")
	}

	ms.api.Router().ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
	}

	if err != nil {
		t.Fatal(err)
	}

	response := []httpresponse.CheckResponse{}
	body, err := ioutil.ReadAll(rr.Body)
	json.Unmarshal(body, &response)

	if len(response) != 1 {
		t.Fatalf("unexpected metrics endpont response, got %d expected %d", len(response), 1)
	}

}

func TestApplicationAlertsEndpointWithInvalidQueryParameters(t *testing.T) {
	alerts := testutil.NewMultipleMockAlerts()
	ms := MockServer(t, "", nil, alerts)
	ms.api.BindEndpoints()
	ms.api.Serve()

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
				t.Fatalf("Http request returned with error")
			}

			ms.api.Router().ServeHTTP(rr, req)
			if rr.Code != test.expectedStatusCode {
				t.Fatalf("handler returned wrong status code: got %v want %v", rr.Code, test.expectedStatusCode)
			}

			if err != nil {
				t.Fatal(err)
			}

			response := map[string]map[string][]string{}
			body, err := ioutil.ReadAll(rr.Body)
			err = json.Unmarshal(body, &response)

			if len(response["Validation"]) != test.expectedValidationCount {
				t.Fatalf("unexpected validation message count response, got %d, expected %d", len(response["Validation"]), test.expectedValidationCount)
			}
		})
	}

}
