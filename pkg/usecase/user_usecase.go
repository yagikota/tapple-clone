//go:generate mockgen -source $GOFILE -package=mock -destination ../mock/usecase/$GOFILE
package usecase

import (
	"context"

	"github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/service"
	"github.com/CyberAgentHack/2208-ace-go-server/pkg/usecase/model"
)

type IUserUsecase interface {
	FindUserByUserID(ctx context.Context, userID int) (*model.User, error)
	FindAllUsers(ctx context.Context) (model.UserSlice, error)
	FindUserDetailByUserID(ctx context.Context, userID int) (*model.UserDetail, error)
}

type userUsecase struct {
	userService service.IUserService
}

func NewUserUsecase(us service.IUserService) IUserUsecase {
	return &userUsecase{
		userService: us,
	}
}

func (uu *userUsecase) FindUserByUserID(ctx context.Context, userID int) (*model.User, error) {
	mu, err := uu.userService.FindUserByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return model.UserFromDomainModel(mu), err
}

func (uu *userUsecase) FindAllUsers(ctx context.Context) (model.UserSlice, error) {
	muSlice, err := uu.userService.FindAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	uSlice := make(model.UserSlice, 0, len(muSlice))
	for _, mu := range muSlice {
		uSlice = append(uSlice, model.UserFromDomainModel(mu))
	}

	return uSlice, nil
}

func (uu *userUsecase) FindUserDetailByUserID(ctx context.Context, userID int) (*model.UserDetail, error) {
	mu, err := uu.userService.FindUserDetailByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return model.UserDetailFromDomainModel(mu), nil
}
