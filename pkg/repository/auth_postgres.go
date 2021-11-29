package repository

import (
	"fmt"

	"github.com/Nurka144/todo-app"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}
func (r *AuthPostgres) CreateToUser(user todo.User) (int, error) {
	var id int
	query := fmt.Sprintf("insert into %s (name, username, password) values ($1, $2, $3) returning id", usersTable)

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(username string, password string) (todo.User, error) {
	var user todo.User
	query := fmt.Sprintf("select * from users where username = $1 and password = $2")

	err := r.db.Get(&user, query, username, password)
	return user, err
}
