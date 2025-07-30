package model

import (
	"go-skeleton/internal/model/enum"
	"time"
)

// TableName overrides the table name used by Credential to `credential`
func (Credential) TableName() string {
	return "credential"
}

type CredentialStatus = enum.GenericStatus

// nb: if using custom type data, eg=json, array, etc. always define the type for gorm
type Credential struct {
	ID        uint             `gorm:"primarykey"`
	UserID    uint             `gorm:"column:user_id"`
	Secret    string           `gorm:"column:secret"`
	Status    CredentialStatus `gorm:"column:status"`
	ExpiredAt time.Time        `gorm:"column:expired_at;type:datetime"`
	CreatedAt time.Time        `gorm:"column:created_at;type:datetime"`
	UpdatedAt time.Time        `gorm:"column:updated_at;type:datetime"`
}
