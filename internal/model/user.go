package model

import "time"

type User struct {
	ID        int64     `sql:"primary_key;auto_increment"`
	CreatedAt time.Time `sql:"type:DATETIME;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `sql:"type:DATETIME;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

func (User) TableName() string {
	return "users"
}
