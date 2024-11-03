package storage

import (
	"context"
	"social-todo-list/common"
	"social-todo-list/modules/item/model"
)

func (sql *sqlStorage) CreateItem(ctx context.Context, data *model.TodoItemCreation) error {
	if err := sql.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
func (sql *sqlStorage) GetItem(ctx context.Context, id int) (*model.TodoItem, error) {
	var data model.TodoItem
	if err := sql.db.First(&data, id).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
func (sql *sqlStorage) ListItem(ctx context.Context, filter *model.Filter, paging *common.Paging) ([]model.TodoItem, error) {
	var result []model.TodoItem
	db := sql.db.Where("status <> ?", "Deleted")
	if f := filter; f != nil {
		if v := f.Status; v != "" {
			db = db.Where("status = ?", v)
		}
	}

	if err := db.Table(model.TodoItem{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := db.Order("id asc").
		Offset((paging.Page - 1) * paging.Size).
		Limit(paging.Size).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
func (sql *sqlStorage) UpdateItem(ctx context.Context, id int, data *model.ToDoItemUpdate) error {
	if err := sql.db.Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

func (sql *sqlStorage) DeleteItem(ctx context.Context, id int) error {
	if err := sql.db.Table(model.TodoItem{}.TableName()).Where("id = ?", id).Update("status", "Deleted").Error; err != nil {
		return err
	}
	return nil
}
