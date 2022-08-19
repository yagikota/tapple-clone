package usecase

import (
	"context"
	"os"
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

	postMessage *entity.Message

	messages1Slice model.MessageSlice

	message11Entity *entity.Message
	message12Entity *entity.Message

	messagesEntity1Slice entity.MessageSlice

	roomID = 1
	room1  *model.Room

	roomSlice model.RoomSlice
	rooms1    *model.Rooms

	room1Entity      *entity.Room
	room1EntitySlice entity.RoomSlice

	roomUser1Entity *entity.RoomUser

	roomDetail1       *model.RoomDetail
	roomDetailUsers1  *model.UserSlice
	roomDetail1Entity *entity.Room

	newMessage     *model.NewMessage
	createdMessage *entity.Message
)

func TestMain(m *testing.M) {
	println("before all...")

	user11Entity = &entity.User{
		ID:       1,
		Name:     "name1",
		Icon:     "icon1",
		Gender:   0,
		Birthday: time.Date(2022, 4, 1, 0, 0, 0, 0, time.UTC),
		Location: 0,
	}

	user12Entity = &entity.User{
		ID:       2,
		Name:     "name2",
		Icon:     "icon2",
		Gender:   1,
		Birthday: time.Date(2022, 4, 1, 0, 0, 0, 0, time.UTC),
		Location: 1,
	}

	users1Entity = entity.UserSlice{user11Entity, user12Entity}

	user11 = &model.User{
		ID:       1,
		Name:     "name1",
		Icon:     "icon1",
		Gender:   0,
		BirthDay: time.Date(2022, 4, 1, 0, 0, 0, 0, time.UTC),
		Location: 0,
	}

	user12 = &model.User{
		ID:       2,
		Name:     "name2",
		Icon:     "icon2",
		Gender:   1,
		BirthDay: time.Date(2022, 4, 1, 0, 0, 0, 0, time.UTC),
		Location: 1,
	}

	users1 = model.UserSlice{user11, user12}

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
			CreatedAt: time.Date(2022, 4, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	rooms1 = &model.Rooms{
		Rooms: []*model.Room{room1},
	}

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
			CreatedAt: time.Date(2022, 4, 1, 0, 0, 0, 0, time.UTC),
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

	roomDetail1Entity = new(entity.Room)
	roomDetail1Entity = &entity.Room{
		ID: 1,
	}
	roomDetail1Entity.R = roomDetail1Entity.R.NewStruct()
	roomDetail1Entity.R.RoomUsers = entity.RoomUserSlice{
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
	roomDetail1Entity.R.RoomUsers[0].R = roomDetail1Entity.R.RoomUsers[0].R.NewStruct()
	roomDetail1Entity.R.RoomUsers[1].R = roomDetail1Entity.R.RoomUsers[0].R.NewStruct()
	roomDetail1Entity.R.RoomUsers[0].R.User = &entity.User{
		ID:       2,
		Name:     "name2",
		Icon:     "icon2",
		Gender:   1,
		Birthday: time.Date(2022, 4, 1, 0, 0, 0, 0, time.UTC),
		Location: 1,
	}
	roomDetail1Entity.R.RoomUsers[1].R.User = &entity.User{
		ID:       1,
		Name:     "name1",
		Icon:     "icon1",
		Gender:   0,
		Birthday: time.Date(2022, 4, 1, 0, 0, 0, 0, time.UTC),
		Location: 0,
	}

	roomDetail1Entity.R.Messages = entity.MessageSlice{
		{
			ID:        1,
			UserID:    1,
			Content:   "content",
			CreatedAt: time.Date(2022, 4, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	message11 = &model.Message{
		ID:        1,
		UserID:    1,
		Content:   "content",
		CreatedAt: time.Date(2022, 4, 1, 0, 0, 0, 0, time.UTC),
	}

	roomDetail1 = &model.RoomDetail{
		ID:       1,
		Name:     "name2",
		Icon:     "icon2",
		Users:    []*model.User{user12, user11},
		Messages: []*model.Message{message11},
	}

	postMessage = &entity.Message{
		UserID:  1,
		RoomID:  1,
		Content: "content",
	}

	createdMessage = &entity.Message{
		ID:        1,
		UserID:    1,
		RoomID:    1,
		Content:   "content",
		CreatedAt: time.Date(2022, 4, 1, 0, 0, 0, 0, time.UTC),
	}

	newMessage = &model.NewMessage{
		Content: "content",
	}

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
			want:    rooms1,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
		name          string
		prepareMockFn func(m *mock.MockIUserService)
		fields        fields
		args          args
		want          *model.RoomDetail
		wantErr       error
	}{
		{
			name: "usecase FindRoomDetailByRoomID success reponse",
			args: args{
				ctx: &gin.Context{},
			},
			prepareMockFn: func(m *mock.MockIUserService) {
				m.EXPECT().FindRoomDetailByRoomID(gomock.Any(), userID, roomID).Return(roomDetail1Entity, nil)
			},
			want:    roomDetail1,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//mock登録
			controller := gomock.NewController(t)
			defer controller.Finish()

			mock := mock.NewMockIUserService(controller)
			tt.prepareMockFn(mock)
			uu := NewUserUsecase(mock)

			res, err := uu.FindRoomDetailByRoomID(tt.args.ctx, userID, roomID)
			assert.Equal(t, res, tt.want)
			assert.Equal(t, err, tt.wantErr)
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
		name          string
		prepareMockFn func(m *mock.MockIUserService)
		fields        fields
		args          args
		want          *model.Message
		wantErr       error
	}{
		{
			name: "usecase SendMessage success response",
			args: args{
				ctx: &gin.Context{},
			},
			prepareMockFn: func(m *mock.MockIUserService) {
				m.EXPECT().SendMessage(gomock.Any(), postMessage).Return(createdMessage, nil)
			},
			want:    message11,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//mock登録
			controller := gomock.NewController(t)
			defer controller.Finish()

			mock := mock.NewMockIUserService(controller)
			tt.prepareMockFn(mock)
			uu := NewUserUsecase(mock)
			res, err := uu.SendMessage(tt.args.ctx, userID, roomID, newMessage)
			assert.Equal(t, res, tt.want)
			assert.Equal(t, err, tt.wantErr)
		})
	}
}
