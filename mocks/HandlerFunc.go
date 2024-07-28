// +build !testmock

// Code generated by mockery v2.43.1. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// HandlerFunc is an autogenerated mock type for the HandlerFunc type
type HandlerFunc struct {
	mock.Mock
}

// Execute provides a mock function with given fields: w, r
func (_m *HandlerFunc) Execute(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// NewHandlerFunc creates a new instance of HandlerFunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewHandlerFunc(t interface {
	mock.TestingT
	Cleanup(func())
}) *HandlerFunc {
	mock := &HandlerFunc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
