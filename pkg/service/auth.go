package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/Nurka144/todo-app"
	"github.com/Nurka144/todo-app/pkg/repository"
)

const salt = "test_password_hash"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateToUser(user todo.User) (int, error) {
	fmt.Println("afa9999")
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateToUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
