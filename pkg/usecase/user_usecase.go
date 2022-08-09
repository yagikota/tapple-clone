package usecase

import (
	"context"

	"github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/service"
	"github.com/CyberAgentHack/2208-ace-go-server/pkg/usecase/model"
)

type IUserUsecase interface {
	User(ctx context.Context, userID int) (*model.User, error)
	Users(ctx context.Context) (model.UserSlice, error)
}

type userUsecase struct {
	userService service.IUserService
}

func NewUserUsecase(uService service.IUserService) IUserUsecase {
	return &userUsecase{
		userService: uService,
	}
}

func (uu *userUsecase) User(ctx context.Context, userID int) (*model.User, error) {
	entity, err := uu.userService.User(ctx, userID)
	if err != nil {
		return nil, err
	}
	return model.UserFromEntity(entity), err
}

func (uu *userUsecase) Users(ctx context.Context) (model.UserSlice, error) {
	entities, err := uu.userService.Users(ctx)
	if err != nil {
		return nil, err
	}

	// メモリ確保
	uSlice := make(model.UserSlice, 0, len(entities))
	for _, entity := range entities {
		uSlice = append(uSlice, model.UserFromEntity(entity))
	}

	return uSlice, nil
}
