package usecase

import (
	"strconv"
	"testing"
	"time"

	"github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/entity"
	mock "github.com/CyberAgentHack/2208-ace-go-server/pkg/mock/service"
	"github.com/CyberAgentHack/2208-ace-go-server/pkg/usecase/model"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

// ----- BEGIN デフォルトのテストデータ -----
var (
	default_time = time.Date(2022, 4, 1, 0, 0, 0, 0, time.UTC)
	userID       = 1
	roomID       = 1
)

func prepareUserEntity(id, gender, location int) *entity.User {
	return &entity.User{
		ID:       id,
		Name:     "name" + strconv.Itoa(id),
		Icon:     "icon" + strconv.Itoa(id),
		Gender:   gender,
		Birthday: default_time,
		Location: location,
	}
}

func prepareUser(id, gender, location int) *model.User {
	return &model.User{
		ID:       model.UserID(id),
		Name:     "name" + strconv.Itoa(id),
		Icon:     "icon" + strconv.Itoa(id),
		Gender:   gender,
		BirthDay: default_time,
		Location: location,
	}
}

func prepareRoomEntity(id int) *entity.Room {
	room := new(entity.Room)
	room.ID = id
	room.R = room.R.NewStruct()
	room.R.Messages = entity.MessageSlice{
		{
			ID:        1,
			UserID:    1,
			Content:   "content",
			CreatedAt: default_time,
		},
	}
	room.R.RoomUsers = entity.RoomUserSlice{
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
	room.R.RoomUsers[0].R = room.R.RoomUsers[0].R.NewStruct()
	room.R.RoomUsers[0].R.User = &entity.User{
		ID:   2,
		Name: "name2",
		Icon: "icon2",
	}
	return room
}

func prepareRoom(id int) *model.Room {
	return &model.Room{
		ID:       model.RoomID(id),
		Unread:   0,
		IsPinned: false,
		Name:     "name2",
		Icon:     "icon2",
		LatestMessage: &model.Message{
			ID:        1,
			UserID:    1,
			Content:   "content",
			CreatedAt: default_time,
		},
	}
}

func prepareRoomDetailEntity(id int) *entity.Room {
	room := new(entity.Room)
	room.ID = id
	room.R = room.R.NewStruct()
	room.R.RoomUsers = entity.RoomUserSlice{
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
	room.R.RoomUsers[0].R = room.R.RoomUsers[0].R.NewStruct()
	room.R.RoomUsers[1].R = room.R.RoomUsers[0].R.NewStruct()
	room.R.RoomUsers[0].R.User = &entity.User{
		ID:       2,
		Name:     "name2",
		Icon:     "icon2",
		Gender:   1,
		Birthday: default_time,
		Location: 1,
	}
	room.R.RoomUsers[1].R.User = &entity.User{
		ID:       1,
		Name:     "name1",
		Icon:     "icon1",
		Gender:   0,
		Birthday: default_time,
		Location: 0,
	}
	room.R.Messages = entity.MessageSlice{
		{
			ID:        1,
			UserID:    1,
			Content:   "content",
			CreatedAt: default_time,
		},
	}
	return room
}

func prepareRoomDetail(id int) *model.RoomDetail {
	return &model.RoomDetail{
		ID:    model.RoomID(id),
		Name:  "name2",
		Icon:  "icon2",
		Users: []*model.User{prepareUser(2, 1, 1), prepareUser(1, 0, 0)},
		Messages: []*model.Message{
			{
				ID:        1,
				UserID:    1,
				Content:   "content",
				CreatedAt: default_time,
			},
		},
	}
}

func prepareMessage(id int) *model.Message {
	return &model.Message{
		ID:        model.MessageID(id),
		UserID:    userID,
		Content:   "content",
		CreatedAt: default_time,
	}
}

func prepareNewMessage() *model.NewMessage {
	return &model.NewMessage{
		Content: "content",
	}
}

func preparePostMessageEntity() *entity.Message {
	return &entity.Message{
		UserID:  userID,
		RoomID:  roomID,
		Content: "content",
	}
}

func prepareCreatedMessageEntity(id int) *entity.Message {
	return &entity.Message{
		ID:        int64(id),
		UserID:    userID,
		RoomID:    roomID,
		Content:   "content",
		CreatedAt: default_time,
	}
}

// ----- END デフォルトのテストデータ -----

type UserUsecaseTestSuite struct {
	suite.Suite
	mock    *mock.MockIUserService
	usecase IUserUsecase
}

func (suite *UserUsecaseTestSuite) SetupSuite() {
	mockCtl := gomock.NewController(suite.T())
	defer mockCtl.Finish()
	suite.mock = mock.NewMockIUserService(mockCtl)
	suite.usecase = NewUserUsecase(suite.mock)
}

func TestUserHandlerSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseTestSuite))
}

func (suite *UserUsecaseTestSuite) Test_userUsecase_FindUserByUserID() {
	suite.mock.EXPECT().FindUserByUserID(&gin.Context{}, 1).Return(prepareUserEntity(1, 0, 0), nil)
	res, err := suite.usecase.FindUserByUserID(&gin.Context{}, 1)
	suite.Equal(err, nil)
	suite.Equal(res, prepareUser(1, 0, 0))
}

func (suite *UserUsecaseTestSuite) Test_userUsecase_FindAllUsers() {
	suite.mock.EXPECT().FindAllUsers(&gin.Context{}).Return(
		entity.UserSlice{prepareUserEntity(1, 0, 0), prepareUserEntity(2, 1, 1)},
		nil,
	)
	res, err := suite.usecase.FindAllUsers(&gin.Context{})
	suite.Equal(err, nil)
	suite.Equal(
		res,
		model.UserSlice{prepareUser(1, 0, 0), prepareUser(2, 1, 1)},
	)
}

func (suite *UserUsecaseTestSuite) Test_userUsecase_FindAllRooms() {
	suite.mock.EXPECT().FindAllRooms(&gin.Context{}, userID).Return(
		entity.RoomSlice{prepareRoomEntity(1)},
		nil,
	)
	res, err := suite.usecase.FindAllRooms(&gin.Context{}, userID)
	suite.Equal(err, nil)
	suite.Equal(
		res,
		&model.Rooms{Rooms: []*model.Room{prepareRoom(1)}},
	)
}

func (suite *UserUsecaseTestSuite) Test_userUsecase_FindRoomDetailByRoomID() {
	suite.mock.EXPECT().FindRoomDetailByRoomID(&gin.Context{}, 1, 1).Return(
		prepareRoomDetailEntity(1),
		nil,
	)
	res, err := suite.usecase.FindRoomDetailByRoomID(&gin.Context{}, 1, 1)
	suite.Equal(err, nil)
	suite.Equal(
		res,
		prepareRoomDetail(1),
	)
}

func (suite *UserUsecaseTestSuite) Test_userUsecase_SendMessage() {
	suite.mock.EXPECT().SendMessage(&gin.Context{}, preparePostMessageEntity()).Return(
		prepareCreatedMessageEntity(1),
		nil,
	)
	res, err := suite.usecase.SendMessage(&gin.Context{}, 1, 1, prepareNewMessage())
	suite.Equal(err, nil)
	suite.Equal(
		res,
		prepareMessage(1),
	)
}
