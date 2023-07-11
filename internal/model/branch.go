package model

import (
	"go/skeleton/internal/model/enum"
	"time"
)

// TableName overrides the table name used by Branch to `branch`
func (Branch) TableName() string {
	return "branch"
}

type BranchStatus = enum.GenericStatus

// nb: if using custom type data, eg=json, array, etc. always define the type for gorm
type Branch struct {
	ID        uint         `gorm:"primarykey"`
	Name      string       `gorm:"column:name"`
	Code      string       `gorm:"column:code"`
	UniqueID  string       `gorm:"column:unique_id"`
	Status    BranchStatus `gorm:"column:status"`
	CreatedAt time.Time    `gorm:"column:created_at;type:datetime"`
	UpdatedAt time.Time    `gorm:"column:updated_at;type:datetime"`
}
