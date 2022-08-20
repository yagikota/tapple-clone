package usecase

import (
	"context"
	"sort"

	"github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/service"
	"github.com/CyberAgentHack/2208-ace-go-server/pkg/usecase/model"
)

type IUserUsecase interface {
	FindUserByUserID(ctx context.Context, userID int) (*model.User, error)
	FindAllUsers(ctx context.Context) (model.UserSlice, error)
	FindAllRooms(ctx context.Context, userID int) (*model.Rooms, error)
	FindRoomDetailByRoomID(ctx context.Context, userID, roomID, messageID int) (*model.RoomDetail, error)
	SendMessage(ctx context.Context, userID, roomID int, m *model.NewMessage) (*model.Message, error)
	FindUserDetailByUserID(ctx context.Context, userID int) (*model.UserDetail, error)
}

type userUsecase struct {
	userService service.IUserService
}

func NewUserUsecase(uService service.IUserService) IUserUsecase {
	return &userUsecase{
		userService: uService,
	}
}

func (uu *userUsecase) FindUserByUserID(ctx context.Context, userID int) (*model.User, error) {
	entity, err := uu.userService.FindUserByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return model.UserFromDomainModel(entity), err
}

func (uu *userUsecase) FindAllUsers(ctx context.Context) (model.UserSlice, error) {
	entities, err := uu.userService.FindAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	uSlice := make(model.UserSlice, 0, len(entities))
	for _, entity := range entities {
		uSlice = append(uSlice, model.UserFromDomainModel(entity))
	}

	return uSlice, nil
}

func (uu *userUsecase) FindAllRooms(ctx context.Context, userID int) (*model.Rooms, error) {
	entities, err := uu.userService.FindAllRooms(ctx, userID)
	if err != nil {
		return nil, err
	}

	rSlice := make(model.RoomSlice, 0, len(entities))
	for _, entity := range entities {
		rSlice = append(rSlice, model.RoomFromDomainModel(entity))
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

func (uu *userUsecase) FindRoomDetailByRoomID(ctx context.Context, userID, roomID, messageID int) (*model.RoomDetail, error) {
	entity, err := uu.userService.FindRoomDetailByRoomID(ctx, userID, roomID, messageID)
	if err != nil {
		return nil, err
	}

	return model.RoomDetailFromDomainModel(entity), nil
}

func (uu *userUsecase) SendMessage(ctx context.Context, userID, roomID int, m *model.NewMessage) (*model.Message, error) {
	entity, err := uu.userService.SendMessage(ctx, m.ToDomainModel(userID, roomID))
	if err != nil {
		return nil, err
	}

	return model.MessageFromDomainModel(entity), nil
}

func (uu *userUsecase) FindUserDetailByUserID(ctx context.Context, userID int) (*model.UserDetail, error) {
	mu, err := uu.userService.FindUserDetailByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return model.UserDetailFromDomainModel(mu), nil
}
