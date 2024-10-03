package entities

import "gorm.io/gorm"

type Todolist struct {
	gorm.Model
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      bool   `json:"status" gorm:"default:false"`
}
