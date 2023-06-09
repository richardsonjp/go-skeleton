package model

import (
	"time"
)

// TableName overrides the table name used by Context to `context`
func (Context) TableName() string {
	return "context"
}

// nb: if using custom type data, eg=json, array, etc. always define the type for gorm
type Context struct {
	ID           uint      `gorm:"primarykey"`
	ContextTag   string    `gorm:"column:context_tag"`
	PrimaryKey   uint      `gorm:"column:primary_key"`
	SecondaryKey uint      `gorm:"column:secondary_key"`
	CreatedAt    time.Time `gorm:"column:created_at;type:datetime"`
	UpdatedAt    time.Time `gorm:"column:updated_at;type:datetime"`
}
