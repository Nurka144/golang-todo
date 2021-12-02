package repository

import (
	"github.com/Nurka144/todo-app"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateToUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

type TodoList interface {
	CreateTodo(todo todo.TodoList) (int, error)
	GetAllTodo() ([]todo.TodoList, error)
}

type Repository struct {
	Authorization
	TodoList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
	}
}
