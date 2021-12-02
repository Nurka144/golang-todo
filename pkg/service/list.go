package service

import (
	"github.com/Nurka144/todo-app"
	"github.com/Nurka144/todo-app/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) CreateTodo(todo todo.TodoList) (int, error) {
	return s.repo.CreateTodo(todo)
}

func (s *TodoListService) GetAllTodo() ([]todo.TodoList, error) {
	return s.repo.GetAllTodo()
}
