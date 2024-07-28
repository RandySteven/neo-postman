// +build !testmock

// Code generated by mockery v2.43.1. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// DevHandler is an autogenerated mock type for the DevHandler type
type DevHandler struct {
	mock.Mock
}

// DummyTester provides a mock function with given fields: w, r
func (_m *DevHandler) DummyTester(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// GetListUrl provides a mock function with given fields: w, r
func (_m *DevHandler) GetListUrl(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// Hello provides a mock function with given fields: w, r
func (_m *DevHandler) Hello(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// NewDevHandler creates a new instance of DevHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDevHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *DevHandler {
	mock := &DevHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
