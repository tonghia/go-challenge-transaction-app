package repository

import (
	"context"
	"fmt"

	"github.com/tonghia/go-challenge-transaction-app/internal/model"
	"gorm.io/gorm"
)

//go:generate mockgen --source=./account.go -destination=./mock/account.go -package=mock
type AccountRepositorier interface {
	GetByID(ctx context.Context, id int64) (*model.Account, error)
}

type AccountRepository struct {
	db        *gorm.DB
	tableName string
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{
		db:        db,
		tableName: model.Account{}.TableName(),
	}
}

func (r *AccountRepository) getDB(ctx context.Context) *gorm.DB {
	return r.db.Table(r.tableName).WithContext(ctx)
}

func (r *AccountRepository) GetByID(ctx context.Context, id int64) (*model.Account, error) {
	var account model.Account

	if err := r.getDB(ctx).Where("id = ?", id).First(&account).Error; err != nil {
		return nil, fmt.Errorf("failed to get account: %w", err)
	}

	return &account, nil
}
