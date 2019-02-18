// Code generated by MockGen. DO NOT EDIT.
// Source: ./docker.go

// Package mock_docker is a generated GoMock package.
package mock_docker

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIClient is a mock of IClient interface
type MockIClient struct {
	ctrl     *gomock.Controller
	recorder *MockIClientMockRecorder
}

// MockIClientMockRecorder is the mock recorder for MockIClient
type MockIClientMockRecorder struct {
	mock *MockIClient
}

// NewMockIClient creates a new mock instance
func NewMockIClient(ctrl *gomock.Controller) *MockIClient {
	mock := &MockIClient{ctrl: ctrl}
	mock.recorder = &MockIClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIClient) EXPECT() *MockIClientMockRecorder {
	return m.recorder
}

// ListImagesWithReference mocks base method
func (m *MockIClient) ListImagesWithReference(ref string) ([]string, error) {
	ret := m.ctrl.Call(m, "ListImagesWithReference", ref)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListImagesWithReference indicates an expected call of ListImagesWithReference
func (mr *MockIClientMockRecorder) ListImagesWithReference(ref interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListImagesWithReference", reflect.TypeOf((*MockIClient)(nil).ListImagesWithReference), ref)
}

// Pull mocks base method
func (m *MockIClient) Pull(image string) error {
	ret := m.ctrl.Call(m, "Pull", image)
	ret0, _ := ret[0].(error)
	return ret0
}

// Pull indicates an expected call of Pull
func (mr *MockIClientMockRecorder) Pull(image interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pull", reflect.TypeOf((*MockIClient)(nil).Pull), image)
}
