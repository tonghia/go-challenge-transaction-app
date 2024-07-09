package repository

import (
	"github.com/tonghia/go-challenge-transaction-app/internal/must"
	"gorm.io/gorm"
)

const (
	_defaultCreateBatchSize = 20
	_defaultMaxLimit        = 100
)

type Repository struct {
	db      *gorm.DB
	isInTxn bool
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
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

func (r Repository) Close() error {
	return must.Close(r.db)
}
