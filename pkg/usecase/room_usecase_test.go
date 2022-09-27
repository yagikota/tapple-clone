package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	dmodel "github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/model"
	mock "github.com/CyberAgentHack/2208-ace-go-server/pkg/mock/service"
	"github.com/CyberAgentHack/2208-ace-go-server/pkg/usecase/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

// ----- BEGIN デフォルトのテストデータ -----

func prepareRoomDomainModel(id int, ispinned bool, day int) *dmodel.Room {
	room := new(dmodel.Room)
	room.ID = id
	room.R = room.R.NewStruct()
	room.R.Messages = dmodel.MessageSlice{
		{
			ID:        1,
			UserID:    1,
			Content:   "content",
			CreatedAt: defaultTime.AddDate(0, 0, day),
		},
	}
	room.R.Messages[0].R = room.R.Messages[0].R.NewStruct()
	room.R.Messages[0].R.User = &dmodel.User{
		ID:          1,
		Name:        "name1",
		Icon:        "icon1",
		Gender:      0,
		Birthday:    time.Date(2000, 5, 9, 23, 59, 59, 0, time.Local),
		Location:    0,
		IsPrincipal: false,
	}
	room.R.RoomUsers = dmodel.RoomUserSlice{
		// 自分自身
		{
			ID:       1,
			UserID:   1,
			RoomID:   1,
			IsPinned: ispinned,
		},
		// 相手側のユーザー
		{
			ID:       2,
			UserID:   2,
			RoomID:   1,
			IsPinned: ispinned,
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

func prepareRoom(id int, day int, ispinned bool) *model.Room {
	return &model.Room{
		ID:          model.RoomID(id),
		Unread:      0,
		IsPinned:    ispinned,
		IsPrincipal: true,
		// 相手側のユーザーの名前とアイコン
		Name:    "name2",
		Icon:    "icon2",
		SubName: "22歳・その他",
		LatestMessage: &model.Message{
			ID: 1,
			User: &model.User{
				ID:          1,
				Name:        "name1",
				Icon:        "icon1",
				Gender:      0,
				BirthDay:    time.Date(2000, 5, 9, 23, 59, 59, 0, time.Local),
				Location:    "その他",
				IsPrincipal: false,
			},
			Content:   "content",
			CreatedAt: defaultTime.AddDate(0, 0, day),
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

	room.R.Messages[0].R = room.R.Messages[0].R.NewStruct()
	room.R.Messages[0].R.User = &dmodel.User{
		ID:          1,
		Name:        "name1",
		Icon:        "icon1",
		Gender:      0,
		Birthday:    time.Date(2000, 5, 9, 23, 59, 59, 0, time.Local),
		Location:    0,
		IsPrincipal: false,
	}
	return room
}

func prepareRoomDetail(id int) *model.RoomDetail {
	return &model.RoomDetail{
		ID: model.RoomID(id),
		// 相手側のユーザーの名前とアイコン
		Name: "name2",
		Icon: "icon2",
		// Users: []*model.User{prepareUser(2, 1, "北海道"), prepareUser(1, 0, "その他")},
		Messages: []*model.Message{
			{
				ID: 1,
				User: &model.User{
					ID:          1,
					Name:        "name1",
					Icon:        "icon1",
					Gender:      0,
					BirthDay:    time.Date(2000, 5, 9, 23, 59, 59, 0, time.Local),
					Location:    "その他",
					IsPrincipal: false,
				},
				Content:   "content",
				CreatedAt: defaultTime,
			},
		},
		IsLast: true,
	}
}

func prepareMessage(id int) *model.Message {
	return &model.Message{
		ID: model.MessageID(id),
		User: &model.User{
			ID:          1,
			Name:        "name1",
			Icon:        "icon1",
			Gender:      0,
			BirthDay:    time.Date(2000, 5, 9, 23, 59, 59, 0, time.Local),
			Location:    "その他",
			IsPrincipal: false,
		},
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
	message := new(dmodel.Message)
	message = &dmodel.Message{
		ID:        int64(id),
		UserID:    userID,
		RoomID:    roomID,
		Content:   "content",
		CreatedAt: defaultTime,
	}

	message.R = message.R.NewStruct()
	message.R.User = &dmodel.User{
		ID:          1,
		Name:        "name1",
		Icon:        "icon1",
		Gender:      0,
		Birthday:    time.Date(2000, 5, 9, 23, 59, 59, 0, time.Local),
		Location:    0,
		IsPrincipal: false,
	}

	return message

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
		dmodel.RoomSlice{
			prepareRoomDomainModel(1, false, 1),
			prepareRoomDomainModel(2, false, -1),
			prepareRoomDomainModel(3, true, 4),
			prepareRoomDomainModel(4, true, 1),
			prepareRoomDomainModel(5, false, 5),
		},
		nil,
	)
	res, err := suite.usecase.FindAllRooms(context.Background(), userID)
	suite.Equal(err, nil)
	suite.Equal(
		res,
		&model.Rooms{
			Rooms: []*model.Room{
				prepareRoom(3, 4, true),
				prepareRoom(4, 1, true),
				prepareRoom(5, 5, false),
				prepareRoom(1, 1, false),
				prepareRoom(2, -1, false),
			}},
	)
}

func (suite *RoomUsecaseTestSuite) Test_roomUsecase_Err_FindAllRooms() {
	suite.mock.EXPECT().FindAllRooms(context.Background(), userID).Return(nil, errors.New("could not find rooms"))
	res, err := suite.usecase.FindAllRooms(context.Background(), userID)
	suite.Nil(res)
	suite.Equal(err, errors.New("could not find rooms"))
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

func (suite *RoomUsecaseTestSuite) Test_roomUsecase_Err_FindRoomDetailByRoomID() {
	suite.mock.EXPECT().FindRoomDetailByRoomID(context.Background(), userID, roomID, messageID).Return(nil, errors.New("could not find room detail"))
	res, err := suite.usecase.FindRoomDetailByRoomID(context.Background(), userID, roomID, messageID)
	suite.Nil(res)
	suite.Equal(err, errors.New("could not find room detail"))
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

func (suite *RoomUsecaseTestSuite) Test_roomUsecase_Err_SendMessage() {
	suite.mock.EXPECT().SendMessage(context.Background(), preparePostMessageDomainModel()).Return(nil, errors.New("could not send message"))
	res, err := suite.usecase.SendMessage(context.Background(), userID, roomID, prepareNewMessage())
	suite.Nil(res)
	suite.Equal(err, errors.New("could not send message"))
}
