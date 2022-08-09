package service

// サービス層の必要性について
// https://christina04.hatenablog.com/entry/go-clean-architecture

import (
	"context"

	"github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/entity"
	domain "github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/repository"
)

type IUserService interface {
	User(ctx context.Context, userID int) (*entity.User, error)
	Users(ctx context.Context) (entity.UserSlice, error)
}

type userService struct {
	userRepository domain.IUserRepository
}

func NewUserService(ur domain.IUserRepository) IUserService {
	return &userService{
		userRepository: ur,
	}
}

func (us *userService) User(ctx context.Context, userID int) (*entity.User, error) {
	return us.userRepository.User(ctx, userID)
}

func (us *userService) Users(ctx context.Context) (entity.UserSlice, error) {
	return us.userRepository.Users(ctx)
}
