package testutil

import (
	"statusbay/api/kubernetes"
	"statusbay/state"
)

var (
	responseTable = []state.TableKubernetes{
		{ID: 1, Name: "foo", Cluster: "cluster1", Namespace: "foo-namespace", Status: "running", Time: 123, DeployBy: "foo@similarweb.com"},
		{ID: 2, Name: "foo", Cluster: "cluster1", Namespace: "foo-namespace", Status: "successful", Time: 1234, DeployBy: "foo@similarweb.com"},
		{ID: 3, Name: "foo-1", Cluster: "cluster1", Namespace: "foo-namespace", Status: "faild", Time: 1234, DeployBy: "foo@similarweb.com"},
	}
)

type MockStorage struct {
	Err error
}

func NewMockStorage() *MockStorage {
	return &MockStorage{}
}

func (m *MockStorage) Applications(queryFillter kubernetes.FilterApplications) (*[]state.TableKubernetes, error) {
	return &responseTable, nil
}

func (m *MockStorage) ApplicationsCount(queryFillter kubernetes.FilterApplications) (int64, error) {
	return 3, nil
}

func (m *MockStorage) GetDeployment(name string, time int64) (state.TableKubernetes, error) {
	return responseTable[0], nil
}
