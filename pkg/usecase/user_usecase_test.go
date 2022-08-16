package usecase

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/entity"
	"github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/service"
	mock "github.com/CyberAgentHack/2208-ace-go-server/pkg/mock/service"
	"github.com/CyberAgentHack/2208-ace-go-server/pkg/usecase/model"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"github.com/volatiletech/null/v8"
)

// テストデータ
var (
	userID = 1

	user11 = &model.User{
		ID:       1,
		Name:     "name1",
		Icon:     "/icon1",
		Gender:   1,
		BirthDay: time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
		Location: 1,
	}
	user12 = &model.User{
		ID:       2,
		Name:     "name2",
		Icon:     "/icon2",
		Gender:   2,
		BirthDay: time.Date(2022, 2, 2, 0, 0, 0, 0, time.Local),
		Location: 2,
	}
	users1 = model.UserSlice{user11, user12}

	user11Entity = &entity.User{
		ID:       1,
		Name:     "name1",
		Icon:     "/icon1",
		Gender:   1,
		Birthday: time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
		Location: 1,
	}
	user12Entity = &entity.User{
		ID:       2,
		Name:     "name2",
		Icon:     "/icon2",
		Gender:   2,
		Birthday: time.Date(2022, 2, 2, 0, 0, 0, 0, time.Local),
		Location: 2,
	}

	users1Entity = entity.UserSlice{user11Entity, user12Entity}

	message1 = &model.Message{
		ID:        1,
		UserID:    1,
		Content:   "content1",
		CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
	}
	messageSlice1 = model.MessageSlice{message1}

	message11Entity = &entity.Message{
		ID:        1,
		UserID:    1,
		Content:   "content1",
		CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
	}
	messageEntitySlice1 = entity.MessageSlice{message11Entity}

	roomID = 1
	room1  = &model.Room{
		ID:            model.RoomID(roomID),
		Unread:        1,
		IsPinned:      true,
		Name:          "name1",
		Icon:          "/icon1",
		LatestMessage: message1,
	}
	roomSlice = model.RoomSlice{room1}
	rooms1    = &model.Rooms{Rooms: roomSlice}

	// roomR1Entity = &entity.

	room1Entity = &entity.Room{
		ID:        0,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeteledAt: null.Time{
			Time:  time.Time{},
			Valid: false,
		},
		// R: &entity.roomR{
		// 	Messages:  []*entity.Message{},
		// 	RoomUsers: []*entity.RoomUser{},
		// },
		// L: entity.roomL{},
	}
)

// type UserUsecaseTestSuite struct {
// 	suite.Suite
// 	router  *gin.Engine
// 	mock    *mock.MockIUserService
// 	usecase *userUsecase
// }

// func (suite *UserUsecaseTestSuite) SetupSuite() {
// 	gin.SetMode(gin.TestMode)

// 	suite.router = gin.Default()
// 	mockCtl := gomock.NewController(suite.T())
// 	defer mockCtl.Finish()
// 	suite.mock = mock.NewMockIUserService(mockCtl)
// }

// func TestUserHandlerSuite(t *testing.T) {
// 	suite.Run(t, new(UserUsecaseTestSuite))
// }

func Test_userUsecase_FindUserByUserID(t *testing.T) {
	type fields struct {
		userService service.IUserService
	}
	type args struct {
		ctx    context.Context
		userID int
	}
	tests := []struct {
		name          string
		prepareMockFn func(m *mock.MockIUserService)
		fields        fields
		args          args
		want          *model.User
		wantErr       error
	}{
		{
			name: "usecase FindUserByUserID success response",
			args: args{
				ctx:    &gin.Context{},
				userID: 1,
			},
			prepareMockFn: func(m *mock.MockIUserService) {
				m.EXPECT().FindUserByUserID(gomock.Any(), userID).Return(user11Entity, nil)
			},
			want:    user11,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			//mock登録
			controller := gomock.NewController(t)
			defer controller.Finish()

			mock := mock.NewMockIUserService(controller)
			tt.prepareMockFn(mock)
			uu := NewUserUsecase(mock)
			res, err := uu.FindUserByUserID(tt.args.ctx, tt.args.userID)
			assert.Equal(t, res, tt.want)
			assert.Equal(t, err, tt.wantErr)
		})
	}
}

func Test_userUsecase_FindAllUsers(t *testing.T) {
	type fields struct {
		userService service.IUserService
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name          string
		prepareMockFn func(m *mock.MockIUserService)
		fields        fields
		args          args
		want          model.UserSlice
		wantErr       error
	}{
		{
			name: "usecase FindAllUsers suceess response",
			args: args{
				ctx: &gin.Context{},
			},
			prepareMockFn: func(m *mock.MockIUserService) {
				m.EXPECT().FindAllUsers(gomock.Any()).Return(users1Entity, nil)
			},
			want:    users1,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			//mock登録
			controller := gomock.NewController(t)
			defer controller.Finish()

			mock := mock.NewMockIUserService(controller)
			tt.prepareMockFn(mock)
			uu := NewUserUsecase(mock)
			res, err := uu.FindAllUsers(tt.args.ctx)
			assert.Equal(t, res, tt.want)
			assert.Equal(t, err, tt.wantErr)
		})
	}
}

func Test_userUsecase_FindAllRooms(t *testing.T) {
	type fields struct {
		userService service.IUserService
	}
	type args struct {
		ctx    context.Context
		userID int
	}
	tests := []struct {
		name          string
		prepareMockFn func(m *mock.MockIUserService)
		fields        fields
		args          args
		want          model.RoomSlice
		wantErr       bool
	}{
		{
			name: "usecase FindAllRooms success response",
			args: args{
				ctx: &gin.Context{},
			},
			prepareMockFn: func(m *mock.MockIUserService) {
				m.EXPECT().FindAllRooms(gomock.Any(), userID).Return(rooms1, nil)
			},
			want: rooms1.Rooms,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uu := &userUsecase{
				userService: tt.fields.userService,
			}
			got, err := uu.FindAllRooms(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.FindAllRooms() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUsecase.FindAllRooms() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userUsecase_FindRoomDetailByRoomID(t *testing.T) {
	type fields struct {
		userService service.IUserService
	}
	type args struct {
		ctx    context.Context
		userID int
		roomID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Room
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uu := &userUsecase{
				userService: tt.fields.userService,
			}
			got, err := uu.FindRoomDetailByRoomID(tt.args.ctx, tt.args.userID, tt.args.roomID)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.FindRoomDetailByRoomID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUsecase.FindRoomDetailByRoomID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userUsecase_SendMessage(t *testing.T) {
	type fields struct {
		userService service.IUserService
	}
	type args struct {
		ctx    context.Context
		userID int
		roomID int
		m      *model.NewMessage
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uu := &userUsecase{
				userService: tt.fields.userService,
			}
			if err := uu.SendMessage(tt.args.ctx, tt.args.userID, tt.args.roomID, tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.SendMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
