package model

import (
	"go-skeleton/internal/model/enum"
	"time"
)

// TableName overrides the table name used by Admin to `admin`
func (Admin) TableName() string {
	return "admin"
}

type AdminStatus = enum.AdminStatus

// nb: if using custom type data, eg=json, array, etc. always define the type for gorm
type Admin struct {
	ID          uint        `gorm:"primarykey"`
	UniqueID    string      `gorm:"column:unique_id"`
	PartnerID   uint        `gorm:"column:partner_id"`
	Email       string      `gorm:"column:email"`
	Name        string      `gorm:"column:name"`
	PhoneNumber string      `gorm:"column:phone_number"`
	Password    string      `gorm:"column:password"`
	Role        StringArr   `gorm:"column:role;type:json"`
	Status      AdminStatus `gorm:"column:status"`
	LastLoginAt time.Time   `gorm:"column:last_login_at;type:datetime"`
	CreatedAt   time.Time   `gorm:"column:created_at;type:datetime"`
	UpdatedAt   time.Time   `gorm:"column:updated_at;type:datetime"`
}
