package model

import (
	"time"
)

// TableName overrides the table name used by ContextPath to `context_path`
func (ContextPath) TableName() string {
	return "context_path"
}

// nb: if using custom type data, eg=json, array, etc. always define the type for gorm
type ContextPath struct {
	ID         uint      `gorm:"primarykey"`
	ContextTag string    `gorm:"column:context_tag"`
	RoleID     uint      `gorm:"column:role_id"`
	PathID     uint      `gorm:"column:path_id"`
	CreatedAt  time.Time `gorm:"column:created_at;type:datetime"`
	UpdatedAt  time.Time `gorm:"column:updated_at;type:datetime"`
}
