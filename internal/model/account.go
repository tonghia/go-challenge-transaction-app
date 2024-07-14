package model

import "time"

type Account struct {
	ID        int64     `sql:"primary_key;auto_increment"`
	UserID    int64     `sql:"type:BIGINT;index"`
	Name      string    `sql:"type:VARCHAR(255)"`
	Bank      string    `sql:"type:VARCHAR(10)"`
	CreatedAt time.Time `sql:"type:DATETIME;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `sql:"type:DATETIME;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

func (Account) TableName() string {
	return "accounts"
}
