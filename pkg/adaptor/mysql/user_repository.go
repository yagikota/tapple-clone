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

func (ur *userRepository) FindByUserID(ctx context.Context, userID int) (*entity.User, error) {
	whereID := fmt.Sprintf("%s = ?", entity.UserColumns.ID)
	return entity.Users(
		qm.Where(whereID, userID),
	).One(ctx, ur.DB)
}

func (ur *userRepository) FindAll(ctx context.Context) (entity.UserSlice, error) {
	boil.DebugMode = true
	return entity.Users().All(ctx, ur.DB)
}

func (ur *userRepository) FindAllRooms(ctx context.Context, UserID int) (entity.RoomSlice, error) {
	boil.DebugMode = true
	// whereUserID := fmt.Sprintf("%s = ?", entity.MessageColumns.UserID)

	// // 未読数取得
	// fmt.Println(entity.Messages(qm.Where(whereUserID, UserID), qm.Where("is_read=?", false)).Count(ctx, ur.DB))

	return entity.Rooms(qm.Load(entity.RoomRels.Messages, qm.OrderBy(entity.MessageColumns.CreatedAt+" DESC")), qm.Load(entity.RoomRels.RoomUsers, qm.And("user_id=?", 2))).All(ctx, ur.DB)
}
