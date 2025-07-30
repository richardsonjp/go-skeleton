package model

import (
	"time"
)

// TableName overrides the table name used by BackendPath to `backend_path`
func (BackendPath) TableName() string {
	return "backend_path"
}

// nb: if using custom type data, eg=json, array, etc. always define the type for gorm
type BackendPath struct {
	ID        uint      `gorm:"primarykey"`
	Name      string    `gorm:"column:name"`
	Method    string    `gorm:"column:method"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime"`
}
