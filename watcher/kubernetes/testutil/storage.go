package testutil

import (
	kuberneteswatcher "statusbay/watcher/kubernetes"
	"statusbay/watcher/kubernetes/common"
	"strconv"
)

type MockStorageDeployment struct {
	ApplyID string
	Status  common.DeploymentStatus
	Schema  kuberneteswatcher.DBSchema
}

type MockStorage struct {
	MockUpdateDeployment  map[string]MockStorageDeployment
	MockWriteDeployment   map[string]MockStorageDeployment
	MockDeploymentHistory map[string]uint64
	MockFile              string
}

func NewMockStorage() *MockStorage {
	return &MockStorage{
		MockUpdateDeployment:  map[string]MockStorageDeployment{},
		MockWriteDeployment:   map[string]MockStorageDeployment{},
		MockDeploymentHistory: map[string]uint64{},
	}
}
func (m *MockStorage) CreateApply(data *kuberneteswatcher.RegistryRow, status common.DeploymentStatus) (string, error) {

	id := strconv.Itoa(len(m.MockWriteDeployment) + 1)
	m.MockWriteDeployment[id] = MockStorageDeployment{
		ApplyID: id,
		Status:  status,
		Schema:  data.DBSchema,
	}

	return id, nil
}

func (m *MockStorage) UpdateApply(applyID string, data *kuberneteswatcher.RegistryRow, status common.DeploymentStatus) (bool, error) {

	m.MockWriteDeployment[applyID] = MockStorageDeployment{
		ApplyID: applyID,
		Status:  status,
		Schema:  data.DBSchema,
	}

	return true, nil
}

func (m *MockStorage) GetAppliesByStatus(status common.DeploymentStatus) (map[string]kuberneteswatcher.DBSchema, error) {

	return map[string]kuberneteswatcher.DBSchema{}, nil

}

func (m *MockStorage) UpdateAppliesVersionHistory(deploymentName string, hash uint64) bool {

	if _, ok := m.MockDeploymentHistory[deploymentName]; ok {

		return false
	}

	m.MockDeploymentHistory[deploymentName] = hash
	return true

}

func (m *MockStorage) DeleteAppliedVersion(deploymentName string) bool {

	return true

}
