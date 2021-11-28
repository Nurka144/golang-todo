package repository

import (
	"github.com/Nurka144/todo-app"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateToUser(user todo.User) (int, error)
}

type TodoList interface {
}

type Repository struct {
	Authorization
	TodoList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
