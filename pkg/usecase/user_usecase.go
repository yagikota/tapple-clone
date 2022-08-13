package usecase

import (
	"context"
	"sort"

	"github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/service"
	"github.com/CyberAgentHack/2208-ace-go-server/pkg/usecase/model"
)

type IUserUsecase interface {
	FindByUserID(ctx context.Context, userID int) (*model.User, error)
	FindAllUsers(ctx context.Context) (model.UserSlice, error)
	FindAllRooms(ctx context.Context, userID int) (model.RoomSlice, error)
	FindRoomDetailByRoomID(ctx context.Context, userID, roomID int) (*model.Room, error) // TODO: 引数の型を省略すべきかどうか調べる
	SendMessage(ctx context.Context, userID int, roomID int, m *model.NewMessage) error
}

type userUsecase struct {
	userService service.IUserService
}

func NewUserUsecase(uService service.IUserService) IUserUsecase {
	return &userUsecase{
		userService: uService,
	}
}

func (uu *userUsecase) FindByUserID(ctx context.Context, userID int) (*model.User, error) {
	entity, err := uu.userService.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return model.UserFromEntity(entity), err
}

func (uu *userUsecase) FindAllUsers(ctx context.Context) (model.UserSlice, error) {
	entities, err := uu.userService.FindAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	// メモリ確保
	uSlice := make(model.UserSlice, 0, len(entities))
	for _, entity := range entities {
		uSlice = append(uSlice, model.UserFromEntity(entity))
	}

	return uSlice, nil
}

func (uu *userUsecase) FindAllRooms(ctx context.Context, userID int) (model.RoomSlice, error) {
	entities, err := uu.userService.FindAllRooms(ctx, userID)
	if err != nil {
		return nil, err
	}

	// メモリ確保
	rSlice := make(model.RoomSlice, 0, len(entities))
	for _, entity := range entities {
		rSlice = append(rSlice, model.RoomFromEntity(entity))
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

	return rSlice, nil
}

func (uu *userUsecase) FindRoomDetailByRoomID(ctx context.Context, userID, roomID int) (*model.Room, error) {
	entity, err := uu.userService.FindRoomDetailByRoomID(ctx, userID, roomID)
	if err != nil {
		return nil, err
	}

	return model.RoomDetailFromEntity(entity), nil
}

func (uu *userUsecase) SendMessage(ctx context.Context, userID int, roomID int, m *model.NewMessage) error {
	return uu.userService.SendMessage(ctx, model.MessageToEntity(m, userID, roomID))
}
