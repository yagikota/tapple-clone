// TODO: 下記コマンド動作しないので、動作するようにする
//
//go:generate mockgen -source pkg/domain/repository/$GOFILE -package=mock -destination pkg/mock/repository/$GOFILE
package domain

import (
	"context"

	"github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/model"
)

type IUserRepository interface {
	FindUserByUserID(ctx context.Context, userID int) (*model.User, error)
	FindAllUsers(ctx context.Context) (model.UserSlice, error)
	FindUserDetailByUserID(ctx context.Context, userID int) (*model.User, error)
}
