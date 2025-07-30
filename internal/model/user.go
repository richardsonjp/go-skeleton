package model

import (
	"go-skeleton/internal/model/enum"
	"time"
)

// TableName overrides the table name used by User to `user`
func (User) TableName() string {
	return "user"
}

type UserStatus = enum.GenericStatus

// nb: if using custom type data, eg=json, array, etc. always define the type for gorm
type User struct {
	ID          uint       `gorm:"primarykey"`
	Name        string     `gorm:"column:name"`
	Password    string     `gorm:"column:password"`
	Email       string     `gorm:"column:email"`
	Phone       string     `gorm:"column:phone"`
	RoleID      uint       `gorm:"column:role_id"`
	Status      UserStatus `gorm:"column:status"`
	LastLoginAt *time.Time `gorm:"column:last_login_at;type:datetime"`
	CreatedAt   time.Time  `gorm:"column:created_at;type:datetime"`
	UpdatedAt   time.Time  `gorm:"column:updated_at;type:datetime"`
}
