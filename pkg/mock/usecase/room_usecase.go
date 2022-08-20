// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/usecase/room_usecase.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	model "github.com/CyberAgentHack/2208-ace-go-server/pkg/usecase/model"
	gomock "github.com/golang/mock/gomock"
)

// MockIRoomUsecase is a mock of IRoomUsecase interface.
type MockIRoomUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockIRoomUsecaseMockRecorder
}

// MockIRoomUsecaseMockRecorder is the mock recorder for MockIRoomUsecase.
type MockIRoomUsecaseMockRecorder struct {
	mock *MockIRoomUsecase
}

// NewMockIRoomUsecase creates a new mock instance.
func NewMockIRoomUsecase(ctrl *gomock.Controller) *MockIRoomUsecase {
	mock := &MockIRoomUsecase{ctrl: ctrl}
	mock.recorder = &MockIRoomUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRoomUsecase) EXPECT() *MockIRoomUsecaseMockRecorder {
	return m.recorder
}

// FindAllRooms mocks base method.
func (m *MockIRoomUsecase) FindAllRooms(ctx context.Context, userID int) (*model.Rooms, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllRooms", ctx, userID)
	ret0, _ := ret[0].(*model.Rooms)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllRooms indicates an expected call of FindAllRooms.
func (mr *MockIRoomUsecaseMockRecorder) FindAllRooms(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllRooms", reflect.TypeOf((*MockIRoomUsecase)(nil).FindAllRooms), ctx, userID)
}

// FindRoomDetailByRoomID mocks base method.
func (m *MockIRoomUsecase) FindRoomDetailByRoomID(ctx context.Context, userID, roomID, messageID int) (*model.RoomDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindRoomDetailByRoomID", ctx, userID, roomID, messageID)
	ret0, _ := ret[0].(*model.RoomDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindRoomDetailByRoomID indicates an expected call of FindRoomDetailByRoomID.
func (mr *MockIRoomUsecaseMockRecorder) FindRoomDetailByRoomID(ctx, userID, roomID, messageID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindRoomDetailByRoomID", reflect.TypeOf((*MockIRoomUsecase)(nil).FindRoomDetailByRoomID), ctx, userID, roomID, messageID)
}

// SendMessage mocks base method.
func (m_2 *MockIRoomUsecase) SendMessage(ctx context.Context, userID, roomID int, m *model.NewMessage) (*model.Message, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMessage", ctx, userID, roomID, m)
	ret0, _ := ret[0].(*model.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendMessage indicates an expected call of SendMessage.
func (mr *MockIRoomUsecaseMockRecorder) SendMessage(ctx, userID, roomID, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockIRoomUsecase)(nil).SendMessage), ctx, userID, roomID, m)
}
