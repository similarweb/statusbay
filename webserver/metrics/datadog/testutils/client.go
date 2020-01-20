package testutil

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"

	"github.com/zorkian/go-datadog-api"
)

func MockQueryMetrics(from int64, to int64, query string) ([]datadog.Series, error) {

	_, filename, _, _ := runtime.Caller(0)
	currentFolderPath := filepath.Dir(filename)
	reader, err := ioutil.ReadFile(fmt.Sprintf("%s/mock/%s.json", currentFolderPath, query))

	if err != nil {
		return nil, err
	}
	data := []datadog.Series{}
	err = json.Unmarshal(reader, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

type MockDatadog struct {
}

func NewMockDatadog() *MockDatadog {
	return &MockDatadog{}
}

func (m *MockDatadog) QueryMetrics(from int64, to int64, query string) ([]datadog.Series, error) {
	return MockQueryMetrics(from, to, query)
}
