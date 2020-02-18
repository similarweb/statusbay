package testutil

import (
	"statusbay/api/kubernetes"
	"statusbay/state"
)

var (
	responseTable = []state.TableKubernetes{
		{ApplyId: "c60c45dc08b369ec8a4ee89bcf37c96eaa1b81cb", Name: "foo", Cluster: "cluster1", Namespace: "foo-namespace", Status: "running", Time: 123, DeployBy: "foo@example.com"},
		{ApplyId: "cbd69b781769cbf090662f46dd3bbef10f3103c2", Name: "foo", Cluster: "cluster1", Namespace: "foo-namespace", Status: "successful", Time: 1234, DeployBy: "foo@example.com"},
		{ApplyId: "asdmken3rnuiweu423ihndscsdfalwelk2223usd", Name: "foo-1", Cluster: "cluster1", Namespace: "foo-namespace", Status: "faild", Time: 1234, DeployBy: "foo@example.com"},
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

func (m *MockStorage) GetDeployment(applyID string) (state.TableKubernetes, error) {
	return responseTable[0], nil
}

func (m *MockStorage) GetUniqueFieldValues(tableName, columnName string) ([]string, error) {
	values := []string{"foo", "foo1"}
	return values, nil
}
