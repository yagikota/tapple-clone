package usecase

import (
	"context"
	"sort"

	"github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/service"
	"github.com/CyberAgentHack/2208-ace-go-server/pkg/usecase/model"
)

type IRoomUsecase interface {
	FindAllRooms(ctx context.Context, userID int) (*model.Rooms, error)
	FindRoomDetailByRoomID(ctx context.Context, userID, roomID, messageID int) (*model.RoomDetail, error)
	SendMessage(ctx context.Context, userID, roomID int, m *model.NewMessage) (*model.Message, error)
}

type roomUsecase struct {
	roomService service.IRoomService
}

func NewRoomUsecase(rs service.IRoomService) IRoomUsecase {
	return &roomUsecase{
		roomService: rs,
	}
}

func (ru *roomUsecase) FindAllRooms(ctx context.Context, userID int) (*model.Rooms, error) {
	mrSlice, err := ru.roomService.FindAllRooms(ctx, userID)
	if err != nil {
		return nil, err
	}

	rSlice := make(model.RoomSlice, 0, len(mrSlice))
	for _, mr := range mrSlice {
		rSlice = append(rSlice, model.RoomFromDomainModel(mr))
	}

	sort.Slice(rSlice, func(i, j int) bool {
		//is_pinnedがtrueを優先
		if rSlice[i].IsPinned && !rSlice[j].IsPinned {
			return true
		} else if !rSlice[i].IsPinned && rSlice[j].IsPinned {
			return false
		}

		//LatestMessageの降順
		return rSlice[i].LatestMessage.CreatedAt.After(rSlice[j].LatestMessage.CreatedAt)
	})

	return &model.Rooms{Rooms: rSlice}, nil
}

func (ru *roomUsecase) FindRoomDetailByRoomID(ctx context.Context, userID, roomID, messageID int) (*model.RoomDetail, error) {
	mr, err := ru.roomService.FindRoomDetailByRoomID(ctx, userID, roomID, messageID)
	if err != nil {
		return nil, err
	}

	return model.RoomDetailFromDomainModel(mr), nil
}

func (ru *roomUsecase) SendMessage(ctx context.Context, userID, roomID int, m *model.NewMessage) (*model.Message, error) {
	mm, err := ru.roomService.SendMessage(ctx, m.ToDomainModel(userID, roomID))
	if err != nil {
		return nil, err
	}

	return model.MessageFromDomainModel(mm), nil
}
