package kubernetes_test

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path"
	"path/filepath"
	"runtime"
	"statusbay/api"
	"statusbay/api/alerts"
	"statusbay/api/kubernetes"
	"statusbay/api/metrics"
	"statusbay/api/testutil"
	"statusbay/config"
	"sync"
	"testing"
)

type testServer struct {
	api *api.Server
}

func MockServer(t *testing.T, storageMockFile string, metrics map[string]metrics.MetricManagerDescriber, alertsClient map[string]alerts.AlertsManagerDescriber) testServer {

	_, filename, _, _ := runtime.Caller(0)
	dir, err := filepath.Abs(path.Join(filepath.Dir(filename), "testutil", "pod-logs"))
	if err != nil {
		dir = ""
	}

	version := testutil.NewMockVersion()
	storage := testutil.NewMockStorage()
	return testServer{
		api: api.NewServer(storage, "8080", config.KubernetesMarksEvents{}, metrics, alertsClient, version, dir),
	}
}

func TestApplicationsData(t *testing.T) {
	var wg sync.WaitGroup
	ctx := context.Background()

	ms := MockServer(t, "", nil, nil)
	ms.api.BindEndpoints()
	ms.api.Serve(ctx, &wg)

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
	var wg sync.WaitGroup
	ctx := context.Background()

	ms := MockServer(t, "", nil, nil)
	ms.api.BindEndpoints()
	ms.api.Serve(ctx, &wg)

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

func TestPodLogs(t *testing.T) {
	var wg sync.WaitGroup
	ctx, cancelFn := context.WithCancel(context.Background())
	defer cancelFn()

	ms := MockServer(t, "", nil, nil)
	ms.api.BindEndpoints()
	ms.api.Serve(ctx, &wg)

	t.Run("valid response", func(t *testing.T) {

		testsResponseExcepted := []struct {
			ContainerName string
			Lines         int
		}{
			{"statusbay-deployment-0.log", 12},
			{"statusbay-deployment-1.log", 2},
		}

		rr := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/api/v1/kubernetes/application/492542c49ffa6b2f2f0fa34b2e666cf523af74ca/logs/pod/statusbay-deployment-0-584ccd4785-2qkk9", nil)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		ms.api.Router().ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Fatalf("unexpected status code: got %d want %d", rr.Code, http.StatusOK)
		}

		response := []kubernetes.ResponseContainerLogs{}
		body, err := ioutil.ReadAll(rr.Body)
		err = json.Unmarshal(body, &response)

		if len(response) != len(testsResponseExcepted) {
			t.Fatalf("unexpected container logs length, got %d expected %d", len(response), len(testsResponseExcepted))
		}

		for i, r := range response {
			if r.ContainerName != testsResponseExcepted[i].ContainerName {
				t.Fatalf("unexpected container name, got %s expected %s", r.ContainerName, testsResponseExcepted[i].ContainerName)
			}
			if len(r.Lines) != testsResponseExcepted[i].Lines {
				t.Fatalf("unexpected container: %s logs lines,  got %d expected %d", r.ContainerName, len(r.Lines), testsResponseExcepted[i].Lines)
			}
		}
	})
}
