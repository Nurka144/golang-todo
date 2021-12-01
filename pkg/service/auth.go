package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/Nurka144/todo-app"
	"github.com/Nurka144/todo-app/pkg/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

const (
	salt    = "test_password_hash"
	signkey = "213edasf12efe"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	password = generatePasswordHash(password)

	user, err := s.repo.GetUser(username, password)

	if err != nil {
		logrus.Fatalf("Error in get user: %s", err.Error())
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		}, UserId: user.Id})

	return token.SignedString([]byte(signkey))
}

func (s *AuthService) CreateToUser(user todo.User) (int, error) {

	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateToUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) PaseTokenMiddleware(accessToken string) (int, error) {

	fmt.Println(accessToken)

	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("vsdvsdvsdsv")
			return nil, errors.New("invalid signing method")
		}

		return []byte(signkey), nil
	})

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)

	if !ok {
		return 0, errors.New("invalid parse")
	}

	return claims.UserId, nil

}
