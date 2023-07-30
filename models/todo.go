package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Todo struct {
	Id        uuid.UUID `json:"id" gorm:"primaryKey; type:uuid;"`
	Title     string    `json:"title" gorm:"not null"`
	Completed bool      `json:"completed" gorm:"not null; default:false"`
}

func (todo *Todo) BeforeCreate(tx *gorm.DB) (err error) {
	todo.Id = uuid.NewV4()
	return
}