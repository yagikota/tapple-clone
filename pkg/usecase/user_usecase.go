package usecase

import (
	"context"

	"github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/service"
	"github.com/CyberAgentHack/2208-ace-go-server/pkg/usecase/model"
)

type IUserUsecase interface {
	FindByUserID(ctx context.Context, userID int) (*model.User, error)
	FindAll(ctx context.Context) (model.UserSlice, error)
	SendMessage(ctx context.Context, userID int, roomID int, m *model.NewMessage) error
}

type userUsecase struct {
	userService service.IUserService
}

func NewUserUsecase(uService service.IUserService) IUserUsecase {
	return &userUsecase{
		userService: uService,
	}
}

func (uu *userUsecase) FindByUserID(ctx context.Context, userID int) (*model.User, error) {
	entity, err := uu.userService.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return model.UserFromEntity(entity), err
}

func (uu *userUsecase) FindAll(ctx context.Context) (model.UserSlice, error) {
	entities, err := uu.userService.FindAll(ctx)
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

func (uu *userUsecase) SendMessage(ctx context.Context, userID int, roomID int, m *model.NewMessage) error {
	return uu.userService.SendMessage(ctx, model.MessageToEntity(m, userID, roomID))
}
