package repository

import (
	"context"
	"time"

	"github.com/tonghia/go-challenge-transaction-app/internal/model"
	"gorm.io/gorm"
)

//go:generate mockgen --source=./account_transaction.go -destination=./mock/account_transaction.go -package=mock
type AccountTransactionRepositorier interface {
	CreateOne(ctx context.Context, txn *model.AccountTransaction) error
	GetByID(ctx context.Context, id int64) (*model.AccountTransaction, error)
	GetByUser(ctx context.Context, userID int64) ([]*model.AccountTransaction, error)
	GetByUserAccount(ctx context.Context, userID, accountID int64) ([]*model.AccountTransaction, error)
	UpdateOne(ctx context.Context, txn *model.AccountTransaction) error
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

func (r *AccountTransactionRepository) CreateOne(ctx context.Context, txn *model.AccountTransaction) error {

	return r.getDB(ctx).Create(txn).Error
}

func (r *AccountTransactionRepository) GetByID(ctx context.Context, id int64) (*model.AccountTransaction, error) {
	var rs *model.AccountTransaction

	return rs, r.getDB(ctx).Where("id = ?", id).First(&rs).Error
}

func (r *AccountTransactionRepository) GetByUser(ctx context.Context, userID int64) ([]*model.AccountTransaction, error) {
	var rs []*model.AccountTransaction

	return rs, r.getDB(ctx).Where("user_id = ?", userID).Find(&rs).Error
}

func (r *AccountTransactionRepository) GetByUserAccount(ctx context.Context, userID, accountID int64) ([]*model.AccountTransaction, error) {
	var rs []*model.AccountTransaction

	return rs, r.getDB(ctx).Where("user_id = ? AND account_id = ?", userID, accountID).Find(&rs).Error
}

func (r *AccountTransactionRepository) UpdateOne(ctx context.Context, txn *model.AccountTransaction) error {

	return r.getDB(ctx).Save(txn).Error
}

func (r *AccountTransactionRepository) DeleteByTransactionID(ctx context.Context, txnID int64) error {

	return r.getDB(ctx).Where("id = ?", txnID).Update("deleted_at", time.Now()).Error
}
