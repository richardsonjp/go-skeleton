package model

import (
	"time"
)

// TableName overrides the table name used by Country to `country`
func (Country) TableName() string {
	return "country"
}

// nb: if using custom type data, eg=json, array, etc. always define the type for gorm
type Country struct {
	ID        uint      `gorm:"primarykey"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime"`
}
