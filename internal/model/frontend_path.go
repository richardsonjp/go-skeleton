package model

import (
	"time"
)

// TableName overrides the table name used by FrontendPath to `frontend_path`
func (FrontendPath) TableName() string {
	return "frontend_path"
}

// nb: if using custom type data, eg=json, array, etc. always define the type for gorm
type FrontendPath struct {
	ID        uint      `gorm:"primarykey"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime"`
}
