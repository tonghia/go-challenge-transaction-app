package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/tonghia/go-challenge-transaction-app/internal/model"
	"gorm.io/gorm"
)

//go:generate mockgen --source=./account_transaction.go -destination=./mock/account_transaction.go -package=mock
type AccountTransactionRepositorier interface {
	GetByUserAccount(ctx context.Context, userID, accountID int64) ([]*model.AccountTransaction, error)
	DeleteByTransactionID(ctx context.Context, txnID int64) error
}

type AccountTransactionRepository struct {
	db        *gorm.DB
	tableName string
}

func NewAccountTransactionRepository(db *gorm.DB) *AccountTransactionRepository {
	return &AccountTransactionRepository{
		db:        db,
		tableName: model.AccountTransaction{}.TableName(),
	}
}

func (r *AccountTransactionRepository) getDB(ctx context.Context) *gorm.DB {
	return r.db.Table(r.tableName).WithContext(ctx)
}

func (r *AccountTransactionRepository) GetByUserAccount(ctx context.Context, userID, accountID int64) ([]*model.AccountTransaction, error) {
	var rs []*model.AccountTransaction

	q := r.getDB(ctx).Where("user_id = ?", userID)

	if accountID != 0 {
		q.Where("account_id = ?", accountID)
	}

	if err := q.Find(&rs).Error; err != nil {
		return rs, fmt.Errorf("failed to get account transactions: %w", err)
	}

	return rs, nil
}

func (r *AccountTransactionRepository) DeleteByTransactionID(ctx context.Context, txnID int64) error {
	return r.getDB(ctx).Where("id = ?", txnID).Update("deleted_at", time.Now()).Error
}
