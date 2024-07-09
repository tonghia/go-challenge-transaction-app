package repository

import (
	"github.com/tonghia/go-challenge-transaction-app/internal/model"
	"gorm.io/gorm"
)

type UserRepositorier interface {
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

func (r *UserRepository) getDB() *gorm.DB {
	return r.db.Table(r.tableName)
}
