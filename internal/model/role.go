package model

import (
	"go-skeleton/internal/model/enum"
	"time"
)

// TableName overrides the table name used by Role to `role`
func (Role) TableName() string {
	return "role"
}

type RoleStatus = enum.GenericStatus

// nb: if using custom type data, eg=json, array, etc. always define the type for gorm
type Role struct {
	ID        uint       `gorm:"primarykey"`
	Name      string     `gorm:"column:name"`
	Status    RoleStatus `gorm:"column:status"`
	CreatedAt time.Time  `gorm:"column:created_at;type:datetime"`
	UpdatedAt time.Time  `gorm:"column:updated_at;type:datetime"`
}
