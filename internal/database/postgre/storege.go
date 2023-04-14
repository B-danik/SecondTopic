package repository

import (
	sql "github.com/B-danik/SecondTopic/internal/database/postgre/sql"
	"github.com/B-danik/SecondTopic/todo"
	"github.com/jmoiron/sqlx"
)

type IAuthorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(email, password string) (todo.User, error)
}

type IBook interface {
	CreateBook(name string) (int, error)
	GetBook()
}

type Repository struct {
	Authorization IAuthorization
	Book          IBook
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: sql.NewAuthPostgres(db),
		Book:          sql.NewBook(db),
	}
}
