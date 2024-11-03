package model

import (
	"time"
)

type TodoItem struct {
	Id          int         `json:"id" gorm:"column:id;"`
	Titles      string      `json:"titles" gorm:"column:titles;"`
	Description string      `json:"description" gorm:"column:description;"`
	Status      *ItemStatus `json:"status" gorm:"column:status;"`
	CreatedAt   time.Time   `json:"createdAt" gorm:"column:createdAt;"`
	UpdatedAt   time.Time   `json:"updatedAt,omitempty" gorm:"column:updatedAt;"`
}

func (TodoItem) TableName() string { return "tasks" }

type TodoItemCreation struct {
	Id          int         `json:"id" gorm:"column:id;"`
	Titles      string      `json:"titles" gorm:"column:titles;"`
	Description string      `json:"description" gorm:"column:description;"`
	Status      *ItemStatus `json:"status" gorm:"column:status;"`
	CreatedAt   time.Time   `json:"createdAt" gorm:"column:createdAt;"`
	UpdatedAt   time.Time   `json:"updatedAt,omitempty" gorm:"column:updatedAt;"`
}

func (TodoItemCreation) TableName() string { return TodoItem{}.TableName() }

type ToDoItemUpdate struct {
	Title       *string   `json:"title" gorm:"column:titles;"`
	Description *string   `json:"description" gorm:"column:description;"`
	Status      *string   `json:"status" gorm:"column:status;"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty" gorm:"column:updatedAt;"`
}

func (ToDoItemUpdate) TableName() string { return TodoItem{}.TableName() }
