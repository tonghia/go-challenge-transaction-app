package model

import (
	"time"

	"github.com/shopspring/decimal"
)

const (
	TransactionTypeWithdraw = "withdraw"
	TransactionTypeDeposit  = "deposit"
)

type AccountTransaction struct {
	ID              int64           `sql:"primary_key;auto_increment"`
	UserID          int64           `sql:"type:BIGINT;index;index_columns:user_id,account_id"`
	AccountID       int64           `sql:"type:BIGINT"`
	Bank            string          `sql:"type:VARCHAR(20)"`
	Amount          decimal.Decimal `sql:"type:DECIMAL(20,8)"`
	TransactionType string          `sql:"type:VARCHAR(20)"`
	DeleteAt        *time.Time      `sql:"type:DATETIME"`
	CreatedAt       time.Time       `sql:"type:DATETIME;default:CURRENT_TIMESTAMP"`
	UpdatedAt       time.Time       `sql:"type:DATETIME;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

func (AccountTransaction) TableName() string {
	return "account_transactions"
}
