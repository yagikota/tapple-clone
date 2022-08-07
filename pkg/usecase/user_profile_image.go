package usecase

import (
	"context"

	"github.com/CyberAgentHack/2208-ace-go-server/domain/service"
	"github.com/CyberAgentHack/2208-ace-go-server/usecase/model"
)

type IUserProfileImageUsecase interface {
	UserProfileImage(ctx context.Context) (*model.UserProfileImage, error)
	UserProfileImages(ctx context.Context) (model.UserProfileImageSlice, error)
}

type userProfileImageUsecase struct {
	upiService service.IUserProfileImageService
}

func NewuserProfileImageUsecase(upis service.IUserProfileImageService) IUserProfileImageUsecase {
	return &userProfileImageUsecase{
		upiService: upis,
	}
}

func (upiu *userProfileImageUsecase) UserProfileImage(ctx context.Context) (*model.UserProfileImage, error) {
	entity, err := upiu.upiService.UserProfileImage(ctx)
	if err != nil {
		return nil, err
	}

	return model.UserProfileImageFromEntity(entity), nil
}

func (upiu *userProfileImageUsecase) UserProfileImages(ctx context.Context) (model.UserProfileImageSlice, error) {
	entities, err := upiu.upiService.UserProfileImages(ctx)
	if err != nil {
		return nil, err
	}

	upiSlice := make(model.UserProfileImageSlice, 0, len(entities))
	for _, entity := range entities {
		upiSlice = append(upiSlice, model.UserProfileImageFromEntity(entity))
	}

	return upiSlice, nil
}
