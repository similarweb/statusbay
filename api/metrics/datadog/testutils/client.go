package testutil

import (
	"context"
	"encoding/json"
	"fmt"
	_nethttp "net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

type MockMetricsApi struct {
}

func NewMockMockMetricsApi() *MockMetricsApi {
	return &MockMetricsApi{}
}

func (m *MockMetricsApi) QueryMetrics(ctx context.Context, from int64, to int64, query string) (datadogV1.MetricsQueryResponse, *_nethttp.Response, error) {
	var (
		resp     datadogV1.MetricsQueryResponse
		httpResp *_nethttp.Response
	)
	_, filename, _, _ := runtime.Caller(0)

	// Read directory contents
	mockQueryDir := filepath.Join(filepath.Dir(filename), "mock")
	files, err := os.ReadDir(mockQueryDir)
	if err != nil {
		return resp, httpResp, err
	}

	// Create a list of accessible queries
	mockQueryList := make(map[string]string, len(files))
	for _, file := range files {
		if !file.IsDir() {
			mockQueryName := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
			mockQueryFilePath := filepath.Join(mockQueryDir, file.Name())
			mockQueryList[mockQueryName] = mockQueryFilePath
		}
	}

	// Check if the query parameter is among the accessible queries
	mockQueryFilePath, found := mockQueryList[query]
	if !found {
		return resp, httpResp, fmt.Errorf("Mock file not found for query: %s", query)
	}

	reader, err := os.ReadFile(mockQueryFilePath)
	if err != nil {
		return resp, httpResp, err
	}

	err = json.Unmarshal(reader, &resp)
	if err != nil {
		fmt.Printf("Error reading mock response:\n %s\n", err)
		return resp, httpResp, err
	}

	err = json.Unmarshal(reader, &resp)
	if err != nil {
		fmt.Printf("Error reading mock response:\n %s\n", err)
		return resp, httpResp, err
	}
	return resp, httpResp, nil
}
