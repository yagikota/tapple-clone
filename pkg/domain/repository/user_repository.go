package domain

import (
	"context"

	"github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/entity"
)

type IUserRepository interface {
	User(ctx context.Context, userID int) (*entity.User, error) // 1ユーザー取得
	Users(ctx context.Context) (entity.UserSlice, error)        // 全ユーザー取得
}
