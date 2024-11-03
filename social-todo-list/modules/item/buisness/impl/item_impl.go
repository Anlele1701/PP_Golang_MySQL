package impl

import (
	"context"
	"errors"
	"social-todo-list/common"
	"social-todo-list/modules/item/buisness"
	"social-todo-list/modules/item/model"
)

type ItemBuisnessImpl struct {
	store buisness.ItemInterface
}

func NewItemBuisness(store buisness.ItemInterface) *ItemBuisnessImpl {
	return &ItemBuisnessImpl{store: store}
}

func (buisness *ItemBuisnessImpl) CreateItem(ctx context.Context, data *model.TodoItemCreation) error {
	return buisness.store.CreateItem(ctx, data)
}

func (buisness *ItemBuisnessImpl) GetItem(ctx context.Context, id int) (*model.TodoItem, error) {
	return buisness.store.GetItem(ctx, id)
}
func (buisness *ItemBuisnessImpl) ListItem(ctx context.Context,
	filter *model.Filter,
	paging *common.Paging) ([]model.TodoItem, error) {
	data, err := buisness.store.ListItem(ctx, filter, paging)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (buisness *ItemBuisnessImpl) UpdateItem(ctx context.Context, id int, data *model.ToDoItemUpdate) error {
	dataGet, err := buisness.store.GetItem(ctx, id)
	if err != nil {
		return err
	}
	if dataGet.Status != nil && *dataGet.Status == model.ItemStatusDeleted {
		return errors.New("item is not found")

	}
	return buisness.store.UpdateItem(ctx, id, data)
}
func (buisness *ItemBuisnessImpl) DeleteItem(ctx context.Context, id int) error {
	data, err := buisness.store.GetItem(ctx, id)
	if err != nil {
		return err
	}
	if data.Status != nil && *data.Status == model.ItemStatusDeleted {
		return errors.New("item is not found")
	}
	return buisness.store.DeleteItem(ctx, id)
}
