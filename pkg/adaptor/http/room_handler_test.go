package http

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	mock "github.com/CyberAgentHack/2208-ace-go-server/pkg/mock/usecase"
	"github.com/CyberAgentHack/2208-ace-go-server/pkg/usecase/model"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type RoomHandlerTestSuite struct {
	suite.Suite
	router  *gin.Engine
	mock    *mock.MockIRoomUsecase
	handler *roomHandler
	rec     *httptest.ResponseRecorder
}

// https://stackoverflow.com/questions/50200933/difference-between-setupsuite-setuptest-in-testify-suites
func (suite *RoomHandlerTestSuite) SetupSuite() {
	gin.SetMode(gin.TestMode)

	suite.router = gin.Default()
	mockCtl := gomock.NewController(suite.T())
	defer mockCtl.Finish()
	suite.mock = mock.NewMockIRoomUsecase(mockCtl)
	suite.handler = NewRoomHandler(suite.mock)

	suite.router.Use(
		func(ctx *gin.Context) {
			ctx.Set("user_id", "1")
			ctx.Set("room_id", "1")
			ctx.Set("message_id", "1")
		},
		checkStatusMiddleware(),
	)

	// v1/users/{user_id}/rooms
	path := usersAPIRoot + fmt.Sprintf("/:%s/rooms", userIDParam)
	suite.router.GET(path, suite.handler.findRooms())
	// v1/users/{user_id}/rooms/{room_id}
	path = usersAPIRoot + fmt.Sprintf("/:%s/rooms/:%s", userIDParam, roomIDParam)
	suite.router.GET(path, suite.handler.findRoomDetailByRoomID())
	// v1/users/{user_id}/rooms/{room_id}/messages
	path = usersAPIRoot + fmt.Sprintf("/:%s/rooms/:%s/messages", userIDParam, roomIDParam)
	suite.router.POST(path, suite.handler.sendMessage())
}

