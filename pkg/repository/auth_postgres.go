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
	fmt.Println("ascascascascacasc")
	fmt.Println(query, user.Name, user.Username, user.Password)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	fmt.Println("qweqweqweqweqwe")
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
