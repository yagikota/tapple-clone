package service

// サービス層の必要性について
// https://christina04.hatenablog.com/entry/go-clean-architecture

import (
	"context"

	"github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/entity"
	domain "github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/repository"
)

type IUserService interface {
	FindByUserID(ctx context.Context, userID int) (*entity.User, error)
	FindAll(ctx context.Context) (entity.UserSlice, error)
	FindAllRooms(ctx context.Context, userID int) (entity.RoomSlice, error)
	FindAllRoomMessages(ctx context.Context, userID, roomID int) (entity.MessageSlice, error)
}

type userService struct {
	userRepository domain.IUserRepository
}

func NewUserService(ur domain.IUserRepository) IUserService {
	return &userService{
		userRepository: ur,
	}
}

func (us *userService) FindByUserID(ctx context.Context, userID int) (*entity.User, error) {
	return us.userRepository.FindByUserID(ctx, userID)
}

func (us *userService) FindAll(ctx context.Context) (entity.UserSlice, error) {
	return us.userRepository.FindAll(ctx)
}

func (us *userService) FindAllRooms(ctx context.Context, userID int) (entity.RoomSlice, error) {
	return us.userRepository.FindAllRooms(ctx, userID)
}

func (us *userService) FindAllRoomMessages(ctx context.Context, userID, roomID int) (entity.MessageSlice, error) {
	return us.userRepository.FindAllRoomMessages(ctx, userID, roomID)
}
