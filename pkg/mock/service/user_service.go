// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/domain/service/user_service.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	model "github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/model"
	gomock "github.com/golang/mock/gomock"
)

// MockIUserService is a mock of IUserService interface.
type MockIUserService struct {
	ctrl     *gomock.Controller
	recorder *MockIUserServiceMockRecorder
}

// MockIUserServiceMockRecorder is the mock recorder for MockIUserService.
type MockIUserServiceMockRecorder struct {
	mock *MockIUserService
}

// NewMockIUserService creates a new mock instance.
func NewMockIUserService(ctrl *gomock.Controller) *MockIUserService {
	mock := &MockIUserService{ctrl: ctrl}
	mock.recorder = &MockIUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserService) EXPECT() *MockIUserServiceMockRecorder {
	return m.recorder
}

// FindAllUsers mocks base method.
func (m *MockIUserService) FindAllUsers(ctx context.Context) (model.UserSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllUsers", ctx)
	ret0, _ := ret[0].(model.UserSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllUsers indicates an expected call of FindAllUsers.
func (mr *MockIUserServiceMockRecorder) FindAllUsers(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllUsers", reflect.TypeOf((*MockIUserService)(nil).FindAllUsers), ctx)
}

// FindUserByUserID mocks base method.
func (m *MockIUserService) FindUserByUserID(ctx context.Context, userID int) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByUserID", ctx, userID)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByUserID indicates an expected call of FindUserByUserID.
func (mr *MockIUserServiceMockRecorder) FindUserByUserID(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByUserID", reflect.TypeOf((*MockIUserService)(nil).FindUserByUserID), ctx, userID)
}

// FindUserDetailByUserID mocks base method.
func (m *MockIUserService) FindUserDetailByUserID(ctx context.Context, userID int) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserDetailByUserID", ctx, userID)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserDetailByUserID indicates an expected call of FindUserDetailByUserID.
func (mr *MockIUserServiceMockRecorder) FindUserDetailByUserID(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserDetailByUserID", reflect.TypeOf((*MockIUserService)(nil).FindUserDetailByUserID), ctx, userID)
}
