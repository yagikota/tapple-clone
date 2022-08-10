package usecase

import (
	"context"

	"github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/service"
	"github.com/CyberAgentHack/2208-ace-go-server/pkg/usecase/model"
)

type IUserUsecase interface {
	FindByUserID(ctx context.Context, userID int) (*model.User, error)
	FindAll(ctx context.Context) (model.UserSlice, error)
	FindAllRooms(ctx context.Context, UserID int) (model.RoomSlice, error)
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

func (uu *userUsecase) FindAllRooms(ctx context.Context, UserID int) (model.RoomSlice, error) {
	entities, err := uu.userService.FindAllRooms(ctx, UserID)
	if err != nil {
		return nil, err
	}

	// メモリ確保
	rSlice := make(model.RoomSlice, 0, len(entities))
	for _, entity := range entities {
		rSlice = append(rSlice, model.RoomFromEntity(entity))
	}

	return rSlice, nil
}
