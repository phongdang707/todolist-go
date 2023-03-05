package todolistmodel

type TodoListModel struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (t *TodoListModel) TableName() string {
	return "notes"
}
