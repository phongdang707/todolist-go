package todoliststorage

import (
	"ToDoList/modules/todolistmodel"
)

func (s *sqlStore) Create(data todolistmodel.TodoListModel) error {
	_, err := s.db.Exec(`INSERT INTO notes (name, description) VALUES ($1, $2)`, data.Name, data.Description)
	return err
}
