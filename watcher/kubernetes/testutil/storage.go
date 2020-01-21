package testutil

import (
	kuberneteswatcher "statusbay/watcher/kubernetes"
	"statusbay/watcher/kubernetes/common"
)

type MockStorageDeployment struct {
	ID     uint
	Status common.DeploymentStatus
	Schema kuberneteswatcher.DBSchema
}

type MockStorage struct {
	MockUpdateDeployment  map[uint]MockStorageDeployment
	MockWriteDeployment   map[uint]MockStorageDeployment
	MockDeploymentHistory map[string]uint64
	MockFile              string
}

func NewMockStorage() *MockStorage {
	return &MockStorage{
		MockUpdateDeployment:  map[uint]MockStorageDeployment{},
		MockWriteDeployment:   map[uint]MockStorageDeployment{},
		MockDeploymentHistory: map[string]uint64{},
	}
}
func (m *MockStorage) CreateApply(data *kuberneteswatcher.RegistryRow, status common.DeploymentStatus) (uint, error) {

	id := uint(len(m.MockWriteDeployment) + 1)
	m.MockWriteDeployment[id] = MockStorageDeployment{
		ID:     id,
		Status: status,
		Schema: data.DBSchema,
	}

	return id, nil
}

func (m *MockStorage) UpdateApply(id uint, data *kuberneteswatcher.RegistryRow, status common.DeploymentStatus) (bool, error) {

	m.MockWriteDeployment[id] = MockStorageDeployment{
		ID:     id,
		Status: status,
		Schema: data.DBSchema,
	}

	return true, nil
}

func (m *MockStorage) GetAppliesByStatus(status common.DeploymentStatus) (map[uint]kuberneteswatcher.DBSchema, error) {

	return map[uint]kuberneteswatcher.DBSchema{}, nil

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
