package service

import (
	"github.com/Nurka144/todo-app"
	"github.com/Nurka144/todo-app/pkg/repository"
)

type Authorization interface {
	CreateToUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	PaseTokenMiddleware(token string) (int, error)
}

type TodoList interface {
	CreateTodo(todo todo.TodoList) (int, error)
}

type Service struct {
	Authorization
	TodoList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
	}
}
