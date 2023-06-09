package model

import (
	"mgw/mgw-resi/internal/model/enum"
	"time"
)

// TableName overrides the table name used by Page to `page`
func (Page) TableName() string {
	return "page"
}

type PageStatus = enum.GenericStatus

// nb: if using custom type data, eg=json, array, etc. always define the type for gorm
type Page struct {
	ID        uint       `gorm:"primarykey"`
	Name      string     `gorm:"column:name"`
	Status    PageStatus `gorm:"column:status"`
	CreatedAt time.Time  `gorm:"column:created_at;type:datetime"`
	UpdatedAt time.Time  `gorm:"column:updated_at;type:datetime"`
}
