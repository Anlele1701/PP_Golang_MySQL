package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type TodoItem struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreatedAt   *time.Time `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
}

func main() {
	fmt.Println("Hello, World!")
	now := time.Now().UTC()
	item := TodoItem{
		Id:          1,
		Title:       "Create a todo list",
		Description: "Create a todo list using Go",
		Status:      "In Progress",
		CreatedAt:   &now,
		UpdatedAt:   nil,
	}
	jsonData, error := json.MarshalIndent(item, "", "    ")
	if error != nil {
		fmt.Println(error)
		return
	}
	fmt.Println(string(jsonData))

	jsonString := `{"id":1,"title":"Create a todo list","description":"Create a todo list using Go","status":"In Progress","createdAt":"2021-07-01T00:00:00Z","updatedAt":"2021-07-01T00:00:00Z"}`
	var item2 TodoItem
	if err := json.Unmarshal([]byte(jsonString), &item2); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(item2)
}
