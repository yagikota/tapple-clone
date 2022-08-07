package service

import (
	"context"

	"github.com/CyberAgentHack/2208-ace-go-server/domain"
	"github.com/CyberAgentHack/2208-ace-go-server/domain/entity"
)

type IUserProfileImageService interface {
	UserProfileImage(ctx context.Context) (*entity.UserProfileImage, error)
	UserProfileImages(ctx context.Context) (entity.UserProfileImageSlice, error)
}

type userProfileImageService struct {
	userProfileImageRepository domain.IUserProfileImageRepository
}

func NewuserProfileImageService(upir domain.IUserProfileImageRepository) IUserProfileImageService {
	return &userProfileImageService{
		userProfileImageRepository: upir,
	}
}

func (upis *userProfileImageService) UserProfileImage(ctx context.Context) (*entity.UserProfileImage, error) {
	return upis.userProfileImageRepository.UserProfileImage(ctx)
}

func (upis *userProfileImageService) UserProfileImages(ctx context.Context) (entity.UserProfileImageSlice, error) {
	return upis.userProfileImageRepository.UserProfileImages(ctx)
}
