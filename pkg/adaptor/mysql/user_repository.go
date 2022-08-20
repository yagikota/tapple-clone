package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/model"
	domain "github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/repository"
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

func (ur *userRepository) FindUserDetailByUserID(ctx context.Context, userID int) (*model.User, error) {
	tx, err := txFromContext(ctx)
	if err != nil {
		return nil, err
	}

	whereID := fmt.Sprintf("%s = ?", model.UserColumns.ID)

	return model.Users(
		qm.Where(whereID, userID),
		qm.Load(model.UserRels.Hobbies),
		qm.Load(model.UserRels.UserProfileImages),
	).One(ctx, tx)
}
