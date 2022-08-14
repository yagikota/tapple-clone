package http

import (
	"context"
	"log"
	"net/http"
	"reflect"
	"testing"

	mock "github.com/CyberAgentHack/2208-ace-go-server/pkg/mock/usecase"
	"github.com/CyberAgentHack/2208-ace-go-server/pkg/usecase"
	"github.com/CyberAgentHack/2208-ace-go-server/pkg/usecase/model"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

func TestNewUserHandler(t *testing.T) {
	type args struct {
		uu usecase.IUserUsecase
	}
	tests := []struct {
		name string
		args args
		want *userHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserHandler(tt.args.uu); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userHandler_findUserByUserID_200(t *testing.T) {
	gin.SetMode(gin.TestMode)

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
				m.EXPECT().FindUserByUserID(args.ctx, args.userID).Return(&model.User{}, nil)
			},
			want: func(c *gin.Context) {
				c.JSON(http.StatusOK, &model.User{})
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()
			mock := mock.NewMockIUserUsecase(controller)
			// NewUserHandler(mock)
			tt.prepareMockFn(mock)
			uh := &userHandler{
				uUsecase: mock,
			}

			if got := uh.findUserByUserID(); cmp.Diff(got, tt.want) == "" {
				log.Println(cmp.Diff(got, tt.want))
				t.Errorf("userHandler.findUserByUserID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userHandler_findUsers(t *testing.T) {
	type fields struct {
		uUsecase usecase.IUserUsecase
	}
	tests := []struct {
		name   string
		fields fields
		want   gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uh := &userHandler{
				uUsecase: tt.fields.uUsecase,
			}
			if got := uh.findUsers(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userHandler.findUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userHandler_findRooms(t *testing.T) {
	type fields struct {
		uUsecase usecase.IUserUsecase
	}
	tests := []struct {
		name   string
		fields fields
		// buildStubs func(store *mock.Mo)
		want gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uh := &userHandler{
				uUsecase: tt.fields.uUsecase,
			}
			if got := uh.findRooms(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userHandler.findRooms() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userHandler_findRoomDetailByRoomID(t *testing.T) {
	type fields struct {
		uUsecase usecase.IUserUsecase
	}
	tests := []struct {
		name   string
		fields fields
		want   gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uh := &userHandler{
				uUsecase: tt.fields.uUsecase,
			}
			if got := uh.findRoomDetailByRoomID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userHandler.findRoomDetailByRoomID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userHandler_sendMessage(t *testing.T) {
	type fields struct {
		uUsecase usecase.IUserUsecase
	}
	tests := []struct {
		name   string
		fields fields
		want   gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uh := &userHandler{
				uUsecase: tt.fields.uUsecase,
			}
			if got := uh.sendMessage(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userHandler.sendMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
