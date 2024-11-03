package model

import (
	"social-todo-list/common"
)

type TodoItem struct {
	common.SQLModel
	Titles      string      `json:"titles" gorm:"column:titles;"`
	Description string      `json:"description" gorm:"column:description;"`
	Status      *ItemStatus `json:"status" gorm:"column:status;"`
}

func (TodoItem) TableName() string { return "tasks" }

type TodoItemCreation struct {
	Id          int         `json:"-" gorm:"column:id;"`
	Title       string      `json:"title" gorm:"column:titles;"`
	Description string      `json:"description" gorm:"column:description;"`
	Status      *ItemStatus `json:"status" gorm:"column:status;"`
}

func (TodoItemCreation) TableName() string { return TodoItem{}.TableName() }

type ToDoItemUpdate struct {
	Title       *string `json:"title" gorm:"column:titles;"`
	Description *string `json:"description" gorm:"column:description;"`
	Status      *string `json:"status" gorm:"column:status;"`
}

func (ToDoItemUpdate) TableName() string { return TodoItem{}.TableName() }
