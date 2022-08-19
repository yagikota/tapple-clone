// TODO: 下記コマンド動作しないので、動作するようにする
//
//go:generate mockgen -source pkg/domain/repository/$GOFILE -package=mock -destination pkg/mock/repository/$GOFILE
package domain

import (
	"context"

	"github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/entity"
)

type IUserRepository interface {
	FindUserByUserID(ctx context.Context, userID int) (*entity.User, error)
	FindAllUsers(ctx context.Context) (entity.UserSlice, error)
	FindAllRooms(ctx context.Context, userID int) (entity.RoomSlice, error)
	FindRoomDetailByRoomID(ctx context.Context, userID, roomID, messageID int) (*entity.Room, error)
	SendMessage(ctx context.Context, m *entity.Message) (*entity.Message, error)
}
