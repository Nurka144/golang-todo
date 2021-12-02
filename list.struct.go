package todo

type TodoList struct {
	Id          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Completed   bool   `json:"completed" db:"completed"`
}
