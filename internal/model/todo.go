package model

//example file

import (
	"gopkg.in/guregu/null.v4"
)

type TodoStatus uint8

const (
	TodoStatusPending TodoStatus = iota + 1
	TodoStatusProcess
	TodoStatusDone
)

// Todo define table columns
// struct validate only do it before database create or update
type Todo struct {
	ID        uint       `gorm:"primarykey"`
	Task      string     `gorm:"column:task" validate:"required"`
	Status    TodoStatus `gorm:"column:status" validate:"required,gt=0,lte=3"`
	CreatedAt TimeAt     `gorm:"created_at;not null"`
	UpdatedAt TimeAt     `gorm:"updated_at;not null"`
	DeletedAt null.Time  `gorm:"deleted_at;type;datetime"`
}

func (todo *Todo) StatusStr() string {
	return []string{"Null", "Pending", "Process", "Done"}[todo.Status]
}
