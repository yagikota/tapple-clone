package mysql

import (
	"context"
	"database/sql"
	"fmt"

	constant "github.com/CyberAgentHack/2208-ace-go-server/pkg"
	"github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/entity"
	"github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/model"
	domain "github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/repository"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type userRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) domain.IUserRepository {
	return &userRepository{
		DB: db,
	}
}

func (ur *userRepository) FindUserByUserID(ctx context.Context, userID int) (*model.User, error) {
	tx, err := txFromContext(ctx)
	if err != nil {
		return nil, err
	}

	whereID := fmt.Sprintf("%s = ?", model.UserColumns.ID)
	return model.Users(
		qm.Where(whereID, userID),
	).One(ctx, tx)
}

func (ur *userRepository) FindAllUsers(ctx context.Context) (model.UserSlice, error) {
	tx, err := txFromContext(ctx)
	if err != nil {
		return nil, err
	}

	return model.Users().All(ctx, tx)
}

func (ur *userRepository) FindAllRooms(ctx context.Context, userID int) (model.RoomSlice, error) {
	tx, err := txFromContext(ctx)
	if err != nil {
		return nil, err
	}

	whereRoomID := fmt.Sprintf("%s = ?)", "rooms.id in (select room_id from room_users where user_id")
	wherePartnerID := fmt.Sprintf("%s <> ?", model.RoomUserColumns.UserID)
	orderBy := fmt.Sprintf("%s DESC", model.MessageColumns.CreatedAt)

	return model.Rooms(
		qm.Where(whereRoomID, userID),
		qm.Load(model.RoomRels.Messages, qm.OrderBy(orderBy)),
		qm.Load(model.RoomRels.RoomUsers, qm.Where(wherePartnerID, userID)),
		qm.Load(qm.Rels(model.RoomRels.RoomUsers, model.RoomUserRels.User)),
	).All(ctx, tx)
}

// TODO: 例えば、localhost:8080/v1/users/2/rooms/３でもアクセスできてしまうので、改善が必要
// 認証機能を導入すれば改善できそう(アクセストークンをヘッダーに乗せるとか)
func (ur *userRepository) FindRoomDetailByRoomID(ctx context.Context, userID, roomID, messageID int) (*model.Room, error) {
	boil.DebugMode = true
	tx, err := txFromContext(ctx)
	if err != nil {
		return nil, err
	}

	whereRoomID := fmt.Sprintf("%s = ?", model.RoomColumns.ID)
	// 自分が2番目にくるようにsort←チャットルームのNameとIconを[0]で取得するため
	orderBy := fmt.Sprintf("%s = ?", model.RoomUserColumns.UserID)
	orderByMessage := fmt.Sprintf("%s DESC", model.MessageColumns.CreatedAt)
	whereMessageCreatedAt := fmt.Sprintf("%s <= ?", model.MessageColumns.CreatedAt)

	if messageID == 0 {
		return model.Rooms(
			qm.Where(whereRoomID, roomID),
			qm.Load(entity.RoomRels.RoomUsers, qm.OrderBy(orderBy, userID)),
			qm.Load(qm.Rels(entity.RoomRels.RoomUsers, entity.RoomUserRels.User)),
			qm.Load(entity.RoomRels.Messages, qm.OrderBy(orderByMessage), qm.Limit(constant.LimitRecord)),
		).One(ctx, tx)
	}

	message, err := model.FindMessage(ctx, tx, int64(messageID))
	if err != nil {
		return nil, err
	}
	messageCreatedAt := message.CreatedAt
	return model.Rooms(
		qm.Where(whereRoomID, roomID),
		qm.Load(entity.RoomRels.RoomUsers, qm.OrderBy(orderBy, userID)),
		qm.Load(qm.Rels(entity.RoomRels.RoomUsers, entity.RoomUserRels.User)),
		qm.Load(entity.RoomRels.Messages, qm.Where(whereMessageCreatedAt, messageCreatedAt), qm.OrderBy(orderByMessage), qm.Limit(constant.LimitRecord)),
	).One(ctx, tx)
}

// TODO: 自身が所属しているルームにのみ送信できるようにする 現状localhost:8080/v1/users/2/rooms/3でも送信できてしまう
// 認証機能を導入すれば改善できそう(アクセストークンをヘッダーに乗せるとか)
func (ur *userRepository) SendMessage(ctx context.Context, m *model.Message) (*model.Message, error) {
	tx, err := txFromContext(ctx)
	if err != nil {
		return nil, err
	}

	err = m.Insert(ctx, tx, boil.Infer())
	if err != nil {
		return nil, err
	}
	return m, nil
}
