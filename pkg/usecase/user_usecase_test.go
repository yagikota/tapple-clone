package usecase

import (
	"strconv"
	"testing"
	"time"

	dmodel "github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/model"
	mock "github.com/CyberAgentHack/2208-ace-go-server/pkg/mock/service"
	"github.com/CyberAgentHack/2208-ace-go-server/pkg/usecase/model"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

// ----- BEGIN デフォルトのテストデータ -----
var (
	defaultTime = time.Date(2022, 4, 1, 0, 0, 0, 0, time.UTC)
	userID      = 1
	roomID      = 1
	messageID   = 1
)

func prepareUserDomainModel(id, gender, location int) *dmodel.User {
	return &dmodel.User{
		ID:       id,
		Name:     "name" + strconv.Itoa(id),
		Icon:     "icon" + strconv.Itoa(id),
		Gender:   gender,
		Birthday: defaultTime,
		Location: location,
	}
}

func prepareUser(id, gender, location int) *model.User {
	return &model.User{
		ID:       model.UserID(id),
		Name:     "name" + strconv.Itoa(id),
		Icon:     "icon" + strconv.Itoa(id),
		Gender:   gender,
		BirthDay: defaultTime,
		Location: location,
	}
}

func prepareRoomDomainModel(id int) *dmodel.Room {
	room := new(dmodel.Room)
	room.ID = id
	room.R = room.R.NewStruct()
	room.R.Messages = dmodel.MessageSlice{
		{
			ID:        1,
			UserID:    1,
			Content:   "content",
			CreatedAt: defaultTime,
		},
	}
	room.R.RoomUsers = dmodel.RoomUserSlice{
		// 自分自身
		{
			ID:       1,
			UserID:   1,
			RoomID:   1,
			IsPinned: false,
		},
		// 相手側のユーザー
		{
			ID:       2,
			UserID:   2,
			RoomID:   1,
			IsPinned: false,
		},
	}
	room.R.RoomUsers[0].R = room.R.RoomUsers[0].R.NewStruct()
	// 相手側のユーザー
	room.R.RoomUsers[0].R.User = &dmodel.User{
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
		// 相手側のユーザーの名前とアイコン
		Name: "name2",
		Icon: "icon2",
		LatestMessage: &model.Message{
			ID:        1,
			UserID:    1,
			Content:   "content",
			CreatedAt: defaultTime,
		},
	}
}

func prepareRoomDetailDomainModel(id int) *dmodel.Room {
	room := new(dmodel.Room)
	room.ID = id
	room.R = room.R.NewStruct()
	room.R.RoomUsers = dmodel.RoomUserSlice{
		// 相手側のユーザー
		{
			ID:       1,
			UserID:   1,
			RoomID:   1,
			IsPinned: false,
		},
		// 相手側のユーザー
		{
			ID:       2,
			UserID:   2,
			RoomID:   1,
			IsPinned: false,
		},
	}
	room.R.RoomUsers[0].R = room.R.RoomUsers[0].R.NewStruct()
	room.R.RoomUsers[1].R = room.R.RoomUsers[0].R.NewStruct()
	room.R.RoomUsers[0].R.User = &dmodel.User{
		// 相手側のユーザー
		ID:       2,
		Name:     "name2",
		Icon:     "icon2",
		Gender:   1,
		Birthday: defaultTime,
		Location: 1,
	}
	room.R.RoomUsers[1].R.User = &dmodel.User{
		// 自分自身
		ID:       1,
		Name:     "name1",
		Icon:     "icon1",
		Gender:   0,
		Birthday: defaultTime,
		Location: 0,
	}
	room.R.Messages = dmodel.MessageSlice{
		{
			ID:        1,
			UserID:    1,
			Content:   "content",
			CreatedAt: defaultTime,
		},
	}
	return room
}

func prepareRoomDetail(id int) *model.RoomDetail {
	return &model.RoomDetail{
		ID: model.RoomID(id),
		// 相手側のユーザーの名前とアイコン
		Name:  "name2",
		Icon:  "icon2",
		Users: []*model.User{prepareUser(2, 1, 1), prepareUser(1, 0, 0)},
		Messages: []*model.Message{
			{
				ID:        1,
				UserID:    1,
				Content:   "content",
				CreatedAt: defaultTime,
			},
		},
	}
}

func prepareMessage(id int) *model.Message {
	return &model.Message{
		ID:        model.MessageID(id),
		UserID:    userID,
		Content:   "content",
		CreatedAt: defaultTime,
	}
}

func prepareNewMessage() *model.NewMessage {
	return &model.NewMessage{
		Content: "content",
	}
}

func preparePostMessageDomainModel() *dmodel.Message {
	return &dmodel.Message{
		UserID:  userID,
		RoomID:  roomID,
		Content: "content",
	}
}

func prepareCreatedMessageDomainModel(id int) *dmodel.Message {
	return &dmodel.Message{
		ID:        int64(id),
		UserID:    userID,
		RoomID:    roomID,
		Content:   "content",
		CreatedAt: defaultTime,
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
	suite.mock.EXPECT().FindUserByUserID(&gin.Context{}, 1).Return(prepareUserDomainModel(1, 0, 0), nil)
	res, err := suite.usecase.FindUserByUserID(&gin.Context{}, 1)
	suite.Equal(err, nil)
	suite.Equal(res, prepareUser(1, 0, 0))
}

func (suite *UserUsecaseTestSuite) Test_userUsecase_FindAllUsers() {
	suite.mock.EXPECT().FindAllUsers(&gin.Context{}).Return(
		dmodel.UserSlice{prepareUserDomainModel(1, 0, 0), prepareUserDomainModel(2, 1, 1)},
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
		dmodel.RoomSlice{prepareRoomDomainModel(1)},
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
	suite.mock.EXPECT().FindRoomDetailByRoomID(&gin.Context{}, userID, roomID, messageID).Return(
		prepareRoomDetailDomainModel(1),
		nil,
	)
	res, err := suite.usecase.FindRoomDetailByRoomID(&gin.Context{}, userID, roomID, messageID)
	suite.Equal(err, nil)
	suite.Equal(
		res,
		prepareRoomDetail(1),
	)
}

func (suite *UserUsecaseTestSuite) Test_userUsecase_SendMessage() {
	suite.mock.EXPECT().SendMessage(&gin.Context{}, preparePostMessageDomainModel()).Return(
		prepareCreatedMessageDomainModel(1),
		nil,
	)
	res, err := suite.usecase.SendMessage(&gin.Context{}, userID, roomID, prepareNewMessage())
	suite.Equal(err, nil)
	suite.Equal(
		res,
		prepareMessage(1),
	)
}
