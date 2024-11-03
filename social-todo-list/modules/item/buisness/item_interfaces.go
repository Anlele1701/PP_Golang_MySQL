package buisness

import (
	"context"
	"social-todo-list/common"
	"social-todo-list/modules/item/model"
)

type ItemInterface interface {
	CreateItem(ctx context.Context, data *model.TodoItemCreation) error
	GetItem(ctx context.Context, id int) (*model.TodoItem, error)
	ListItem(ctx context.Context, filter *model.Filter,
		paging *common.Paging) ([]model.TodoItem, error)
	UpdateItem(ctx context.Context, id int, data *model.ToDoItemUpdate) error
	DeleteItem(ctx context.Context, id int) error
}
