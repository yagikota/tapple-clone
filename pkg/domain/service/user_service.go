package service

// サービス層の必要性について
// https://christina04.hatenablog.com/entry/go-clean-architecture

import (
	"context"

	"github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/model"
	"github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/repository"
)

type IUserService interface {
	FindUserByUserID(ctx context.Context, userID int) (*model.User, error)
	FindAllUsers(ctx context.Context) (model.UserSlice, error)
	FindAllRooms(ctx context.Context, userID int) (model.RoomSlice, error)
	FindRoomDetailByRoomID(ctx context.Context, userID, roomID, messageID int) (*model.Room, error)
	SendMessage(ctx context.Context, m *model.Message) (*model.Message, error)
}

type userService struct {
	userRepository domain.IUserRepository
}

func NewUserService(ur domain.IUserRepository) IUserService {
	return &userService{
		userRepository: ur,
	}
}

func (us *userService) FindUserByUserID(ctx context.Context, userID int) (*model.User, error) {
	return us.userRepository.FindUserByUserID(ctx, userID)
}

func (us *userService) FindAllUsers(ctx context.Context) (model.UserSlice, error) {
	return us.userRepository.FindAllUsers(ctx)
}

func (us *userService) FindAllRooms(ctx context.Context, userID int) (model.RoomSlice, error) {
	return us.userRepository.FindAllRooms(ctx, userID)
}

func (us *userService) FindRoomDetailByRoomID(ctx context.Context, userID, roomID, messageID int) (*model.Room, error) {
	return us.userRepository.FindRoomDetailByRoomID(ctx, userID, roomID, messageID)
}

func (us *userService) SendMessage(ctx context.Context, m *model.Message) (*model.Message, error) {
	return us.userRepository.SendMessage(ctx, m)
}
