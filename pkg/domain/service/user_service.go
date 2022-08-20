package service

// サービス層の必要性について
// https://christina04.hatenablog.com/entry/go-clean-architecture

import (
	"context"

	"github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/model"
	domain "github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/repository"
)

type IUserService interface {
	FindUserByUserID(ctx context.Context, userID int) (*model.User, error)
	FindAllUsers(ctx context.Context) (model.UserSlice, error)
	FindUserDetailByUserID(ctx context.Context, userID int) (*model.User, error)
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

func (us *userService) FindUserDetailByUserID(ctx context.Context, userID int) (*model.User, error) {
	return us.userRepository.FindUserDetailByUserID(ctx, userID)
}
