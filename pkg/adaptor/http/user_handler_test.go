package http

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	mock "github.com/CyberAgentHack/2208-ace-go-server/pkg/mock/usecase"
	"github.com/CyberAgentHack/2208-ace-go-server/pkg/usecase/model"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// func TestNewUserHandler(t *testing.T) {
// 	type args struct {
// 		uu usecase.IUserUsecase
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want *userHandler
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := NewUserHandler(tt.args.uu); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("NewUserHandler() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func Test_userHandler_findUserByUserID_200(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID int
	}

	tests := []struct {
		name          string
		prepareMockFn func(m *mock.MockIUserUsecase)
		want          gin.HandlerFunc
	}{
		{
			name: "findUserByUserID success response",
			prepareMockFn: func(m *mock.MockIUserUsecase) {
				args := args{
					ctx:    &gin.Context{},
					userID: 1,
				}
				// gomock.Any()にしないとエラー
				m.EXPECT().FindUserByUserID(gomock.Any(), args.userID).Return(&model.User{
					ID:       1,
					Name:     "name1",
					Icon:     "/icon1",
					Gender:   1,
					BirthDay: time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
					Location: 1,
				}, nil)
			},
			want: func(c *gin.Context) {
				c.JSON(http.StatusOK, &model.User{})
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			// mock登録
			controller := gomock.NewController(t)
			defer controller.Finish()
			mock := mock.NewMockIUserUsecase(controller)
			tt.prepareMockFn(mock)
			// user handler 初期化(内部にmock)
			uh := NewUserHandler(mock)

			// エンドポイントを登録
			// router := InitRouter()
			router := gin.Default()
			// ハンドラー登録(mockが登録させたハンドラー)
			router.GET(usersAPIRoot+"/:user_id", func(c *gin.Context) { c.Set("user_id", 1) }, uh.findUserByUserID())

			// レスポンスを受け止める*httptest.ResponseRecorder
			rec := httptest.NewRecorder()
			// テストで送信するリクエスト
			req := httptest.NewRequest(http.MethodGet, usersAPIRoot+"/1", nil)
			// リクエスト送信
			router.ServeHTTP(rec, req)

			// 結果確認
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.JSONEq(
				t,
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
		})
	}
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
