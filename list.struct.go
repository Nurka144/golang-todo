package todo

type TodoList struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}
