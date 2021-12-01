package repository

import (
	"fmt"

	"github.com/Nurka144/todo-app"
	"github.com/jmoiron/sqlx"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r *TodoListPostgres) CreateTodo(todo todo.TodoList) (int, error) {
	var id int
	query := fmt.Sprintf("insert into %s (name, description, completed) values ($1, $2, $3) returning id", todoListsTable)
	row := r.db.QueryRow(query, todo.Name, todo.Description, false)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
