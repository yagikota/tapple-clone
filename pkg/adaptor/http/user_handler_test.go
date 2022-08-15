package http

import (
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

type UserHandlerTestSuite struct {
	suite.Suite
	router  *gin.Engine
	mock    *mock.MockIUserUsecase
	handler *userHandler
}

// https://stackoverflow.com/questions/50200933/difference-between-setupsuite-setuptest-in-testify-suites
// The SetupSuite method will be run by testify once, at the very
// start of the testing suite, before any tests are run.
func (suite *UserHandlerTestSuite) SetupSuite() {
	gin.SetMode(gin.TestMode)

	suite.router = gin.Default()
	mockCtl := gomock.NewController(suite.T())
	defer mockCtl.Finish()
	suite.mock = mock.NewMockIUserUsecase(mockCtl)
	suite.handler = NewUserHandler(suite.mock)

	// ハンドラー登録(mockが登録させたハンドラー)
	suite.router.GET(usersAPIRoot+"/:user_id", func(c *gin.Context) { c.Set("user_id", 1) }, suite.handler.findUserByUserID())
}

// テストを実行するのに必要
func TestUserHandlerSuite(t *testing.T) {
	suite.Run(t, new(UserHandlerTestSuite))
}

func (suite *UserHandlerTestSuite) Test_userHandler_findUserByUserID_200() {
	suite.mock.EXPECT().FindUserByUserID(gomock.Any(), 1).Return(
		&model.User{
			ID:       1,
			Name:     "name1",
			Icon:     "/icon1",
			Gender:   1,
			BirthDay: time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
			Location: 1,
		},
		nil,
	)

	// レスポンスを受け止める*httptest.ResponseRecorder
	rec := httptest.NewRecorder()
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
			"birthday":"2022-01-01T00:00:00+09:00",
			"location":1
		}`,
		rec.Body.String(),
	)
}

// func Test_userHandler_findUsers(t *testing.T) {
// 	type fields struct {
// 		uUsecase usecase.IUserUsecase
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		want   gin.HandlerFunc
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			uh := &userHandler{
// 				uUsecase: tt.fields.uUsecase,
// 			}
// 			if got := uh.findUsers(); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("userHandler.findUsers() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_userHandler_findRooms(t *testing.T) {
// 	type fields struct {
// 		uUsecase usecase.IUserUsecase
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		// buildStubs func(store *mock.Mo)
// 		want gin.HandlerFunc
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			uh := &userHandler{
// 				uUsecase: tt.fields.uUsecase,
// 			}
// 			if got := uh.findRooms(); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("userHandler.findRooms() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_userHandler_findRoomDetailByRoomID(t *testing.T) {
// 	type fields struct {
// 		uUsecase usecase.IUserUsecase
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		want   gin.HandlerFunc
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			uh := &userHandler{
// 				uUsecase: tt.fields.uUsecase,
// 			}
// 			if got := uh.findRoomDetailByRoomID(); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("userHandler.findRoomDetailByRoomID() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_userHandler_sendMessage(t *testing.T) {
// 	type fields struct {
// 		uUsecase usecase.IUserUsecase
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		want   gin.HandlerFunc
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			uh := &userHandler{
// 				uUsecase: tt.fields.uUsecase,
// 			}
// 			if got := uh.sendMessage(); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("userHandler.sendMessage() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
