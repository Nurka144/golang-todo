package repository

type Authorization interface {
}

type TodoList interface {
}

type Repository struct {
	Authorization
	TodoList
}

func NewRepository() *Repository {
	return &Repository{}
}
