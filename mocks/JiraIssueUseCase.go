// +build !testmock

// Code generated by mockery v2.43.1. DO NOT EDIT.

package mocks

import (
	context "context"

	apperror "github.com/RandySteven/neo-postman/apperror"

	mock "github.com/stretchr/testify/mock"

	requests "github.com/RandySteven/neo-postman/entities/payloads/requests"

	responses "github.com/RandySteven/neo-postman/entities/payloads/responses"
)

// JiraIssueUseCase is an autogenerated mock type for the JiraIssueUseCase type
type JiraIssueUseCase struct {
	mock.Mock
}

// CreateJiraTicket provides a mock function with given fields: ctx, request
func (_m *JiraIssueUseCase) CreateJiraTicket(ctx context.Context, request *requests.CreateJiraIssueRequest) (*responses.CreateJiraIssueResponse, *apperror.CustomError) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for CreateJiraTicket")
	}

	var r0 *responses.CreateJiraIssueResponse
	var r1 *apperror.CustomError
	if rf, ok := ret.Get(0).(func(context.Context, *requests.CreateJiraIssueRequest) (*responses.CreateJiraIssueResponse, *apperror.CustomError)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *requests.CreateJiraIssueRequest) *responses.CreateJiraIssueResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*responses.CreateJiraIssueResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *requests.CreateJiraIssueRequest) *apperror.CustomError); ok {
		r1 = rf(ctx, request)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*apperror.CustomError)
		}
	}

	return r0, r1
}

// GetAllJiraTickets provides a mock function with given fields: ctx
func (_m *JiraIssueUseCase) GetAllJiraTickets(ctx context.Context) ([]*responses.JiraIssueListResponse, *apperror.CustomError) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAllJiraTickets")
	}

	var r0 []*responses.JiraIssueListResponse
	var r1 *apperror.CustomError
	if rf, ok := ret.Get(0).(func(context.Context) ([]*responses.JiraIssueListResponse, *apperror.CustomError)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*responses.JiraIssueListResponse); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*responses.JiraIssueListResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) *apperror.CustomError); ok {
		r1 = rf(ctx)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*apperror.CustomError)
		}
	}

	return r0, r1
}

// NewJiraIssueUseCase creates a new instance of JiraIssueUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewJiraIssueUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *JiraIssueUseCase {
	mock := &JiraIssueUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
