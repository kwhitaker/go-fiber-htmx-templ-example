package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title     string `gorm:"not null;default:null" json:"title" validate:"required"`
	Completed bool   `json:"completed"`
}
