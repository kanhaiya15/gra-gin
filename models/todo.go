package models

import "github.com/jinzhu/gorm"

// TodoModel TodoModel
type TodoModel struct {
	gorm.Model
	Title     string `json:"title" binding:"required"`
	Completed int    `json:"completed" binding:"required"`
}

// TransformedTodo TransformedTodo
type TransformedTodo struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
