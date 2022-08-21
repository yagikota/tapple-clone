package http

import (
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

	userDetail *model.UserDetail
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
		},
		checkStatusMiddleware(),
	)
	// ハンドラー登録
	// v1/users
	suite.router.GET(usersAPIRoot, suite.handler.findUsers())
	// v1/users/{user_id}
	path := usersAPIRoot + fmt.Sprintf("/:%s", userIDParam)
	suite.router.GET(path, suite.handler.findUserByUserID())
	// v1/users/{user_id}/profile
	path = usersAPIRoot + fmt.Sprintf("/:%s/profile", userIDParam)
	suite.router.GET(path, suite.handler.findUserDetailByUserID())
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
		Location: "その他",
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

	userDetail = &model.UserDetail{
		ID:          1,
		Name:        "name1",
		Age:         1,
		Location:    "その他",
		IsPrincipal: false,
		TagCount:    1,
		ProfileImages: []*model.UserProfileImage{
			{
				ID:        1,
				UserID:    1,
				ImagePath: "image_path",
			},
		},
		Hobbies: []*model.Hobby{
			{
				ID:  1,
				Tag: "tag",
			},
		},
	}
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
			"location":"その他",
			"is_principal": false
		}`,
		rec.Body.String(),
	)
}

func (suite *UserHandlerTestSuite) Test_userHandler_findUserByUserID_400() {
	rec := suite.rec
	req := httptest.NewRequest(http.MethodGet, usersAPIRoot+"/dummy", nil)
	suite.router.ServeHTTP(rec, req)
	suite.Equal(http.StatusBadRequest, rec.Code)
	suite.JSONEq(
		`{
			"message": "Bad Request"
		}`,
		rec.Body.String(),
	)
}

func (suite *UserHandlerTestSuite) Test_userHandler_findUserByUserID_500() {
	suite.mock.EXPECT().FindUserByUserID(gomock.Any(), userID).Return(nil, errors.New("dummy_error")).Times(1)
	rec := suite.rec
	req := httptest.NewRequest(http.MethodGet, usersAPIRoot+"/1", nil)
	suite.router.ServeHTTP(rec, req)
	suite.Equal(http.StatusInternalServerError, rec.Code)
	suite.JSONEq(
		`{
			"message":"Internal Server Error"
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
				"location":"その他",
				"is_principal": false
			}
		]`,
		rec.Body.String(),
	)
}

func (suite *UserHandlerTestSuite) Test_userHandler_findUsers_500() {
	suite.mock.EXPECT().FindAllUsers(gomock.Any()).Return(nil, errors.New("dummy_error"))
	rec := suite.rec
	req := httptest.NewRequest(http.MethodGet, usersAPIRoot, nil)
	suite.router.ServeHTTP(rec, req)
	suite.Equal(http.StatusInternalServerError, rec.Code)
	suite.JSONEq(
		`{
			"message":"Internal Server Error"
		}`,
		rec.Body.String(),
	)
}

func (suite *UserHandlerTestSuite) Test_userHandler_findUserDetailByUserID_200() {
	suite.mock.EXPECT().FindUserDetailByUserID(gomock.Any(), userID).Return(userDetail, nil)
	rec := suite.rec
	req := httptest.NewRequest(http.MethodGet, usersAPIRoot+"/1/profile", nil)
	suite.router.ServeHTTP(rec, req)
	suite.Equal(http.StatusOK, rec.Code)
	suite.JSONEq(
		`{
			"id": 1,
			"name": "name1",
			"age": 1,
			"location": "その他",
			"is_principal": false,
			"tag_count": 1,
			"profile_images": [
				{
					"id": 1,
					"user_id": 1,
					"image_path": "image_path"
				}
			],
			"hobbies": [
				{
					"id": 1,
					"tag": "tag"
				}
			]
		}`,
		rec.Body.String(),
	)
}

func (suite *UserHandlerTestSuite) Test_userHandler_findUserDetailByUserID_400() {
	rec := suite.rec
	req := httptest.NewRequest(http.MethodGet, usersAPIRoot+"/dummy/profile", nil)
	suite.router.ServeHTTP(rec, req)
	suite.Equal(http.StatusBadRequest, rec.Code)
	suite.JSONEq(
		`{
			"message": "Bad Request"
		}`,
		rec.Body.String(),
	)
}

func (suite *UserHandlerTestSuite) Test_userHandler_findUserDetailByUserID_500() {
	suite.mock.EXPECT().FindUserDetailByUserID(gomock.Any(), userID).Return(nil, errors.New("dummy_error"))
	rec := suite.rec
	req := httptest.NewRequest(http.MethodGet, usersAPIRoot+"/1/profile", nil)
	suite.router.ServeHTTP(rec, req)
	suite.Equal(http.StatusInternalServerError, rec.Code)
	suite.JSONEq(
		`{
			"message":"Internal Server Error"
		}`,
		rec.Body.String(),
	)
}
