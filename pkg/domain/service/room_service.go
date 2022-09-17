//go:generate mockgen -source $GOFILE -package=mock -destination ../../mock/service/$GOFILE

package service

import (
	"context"

	"github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/model"
	domain "github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/repository"
)

type IRoomService interface {
	FindAllRooms(ctx context.Context, userID int) (model.RoomSlice, error)
	FindRoomDetailByRoomID(ctx context.Context, userID, roomID, messageID int) (*model.Room, error)
	SendMessage(ctx context.Context, m *model.Message) (*model.Message, error)
}

type roomService struct {
	roomRepository domain.IRoomRepository
}

func NewRoomService(rr domain.IRoomRepository) IRoomService {
	return &roomService{
		roomRepository: rr,
	}
}

func (rs *roomService) FindAllRooms(ctx context.Context, userID int) (model.RoomSlice, error) {
	return rs.roomRepository.FindAllRooms(ctx, userID)
}

func (rs *roomService) FindRoomDetailByRoomID(ctx context.Context, userID, roomID, messageID int) (*model.Room, error) {
	return rs.roomRepository.FindRoomDetailByRoomID(ctx, userID, roomID, messageID)
}

func (rs *roomService) SendMessage(ctx context.Context, m *model.Message) (*model.Message, error) {
	return rs.roomRepository.SendMessage(ctx, m)
}
