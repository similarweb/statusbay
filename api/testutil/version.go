package testutil

import (
	notifier "github.com/similarweb/client-notifier"
)

type MockVersion struct {
}

func NewMockVersion() *MockVersion {
	return &MockVersion{}
}

func (v *MockVersion) Get() (*notifier.Response, error) {
	return nil, nil
}
