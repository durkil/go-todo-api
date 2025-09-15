package main

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title     string `json:"title"`
	Completed bool   `json:"completed" gorm:"default:false"`
	UserID    uint   `json:"user_id"`
}
