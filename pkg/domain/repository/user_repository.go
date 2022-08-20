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
	FindAllRooms(ctx context.Context, userID int) (model.RoomSlice, error)
	FindRoomDetailByRoomID(ctx context.Context, userID, roomID, messageID int) (*model.Room, error)
	SendMessage(ctx context.Context, m *model.Message) (*model.Message, error)
	FindUserDetailByUserID(ctx context.Context, userID int) (*model.User, error)
}
