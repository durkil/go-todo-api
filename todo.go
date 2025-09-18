package main

import "gorm.io/gorm"

// Todo represents a todo item
// @Description Todo item information
type Todo struct {
	gorm.Model
	Title     string `json:"title" example:"Complete project"`
	Completed bool   `json:"completed" gorm:"default:false" example:"false"`
	UserID    uint   `json:"user_id" example:"1"`
}
