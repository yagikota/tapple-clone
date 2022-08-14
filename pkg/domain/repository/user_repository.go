package domain

import (
	"context"

	"github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/entity"
)

type IUserRepository interface {
	FindByUserID(ctx context.Context, userID int) (*entity.User, error)
	FindAllUsers(ctx context.Context) (entity.UserSlice, error)
	FindAllRooms(ctx context.Context, userID int) (entity.RoomSlice, error)
	FindRoomDetailByRoomID(ctx context.Context, userID int, roomID int) (*entity.Room, error)
	SendMessage(ctx context.Context, m *entity.Message) error
}
