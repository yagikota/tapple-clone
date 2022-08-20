package http

import (
	"bytes"
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

// テストデータ
var (
	userID     int
	user1      *model.User
	userSlice1 model.UserSlice

	message1      *model.Message
	messageSlice1 model.MessageSlice

	roomID     int
	room1      *model.Room
	roomSlice  model.RoomSlice
	rooms1     *model.Rooms
	roomDetail *model.RoomDetail

	newMessage1 *model.NewMessage

	messageID int
)

//  1.SetupSuite
//	2. 各テスト
//		2.1 SetupTest
//		2.2 BeforeTest
//		2.3 Test(実行されるテスト)
//		2.4 AfterTest
//		2.5 TeardownTest
//	3.TearDownSuite

type UserHandlerTestSuite struct {
	suite.Suite
	router  *gin.Engine
	mock    *mock.MockIUserUsecase
	handler *userHandler
	rec     *httptest.ResponseRecorder
}

// https://stackoverflow.com/questions/50200933/difference-between-setupsuite-setuptest-in-testify-suites
func (suite *UserHandlerTestSuite) SetupSuite() {
	gin.SetMode(gin.TestMode)

	suite.router = gin.Default()
	mockCtl := gomock.NewController(suite.T())
	defer mockCtl.Finish()
	suite.mock = mock.NewMockIUserUsecase(mockCtl)
	suite.handler = NewUserHandler(suite.mock)

	suite.router.Use(
		func(ctx *gin.Context) {
			ctx.Set("user_id", "1")
			ctx.Set("room_id", "1")
			ctx.Set("message_id", "1")
		},
	)
	// ハンドラー登録
	// v1/users
	suite.router.GET(usersAPIRoot, suite.handler.findUsers())
	// v1/users/{user_id}
	path := usersAPIRoot + fmt.Sprintf("/:%s", userIDParam)
	suite.router.GET(path, suite.handler.findUserByUserID())
	// v1/users/{user_id}/rooms
	path = usersAPIRoot + fmt.Sprintf("/:%s/rooms", userIDParam)
	suite.router.GET(path, suite.handler.findRooms())
	// v1/users/{user_id}/rooms/{room_id}
	path = usersAPIRoot + fmt.Sprintf("/:%s/rooms/:%s", userIDParam, roomIDParam)
	suite.router.GET(path, suite.handler.findRoomDetailByRoomID())
	// v1/users/{user_id}/rooms/{room_id}/messages
	path = usersAPIRoot + fmt.Sprintf("/:%s/rooms/:%s/messages", userIDParam, roomIDParam)
	suite.router.POST(path, suite.handler.sendMessage())
}

func (suite *UserHandlerTestSuite) SetupTest() {
	suite.rec = httptest.NewRecorder()
	userID = 1
	user1 = &model.User{
		ID:       model.UserID(userID),
		Name:     "name1",
		Icon:     "/icon1",
		Gender:   1,
		BirthDay: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		Location: 1,
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
func TestUserHandlerSuite(t *testing.T) {
	suite.Run(t, new(UserHandlerTestSuite))
}

func (suite *UserHandlerTestSuite) Test_userHandler_findUserByUserID_200() {
	// TODO: gomock.Any()を貼るべく使わないようにする https://stackoverflow.com/questions/66952761/how-to-unit-test-a-go-gin-handler-function
	suite.mock.EXPECT().FindUserByUserID(gomock.Any(), userID).Return(user1, nil)
	// レスポンスを受け止める*httptest.ResponseRecorder
	rec := suite.rec
	// テストで送信するリクエスト
	req := httptest.NewRequest(http.MethodGet, usersAPIRoot+"/1", nil)
	// リクエスト送信
	suite.router.ServeHTTP(rec, req)
	// 結果確認
	suite.Equal(http.StatusOK, rec.Code)
	suite.JSONEq(
		`{
			"id":1,
			"name":"name1",
			"icon":"/icon1",
			"gender":1,
			"birthday":"2022-01-01T00:00:00Z",
			"location":1
		}`,
		rec.Body.String(),
	)
}

func (suite *UserHandlerTestSuite) Test_userHandler_findUsers_200() {
	suite.mock.EXPECT().FindAllUsers(gomock.Any()).Return(userSlice1, nil)
	rec := suite.rec
	req := httptest.NewRequest(http.MethodGet, usersAPIRoot, nil)
	suite.router.ServeHTTP(rec, req)
	suite.Equal(http.StatusOK, rec.Code)
	suite.JSONEq(
		`[
			{
				"id":1,
				"name":"name1",
				"icon":"/icon1",
				"gender":1,
				"birthday":"2022-01-01T00:00:00Z",
				"location":1
			}
		]`,
		rec.Body.String(),
	)
}

func (suite *UserHandlerTestSuite) Test_userHandler_findRooms_200() {
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

func (suite *UserHandlerTestSuite) Test_userHandler_findRoomDetailByRoomID_200() {
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
					"location": 1
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

func (suite *UserHandlerTestSuite) Test_userHandler_sendMessage_200() {
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
