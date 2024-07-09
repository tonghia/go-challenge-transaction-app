package model

type User struct{}

func (User) TableName() string {
	return "users"
}
