//go:generate mockgen -source $GOFILE -package=mock -destination ../../mock/repository/$GOFILE
package domain

import (
	"context"

	"github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/model"
)

type IRoomRepository interface {
	FindAllRooms(ctx context.Context, userID int) (model.RoomSlice, error)
	FindRoomDetailByRoomID(ctx context.Context, userID, roomID, messageID int) (*model.Room, error)
	SendMessage(ctx context.Context, m *model.Message) (*model.Message, error)
}
