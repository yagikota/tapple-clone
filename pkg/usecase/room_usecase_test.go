package usecase

import (
	"context"
	"testing"
	"time"

	dmodel "github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/model"
	mock "github.com/CyberAgentHack/2208-ace-go-server/pkg/mock/service"
	"github.com/CyberAgentHack/2208-ace-go-server/pkg/usecase/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

// ----- BEGIN デフォルトのテストデータ -----

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
		ID:          2,
		Name:        "name2",
		Icon:        "icon2",
		Birthday:    time.Date(2000, 5, 9, 23, 59, 59, 0, time.Local),
		IsPrincipal: true,
	}
	return room
}

func prepareRoom(id int) *model.Room {
	return &model.Room{
		ID:          model.RoomID(id),
		Unread:      0,
		IsPinned:    false,
		IsPrincipal: true,
		// 相手側のユーザーの名前とアイコン
		Name:    "name2",
		Icon:    "icon2",
		SubName: "22歳・その他",
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
		ID:          2,
		Name:        "name2",
		Icon:        "icon2",
		Gender:      1,
		Birthday:    defaultTime,
		Location:    1,
		IsPrincipal: true,
	}
	room.R.RoomUsers[1].R.User = &dmodel.User{
		// 自分自身
		ID:          1,
		Name:        "name1",
		Icon:        "icon1",
		Gender:      0,
		Birthday:    defaultTime,
		Location:    0,
		IsPrincipal: true,
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
		Users: []*model.User{prepareUser(2, 1, "北海道"), prepareUser(1, 0, "その他")},
		Messages: []*model.Message{
			{
				ID:        1,
				UserID:    1,
				Content:   "content",
				CreatedAt: defaultTime,
			},
		},
		IsLast: true,
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

type RoomUsecaseTestSuite struct {
	suite.Suite
	mock    *mock.MockIRoomService
	usecase IRoomUsecase
}

func (suite *RoomUsecaseTestSuite) SetupSuite() {
	mockCtl := gomock.NewController(suite.T())
	defer mockCtl.Finish()
	suite.mock = mock.NewMockIRoomService(mockCtl)
	suite.usecase = NewRoomUsecase(suite.mock)

	defaultTime = time.Date(2022, 4, 1, 0, 0, 0, 0, time.UTC)
	userID = 1
	roomID = 1
	messageID = 1
}

func TestRoomHandlerSuite(t *testing.T) {
	suite.Run(t, new(RoomUsecaseTestSuite))
}

func (suite *RoomUsecaseTestSuite) Test_roomUsecase_FindAllRooms() {
	suite.mock.EXPECT().FindAllRooms(context.Background(), userID).Return(
		dmodel.RoomSlice{prepareRoomDomainModel(1)},
		nil,
	)
	res, err := suite.usecase.FindAllRooms(context.Background(), userID)
	suite.Equal(err, nil)
	suite.Equal(
		res,
		&model.Rooms{Rooms: []*model.Room{prepareRoom(1)}},
	)
}

func (suite *RoomUsecaseTestSuite) Test_roomUsecase_FindRoomDetailByRoomID() {
	suite.mock.EXPECT().FindRoomDetailByRoomID(context.Background(), userID, roomID, messageID).Return(
		prepareRoomDetailDomainModel(1),
		nil,
	)
	res, err := suite.usecase.FindRoomDetailByRoomID(context.Background(), userID, roomID, messageID)
	suite.Equal(err, nil)
	suite.Equal(
		res,
		prepareRoomDetail(1),
	)
}

func (suite *RoomUsecaseTestSuite) Test_roomUsecase_SendMessage() {
	suite.mock.EXPECT().SendMessage(context.Background(), preparePostMessageDomainModel()).Return(
		prepareCreatedMessageDomainModel(1),
		nil,
	)
	res, err := suite.usecase.SendMessage(context.Background(), userID, roomID, prepareNewMessage())
	suite.Equal(err, nil)
	suite.Equal(
		res,
		prepareMessage(1),
	)
}
