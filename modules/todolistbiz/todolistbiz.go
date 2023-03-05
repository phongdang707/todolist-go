package todolistbiz

import (
	"ToDoList/modules/todolistmodel"
	"errors"
)

type CreateTaskStore interface {
	Create(data todolistmodel.TodoListModel) error
}

type CreateTaskBiz struct {
	store CreateTaskStore
}

func NewCreateRestaurantBiz(store CreateTaskStore) *CreateTaskBiz {
	return &CreateTaskBiz{store: store}
}

func (c CreateTaskBiz) CreateTaskWithBiz(data *todolistmodel.TodoListModel) error {
	if data.Name == "" {
		return errors.New("name can not be blank")
	}

	err := c.store.Create(*data)
	return err
}