func (suite *RoomHandlerTestSuite) SetupTest() {
	suite.rec = httptest.NewRecorder()
	userID = 1
	user1 = &model.User{
		ID:          model.UserID(userID),
		Name:        "name1",
		Icon:        "/icon1",
		Gender:      1,
		BirthDay:    time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		Location:    "その他",
		IsPrincipal: false,
	}
	userSlice1 = model.UserSlice{user1}

	message1 = &model.Message{
		ID:        1,
		UserID:    1,
		Content:   "content1",
		CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	messageSlice1 = model.MessageSlice{message1}

	roomID = 1
	room1 = &model.Room{
		ID:            model.RoomID(roomID),
		Unread:        1,
		IsPinned:      true,
		Name:          "name1",
		SubName:       "sub_name1",
		Icon:          "/icon1",
		LatestMessage: message1,
	}
	roomSlice = model.RoomSlice{room1}
	rooms1 = &model.Rooms{Rooms: roomSlice}
	roomDetail = &model.RoomDetail{
		ID:       model.RoomID(roomID),
		Name:     "name1",
		Icon:     "/icon1",
		Users:    userSlice1,
		Messages: messageSlice1,
		IsLast:   true,
	}

	newMessage1 = &model.NewMessage{
		Content: "content1",
	}

	messageID = 1
}

// テストを実行するのに必要
func TestRoomHandlerSuite(t *testing.T) {
	suite.Run(t, new(RoomHandlerTestSuite))
}

func (suite *RoomHandlerTestSuite) Test_roomHandler_findRooms_200() {
	suite.mock.EXPECT().FindAllRooms(gomock.Any(), userID).Return(rooms1, nil)
	rec := suite.rec
	path := fmt.Sprintf("%s/%d/rooms", usersAPIRoot, userID)
	req := httptest.NewRequest(http.MethodGet, path, nil)
	suite.router.ServeHTTP(rec, req)
	suite.Equal(http.StatusOK, rec.Code)
	suite.JSONEq(
		`{
			"rooms": [
				{
					"id": 1,
					"unread": 1,
					"is_pinned": true,
					"name": "name1",
					"sub_name": "sub_name1",
					"is_principal": false,
					"icon": "/icon1",
					"latest_message": {
						"id": 1,
						"user_id": 1,
						"content": "content1",
						"created_at": "2022-01-01T00:00:00Z"
					}
				}
			]
		}`,
		rec.Body.String(),
	)
}

func (suite *RoomHandlerTestSuite) Test_roomHandler_findRooms_400() {
	rec := suite.rec
	path := fmt.Sprintf("%s/dummy_user_id/rooms", usersAPIRoot)
	req := httptest.NewRequest(http.MethodGet, path, nil)
	suite.router.ServeHTTP(rec, req)
	suite.Equal(http.StatusBadRequest, rec.Code)
	suite.JSONEq(
		`{
			"message": "Bad Request"
		}`,
		rec.Body.String(),
	)
}

func (suite *RoomHandlerTestSuite) Test_roomHandler_findRooms_500() {
	suite.mock.EXPECT().FindAllRooms(gomock.Any(), userID).Return(nil, errors.New("dummy_error"))
	rec := suite.rec
	path := fmt.Sprintf("%s/%d/rooms", usersAPIRoot, userID)
	req := httptest.NewRequest(http.MethodGet, path, nil)
	suite.router.ServeHTTP(rec, req)
	suite.Equal(http.StatusInternalServerError, rec.Code)
	suite.JSONEq(
		`{
			"message":"Internal Server Error"
		}`,
		rec.Body.String(),
	)
}

func (suite *RoomHandlerTestSuite) Test_roomHandler_findRoomDetailByRoomID_200() {
	suite.mock.EXPECT().FindRoomDetailByRoomID(gomock.Any(), userID, roomID, messageID).Return(roomDetail, nil)
	rec := suite.rec
	path := fmt.Sprintf("%s/%d/rooms/%d?message_id=1", usersAPIRoot, userID, roomID)
	req := httptest.NewRequest(http.MethodGet, path, nil)
	suite.router.ServeHTTP(rec, req)
	suite.Equal(http.StatusOK, rec.Code)
	suite.JSONEq(
		`{
			"id": 1,
			"name": "name1",
			"icon": "/icon1",
			"users": [
				{
					"id": 1,
					"name": "name1",
					"icon": "/icon1",
					"gender": 1,
					"birthday": "2022-01-01T00:00:00Z",
					"location": "その他",
					"is_principal": false
				}
			],
			"messages": [
				{
					"id": 1,
					"user_id": 1,
					"content": "content1",
					"created_at": "2022-01-01T00:00:00Z"
				}
			],
			"is_last": true
		}`,
		rec.Body.String(),
	)
}

func (suite *RoomHandlerTestSuite) Test_roomHandler_findRoomDetailByRoomID_userID_400() {
	rec := suite.rec
	path := fmt.Sprintf("%s/dummy_user_id/rooms/%d?message_id=1", usersAPIRoot, roomID)
	req := httptest.NewRequest(http.MethodGet, path, nil)
	suite.router.ServeHTTP(rec, req)
	suite.Equal(http.StatusBadRequest, rec.Code)
	suite.JSONEq(
		`{
			"message": "Bad Request"
		}`,
		rec.Body.String(),
	)
}

func (suite *RoomHandlerTestSuite) Test_roomHandler_findRoomDetailByRoomID_roomID_400() {
	rec := suite.rec
	path := fmt.Sprintf("%s/%d/rooms/dummy_room_id?message_id=1", usersAPIRoot, userID)
	req := httptest.NewRequest(http.MethodGet, path, nil)
	suite.router.ServeHTTP(rec, req)
	suite.Equal(http.StatusBadRequest, rec.Code)
	suite.JSONEq(
		`{
			"message": "Bad Request"
		}`,
		rec.Body.String(),
	)
}

func (suite *RoomHandlerTestSuite) Test_roomHandler_findRoomDetailByRoomID_messageID_400() {
	rec := suite.rec
	path := fmt.Sprintf("%s/%d/rooms/%d?message_id=dummy", usersAPIRoot, userID, roomID)
	req := httptest.NewRequest(http.MethodGet, path, nil)
	suite.router.ServeHTTP(rec, req)
	suite.Equal(http.StatusBadRequest, rec.Code)
	suite.JSONEq(
		`{
			"message": "Bad Request"
		}`,
		rec.Body.String(),
	)
}

func (suite *RoomHandlerTestSuite) Test_roomHandler_findRoomDetailByRoomID_500() {
	suite.mock.EXPECT().FindRoomDetailByRoomID(gomock.Any(), userID, roomID, messageID).Return(nil, errors.New("dummy_error"))
	rec := suite.rec
	path := fmt.Sprintf("%s/%d/rooms/%d?message_id=1", usersAPIRoot, userID, roomID)
	req := httptest.NewRequest(http.MethodGet, path, nil)
	suite.router.ServeHTTP(rec, req)
	suite.Equal(http.StatusInternalServerError, rec.Code)
	suite.JSONEq(
		`{
			"message":"Internal Server Error"
		}`,
		rec.Body.String(),
	)
}

func (suite *RoomHandlerTestSuite) Test_roomHandler_sendMessage_200() {
	suite.mock.EXPECT().SendMessage(gomock.Any(), userID, roomID, newMessage1).Return(message1, nil)
	rec := suite.rec
	path := fmt.Sprintf("%s/%d/rooms/%d/messages", usersAPIRoot, userID, roomID)
	jsonStr := []byte(
		`{
			"content": "content1"
		}`,
	)
	req := httptest.NewRequest(http.MethodPost, path, bytes.NewBuffer(jsonStr))
	suite.router.ServeHTTP(rec, req)
	suite.Equal(http.StatusOK, rec.Code)
	suite.JSONEq(
		`{
			"id": 1,
			"user_id": 1,
			"content": "content1",
			"created_at": "2022-01-01T00:00:00Z"
		}`,
		rec.Body.String(),
	)
}

func (suite *RoomHandlerTestSuite) Test_roomHandler_sendMessage_userID_400() {
	rec := suite.rec
	path := fmt.Sprintf("%s/dummy_user_id/rooms/%d/messages", usersAPIRoot, roomID)
	jsonStr := []byte(
		`{
			"content": "content1"
		}`,
	)
	req := httptest.NewRequest(http.MethodPost, path, bytes.NewBuffer(jsonStr))
	suite.router.ServeHTTP(rec, req)
	suite.Equal(http.StatusBadRequest, rec.Code)
	suite.JSONEq(
		`{
			"message": "Bad Request"
		}`,
		rec.Body.String(),
	)
}

func (suite *RoomHandlerTestSuite) Test_roomHandler_sendMessage_roomID_400() {
	rec := suite.rec
	path := fmt.Sprintf("%s/%d/rooms/dummy_room_id/messages", usersAPIRoot, userID)
	jsonStr := []byte(
		`{
			"content": "content1"
		}`,
	)
	req := httptest.NewRequest(http.MethodPost, path, bytes.NewBuffer(jsonStr))
	suite.router.ServeHTTP(rec, req)
	suite.Equal(http.StatusBadRequest, rec.Code)
	suite.JSONEq(
		`{
			"message": "Bad Request"
		}`,
		rec.Body.String(),
	)
}

func (suite *RoomHandlerTestSuite) Test_roomHandler_sendMessage_requestBody_400() {
	rec := suite.rec
	path := fmt.Sprintf("%s/%d/rooms/%d/messages", usersAPIRoot, userID, roomID)
	jsonStr := []byte(
		`{
			"content_dummy": "content1"
		}`,
	)
	req := httptest.NewRequest(http.MethodPost, path, bytes.NewBuffer(jsonStr))
	suite.router.ServeHTTP(rec, req)
	suite.Equal(http.StatusBadRequest, rec.Code)
	suite.JSONEq(
		`{
			"message": "Bad Request"
		}`,
		rec.Body.String(),
	)
}

func (suite *RoomHandlerTestSuite) Test_roomHandler_sendMessage_500() {
	suite.mock.EXPECT().SendMessage(gomock.Any(), userID, roomID, newMessage1).Return(nil, errors.New("dummy_error"))
	rec := suite.rec
	path := fmt.Sprintf("%s/%d/rooms/%d/messages", usersAPIRoot, userID, roomID)
	jsonStr := []byte(
		`{
			"content": "content1"
		}`,
	)
	req := httptest.NewRequest(http.MethodPost, path, bytes.NewBuffer(jsonStr))
	suite.router.ServeHTTP(rec, req)
	suite.Equal(http.StatusInternalServerError, rec.Code)
	suite.JSONEq(
		`{
			"message":"Internal Server Error"
		}`,
		rec.Body.String(),
	)
}
