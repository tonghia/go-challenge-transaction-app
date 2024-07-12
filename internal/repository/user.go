package repository

import (
	"context"
	"fmt"

	"github.com/tonghia/go-challenge-transaction-app/internal/model"
	"gorm.io/gorm"
)

//go:generate mockgen --source=./user.go -destination=./mock/user.go -package=mock
type UserRepositorier interface {
	GetByID(ctx context.Context, id int64) (*model.User, error)
}

type UserRepository struct {
	db        *gorm.DB
	tableName string
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db:        db,
		tableName: model.User{}.TableName(),
	}
}

func (r *UserRepository) getDB(ctx context.Context) *gorm.DB {
	return r.db.Table(r.tableName).WithContext(ctx)
}

func (r *UserRepository) GetByID(ctx context.Context, id int64) (*model.User, error) {
	var rs *model.User

	if err := r.getDB(ctx).Where("id = ?", id).First(&rs).Error; err != nil {
		return rs, fmt.Errorf("failed to get user: %w", err)
	}

	return rs, nil
}
