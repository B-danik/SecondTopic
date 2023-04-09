package authUser

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/B-danik/SecondTopic/todo"
)

type AuthPostgre struct {
	db *sqlx.DB
}

const (
	usersTable = "users"
)

func NewAuthPostgres(db *sqlx.DB) *AuthPostgre {
	return &AuthPostgre{db: db}
}

func (p *AuthPostgre) CreateUser(user todo.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (email,name, lastname, password_hash) values ($1, $2, $3,$4) RETURNING id", usersTable)

	row := p.db.QueryRow(query, user.Email, user.Name, user.LastName, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (p *AuthPostgre) GetUser(email, password string) (todo.User, error) {
	var user todo.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE email=$1 AND password_hash=$2", usersTable)
	err := p.db.Get(&user, query, email, password)

	return user, err
}
