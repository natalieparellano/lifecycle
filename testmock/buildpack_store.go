// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/buildpacks/lifecycle (interfaces: BuildpackStore)

// Package testmock is a generated GoMock package.
package testmock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	buildpack "github.com/buildpacks/lifecycle/buildpack"
)

// MockBuildpackStore is a mock of BuildpackStore interface.
type MockBuildpackStore struct {
	ctrl     *gomock.Controller
	recorder *MockBuildpackStoreMockRecorder
}

// MockBuildpackStoreMockRecorder is the mock recorder for MockBuildpackStore.
type MockBuildpackStoreMockRecorder struct {
	mock *MockBuildpackStore
}

// NewMockBuildpackStore creates a new mock instance.
func NewMockBuildpackStore(ctrl *gomock.Controller) *MockBuildpackStore {
	mock := &MockBuildpackStore{ctrl: ctrl}
	mock.recorder = &MockBuildpackStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuildpackStore) EXPECT() *MockBuildpackStoreMockRecorder {
	return m.recorder
}

// Lookup mocks base method.
func (m *MockBuildpackStore) Lookup(arg0, arg1 string) (buildpack.Buildpack, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Lookup", arg0, arg1)
	ret0, _ := ret[0].(buildpack.Buildpack)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Lookup indicates an expected call of Lookup.
func (mr *MockBuildpackStoreMockRecorder) Lookup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lookup", reflect.TypeOf((*MockBuildpackStore)(nil).Lookup), arg0, arg1)
}
