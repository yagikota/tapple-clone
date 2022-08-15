package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/entity"
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

func (ur *userRepository) FindUserByUserID(ctx context.Context, userID int) (*entity.User, error) {
	tx, err := TxFromContext(ctx)
	if err != nil {
		return nil, err
	}

	whereID := fmt.Sprintf("%s = ?", entity.UserColumns.ID)
	return entity.Users(
		qm.Where(whereID, userID),
	).One(ctx, tx)
}

func (ur *userRepository) FindAllUsers(ctx context.Context) (entity.UserSlice, error) {
	tx, err := TxFromContext(ctx)
	if err != nil {
		return nil, err
	}

	return entity.Users().All(ctx, tx)
}

func (ur *userRepository) FindAllRooms(ctx context.Context, userID int) (entity.RoomSlice, error) {
	tx, err := TxFromContext(ctx)
	if err != nil {
		return nil, err
	}

	whereRoomID := fmt.Sprintf("%s = ?)", "rooms.id in (select room_id from room_users where user_id")
	wherePartnerID := fmt.Sprintf("%s <> ?", entity.RoomUserColumns.UserID)
	orderBy := fmt.Sprintf("%s DESC", entity.MessageColumns.CreatedAt)

	return entity.Rooms(
		qm.Where(whereRoomID, userID),
		qm.Load(entity.RoomRels.Messages, qm.OrderBy(orderBy)),
		qm.Load(entity.RoomRels.RoomUsers, qm.Where(wherePartnerID, userID)),
		qm.Load(qm.Rels(entity.RoomRels.RoomUsers, entity.RoomUserRels.User)),
	).All(ctx, tx)
}

// TODO: 例えば、localhost:8080/v1/users/2/rooms/３でもアクセスできてしまうので、改善が必要
// 認証機能を導入すれば改善できそう(アクセストークンをヘッダーに乗せるとか)
func (ur *userRepository) FindRoomDetailByRoomID(ctx context.Context, userID, roomID int) (*entity.Room, error) {
	boil.DebugMode = true
	tx, err := TxFromContext(ctx)
	if err != nil {
		return nil, err
	}

	whereRoomID := fmt.Sprintf("%s = ?", entity.RoomColumns.ID)
	// 自分が2番目にくるようにsort←チャットルームのNameとIconを[0]で取得するため
	orderBy := fmt.Sprintf("%s = ?", entity.RoomUserColumns.UserID)
	return entity.Rooms(
		qm.Where(whereRoomID, roomID),
		qm.Load(entity.RoomRels.RoomUsers, qm.OrderBy(orderBy, userID)),
		qm.Load(qm.Rels(entity.RoomRels.RoomUsers, entity.RoomUserRels.User)),
		qm.Load(entity.RoomRels.Messages),
	).One(ctx, tx)
}

// TODO: 自身が所属しているルームにのみ送信できるようにする 現状localhost:8080/v1/users/2/rooms/3でも送信できてしまう
// 認証機能を導入すれば改善できそう(アクセストークンをヘッダーに乗せるとか)
func (ur *userRepository) SendMessage(ctx context.Context, m *entity.Message) error {
	tx, err := TxFromContext(ctx)
	if err != nil {
		return err
	}

	return m.Insert(ctx, tx, boil.Infer())
}
