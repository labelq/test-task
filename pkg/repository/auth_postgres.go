package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"test"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user test.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, password) values ($1, $2) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Name, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(name, password string) (test.User, error) {
	var user test.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE name = $1 AND password = $2", usersTable)
	err := r.db.Get(&user, query, name, password)

	return user, err
}
