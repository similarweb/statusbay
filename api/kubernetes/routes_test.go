package kubernetes_test

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"statusbay/api"
	"statusbay/api/alerts"
	"statusbay/api/kubernetes"
	"statusbay/api/kubernetes/testutil"
	"statusbay/api/metrics"
	"sync"
	"testing"
)

type testServer struct {
	api *api.Server
}

func MockServer(t *testing.T, storageMockFile string, metrics map[string]metrics.MetricManagerDescriber, alertsClient map[string]alerts.AlertsManagerDescriber) testServer {

	storage := testutil.NewMockStorage()
	return testServer{
		api: api.NewServer(storage, "8080", "api/kubernetes/testutil/events.yaml", metrics, alertsClient),
	}
}

func TestApplicationsData(t *testing.T) {
	var wg *sync.WaitGroup
	ctx := context.Background()

	ms := MockServer(t, "", nil, nil)
	ms.api.BindEndpoints()
	ms.api.Serve(ctx, wg)

	testsResponseCount := []struct {
		endpoint              string
		expectedStatusCode    int
		expectedCountResponse int
	}{
		{"/api/v1/kubernetes/applications", http.StatusOK, 3},
	}

	for _, test := range testsResponseCount {
		t.Run(test.endpoint, func(t *testing.T) {
			rr := httptest.NewRecorder()
			req, err := http.NewRequest("GET", test.endpoint, nil)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			ms.api.Router().ServeHTTP(rr, req)
			if rr.Code != test.expectedStatusCode {
				t.Fatalf("unexpected status code: got %d want %d", rr.Code, test.expectedStatusCode)
			}

			response := &kubernetes.ResponseKubernetesApplicationsCount{}
			body, err := ioutil.ReadAll(rr.Body)
			err = json.Unmarshal(body, &response)
			if len(response.Records) != test.expectedCountResponse {
				t.Fatalf("unexpected deployment events length, got %d expected %d", len(response.Records), test.expectedCountResponse)
			}
		})
	}

}

func TestApplicationsFiltersData(t *testing.T) {
	var wg *sync.WaitGroup
	ctx := context.Background()

	ms := MockServer(t, "", nil, nil)
	ms.api.BindEndpoints()
	ms.api.Serve(ctx, wg)

	testsResponseCount := []struct {
		endpoint              string
		expectedStatusCode    int
		expectedCountResponse int
	}{
		{"/api/v1/kubernetes/applications/values/cluster", http.StatusOK, 2},
		{"/api/v1/kubernetes/applications/values/none", http.StatusBadRequest, 0},
	}

	for _, test := range testsResponseCount {
		t.Run(test.endpoint, func(t *testing.T) {

			rr := httptest.NewRecorder()
			req, err := http.NewRequest("GET", test.endpoint, nil)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			ms.api.Router().ServeHTTP(rr, req)
			if rr.Code != test.expectedStatusCode {
				t.Fatalf("unexpected status code: got %d want %d", rr.Code, test.expectedStatusCode)
			}

			response := []string{}
			body, err := ioutil.ReadAll(rr.Body)
			err = json.Unmarshal(body, &response)

			if len(response) != test.expectedCountResponse {
				t.Fatalf("unexpected filters response length, got %d expected %d", len(response), test.expectedCountResponse)
			}

		})
	}

}
