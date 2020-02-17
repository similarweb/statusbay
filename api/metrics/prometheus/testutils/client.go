package testutil

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"

	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
)

type apiResponse struct {
	Status string       `json:"status"`
	Data   dataResponse `json:"data"`
}

type dataResponse struct {
	ResultType model.ValueType `json:"resultType"`
	Results    model.Matrix    `json:"result"`
}

func MockQueryRange(ctx context.Context, query string, r v1.Range) (model.Value, v1.Warnings, error) {

	_, filename, _, _ := runtime.Caller(0)
	currentFolderPath := filepath.Dir(filename)
	reader, err := ioutil.ReadFile(fmt.Sprintf("%s/mock/%s.json", currentFolderPath, query))

	if err != nil {
		return nil, nil, err
	}

	var apiRes apiResponse
	err = json.Unmarshal(reader, &apiRes)
	if err != nil {
		return nil, nil, err
	}

	return apiRes.Data.Results, nil, nil
}

type MockPrometheus struct {
}

func NewMockPrometheus() *MockPrometheus {
	return &MockPrometheus{}
}

func (m *MockPrometheus) QueryRange(ctx context.Context, query string, r v1.Range) (model.Value, v1.Warnings, error) {
	return MockQueryRange(ctx, query, r)
}
