package service

import "github.com/Nurka144/todo-app/pkg/repository"

type Authorization interface {
}

type TodoList interface {
}

type Service struct {
	Authorization
	TodoList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
