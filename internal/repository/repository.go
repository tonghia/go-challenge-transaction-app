package repository

import (
	"gorm.io/gorm"
)

const (
	_defaultCreateBatchSize = 20
	_defaultMaxLimit        = 100
)

type Repository struct {
	db                 *gorm.DB
	isInTxn            bool
	User               UserRepositorier
	Account            AccountRepositorier
	AccountTransaction AccountTransactionRepositorier
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db:                 db,
		User:               NewUserRepository(db),
		Account:            NewAccountRepository(db),
		AccountTransaction: NewAccountTransactionRepository(db),
	}
}

func (r Repository) Transaction(txFunc func(*Repository) error) (err error) {
	return r.db.Transaction(func(tx *gorm.DB) error {
		txRepositories := r
		txRepositories.db = tx
		txRepositories.isInTxn = true

		return txFunc(&txRepositories)
	})
}
