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
