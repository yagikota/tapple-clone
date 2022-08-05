package service

import (
	"context"

	"github.com/CyberAgentHack/2208-ace-go-server/domain"
	"github.com/CyberAgentHack/2208-ace-go-server/domain/entity"
)

type IUserService interface {
	User(ctx context.Context, userID int) (*entity.User, error) // 1ユーザー取得
	Users(ctx context.Context) (entity.UserSlice, error)
}

type userService struct {
	userRepository domain.IUserRepository
}

func NewUserService(userRepository domain.IUserRepository) IUserService {
	return &userService {
		userRepository: userRepository,
	}
}

func (us *userService) User(ctx context.Context, userID int) (*entity.User, error) {
	return nil, nil
}

func (us *userService) Users(ctx context.Context) (entity.UserSlice, error){
	return nil, nil
}