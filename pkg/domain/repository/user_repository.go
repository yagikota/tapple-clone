//go:generate mockgen -source $GOFILE -package=mock -destination ../../mock/repository/$GOFILE
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
