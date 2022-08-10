package domain

import (
	"context"

	"github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/entity"
)

type IUserRepository interface {
	FindByUserID(ctx context.Context, userID int) (*entity.User, error)     // 1ユーザー取得
	FindAll(ctx context.Context) (entity.UserSlice, error)                  // 全ユーザー取得
	FindAllRooms(ctx context.Context, UserID int) (entity.RoomSlice, error) //任意ユーザーの全ルーム取得
}
