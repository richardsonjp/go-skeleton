package model

import (
	"time"
)

// TableName overrides the table name used by Misc to `misc`
func (Misc) TableName() string {
	return "misc"
}

// nb: if using custom type data, eg=json, array, etc. always define the type for gorm
type Misc struct {
	ID        uint      `gorm:"primarykey"`
	Name      string    `gorm:"column:name"`
	Value     string    `gorm:"column:value"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime"`
}
