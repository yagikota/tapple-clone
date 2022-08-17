package usecase

import (
	"context"
	"os"
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
)

// テストデータ
var (
	userID = 1
	user11 *model.User
	user12 *model.User

	users1 model.UserSlice

	user11Entity *entity.User
	user12Entity *entity.User

	users1Entity entity.UserSlice

	message11 *model.Message
	message12 *model.Message

	messages1Slice model.MessageSlice

	message11Entity *entity.Message
	message12Entity *entity.Message

	messagesEntity1Slice entity.MessageSlice

	room1 *model.Room

	roomSlice model.RoomSlice
	rooms1    *model.Rooms

	room1Entity      *entity.Room
	room1EntitySlice entity.RoomSlice

	roomUser1Entity *entity.RoomUser
)

func TestMain(m *testing.M) {
	println("before all...")

	room1 = &model.Room{
		ID:       1,
		Unread:   0,
		IsPinned: false,
		Name:     "name1",
		Icon:     "/icon",
		LatestMessage: &model.Message{
			ID:        1,
			UserID:    1,
			Content:   "content",
			CreatedAt: time.Date(2022, 4, 1, 0, 0, 0, 0, time.Local),
		},
	}

	rooms1 = &model.Rooms{
		Rooms: []*model.Room{room1},
	}

	// fmt.Println(room1Entity)
	room1Entity = new(entity.Room)
	room1Entity = &entity.Room{
		ID: 1,
	}
	room1Entity.R = room1Entity.R.NewStruct() //roomにroomRを作成
	room1Entity.R.Messages = entity.MessageSlice{
		{
			ID:        1,
			UserID:    1,
			Content:   "content",
			CreatedAt: time.Date(2022, 4, 1, 0, 0, 0, 0, time.Local),
		},
	}

	room1Entity.R.RoomUsers = entity.RoomUserSlice{
		{
			ID:       1,
			UserID:   1,
			RoomID:   1,
			IsPinned: false,
		},
		{
			ID:       2,
			UserID:   2,
			RoomID:   1,
			IsPinned: false,
		},
	}

	room1Entity.R.RoomUsers[0].R = room1Entity.R.RoomUsers[0].R.NewStruct()
	room1Entity.R.RoomUsers[0].R.User = &entity.User{
		ID:   1,
		Name: "name1",
		Icon: "/icon",
	}

	room1EntitySlice = entity.RoomSlice{room1Entity}

	code := m.Run()

	println("after all...")

	os.Exit(code)
}

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

// func Test_userUsecase_FindAllUsers(t *testing.T) {
// 	type fields struct {
// 		userService service.IUserService
// 	}
// 	type args struct {
// 		ctx context.Context
// 	}
// 	tests := []struct {
// 		name          string
// 		prepareMockFn func(m *mock.MockIUserService)
// 		fields        fields
// 		args          args
// 		want          model.UserSlice
// 		wantErr       error
// 	}{
// 		{
// 			name: "usecase FindAllUsers suceess response",
// 			args: args{
// 				ctx: &gin.Context{},
// 			},
// 			prepareMockFn: func(m *mock.MockIUserService) {
// 				m.EXPECT().FindAllUsers(gomock.Any()).Return(users1Entity, nil)
// 			},
// 			want:    users1,
// 			wantErr: nil,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			gin.SetMode(gin.TestMode)
// 			//mock登録
// 			controller := gomock.NewController(t)
// 			defer controller.Finish()

// 			mock := mock.NewMockIUserService(controller)
// 			tt.prepareMockFn(mock)
// 			uu := NewUserUsecase(mock)
// 			res, err := uu.FindAllUsers(tt.args.ctx)
// 			assert.Equal(t, res, tt.want)
// 			assert.Equal(t, err, tt.wantErr)
// 		})
// 	}
// }

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
		want          *model.Rooms
		wantErr       error
	}{
		{
			name: "usecase FindAllRooms success response",
			args: args{
				ctx: &gin.Context{},
			},
			prepareMockFn: func(m *mock.MockIUserService) {
				m.EXPECT().FindAllRooms(gomock.Any(), userID).Return(room1EntitySlice, nil)
			},
			want: rooms1,
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

			res, err := uu.FindAllRooms(tt.args.ctx, userID)
			assert.Equal(t, res, tt.want)
			assert.Equal(t, err, tt.wantErr)
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
		want    *model.RoomDetail
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
		want    *model.Message
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uu := &userUsecase{
				userService: tt.fields.userService,
			}
			got, err := uu.SendMessage(tt.args.ctx, tt.args.userID, tt.args.roomID, tt.args.m)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.SendMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUsecase.SendMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
