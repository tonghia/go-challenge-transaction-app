package repository

import (
	"context"

	"github.com/tonghia/go-challenge-transaction-app/internal/model"
	"gorm.io/gorm"
)

//go:generate mockgen --source=./account.go -destination=./mock/account.go -package=mock
type AccountRepositorier interface {
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
