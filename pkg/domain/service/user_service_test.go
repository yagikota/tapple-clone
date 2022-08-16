package service

import (
	"context"
	"reflect"
	"testing"

	"github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/entity"
	domain "github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/repository"
)

func TestNewUserService(t *testing.T) {
	type args struct {
		ur domain.IUserRepository
	}
	tests := []struct {
		name string
		args args
		want IUserService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserService(tt.args.ur); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_FindUserByUserID(t *testing.T) {
	type fields struct {
		userRepository domain.IUserRepository
	}
	type args struct {
		ctx    context.Context
		userID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &userService{
				userRepository: tt.fields.userRepository,
			}
			got, err := us.FindUserByUserID(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.FindUserByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userService.FindUserByUserID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_FindAllUsers(t *testing.T) {
	type fields struct {
		userRepository domain.IUserRepository
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.UserSlice
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &userService{
				userRepository: tt.fields.userRepository,
			}
			got, err := us.FindAllUsers(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.FindAllUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userService.FindAllUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_FindAllRooms(t *testing.T) {
	type fields struct {
		userRepository domain.IUserRepository
	}
	type args struct {
		ctx    context.Context
		userID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.RoomSlice
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &userService{
				userRepository: tt.fields.userRepository,
			}
			got, err := us.FindAllRooms(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.FindAllRooms() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userService.FindAllRooms() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_FindRoomDetailByRoomID(t *testing.T) {
	type fields struct {
		userRepository domain.IUserRepository
	}
	type args struct {
		ctx    context.Context
		userID int
		roomID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Room
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &userService{
				userRepository: tt.fields.userRepository,
			}
			got, err := us.FindRoomDetailByRoomID(tt.args.ctx, tt.args.userID, tt.args.roomID)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.FindRoomDetailByRoomID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userService.FindRoomDetailByRoomID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_SendMessage(t *testing.T) {
	type fields struct {
		userRepository domain.IUserRepository
	}
	type args struct {
		ctx context.Context
		m   *entity.Message
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &userService{
				userRepository: tt.fields.userRepository,
			}
			if err := us.SendMessage(tt.args.ctx, tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("userService.SendMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
