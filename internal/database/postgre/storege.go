package repository

import (
	sql "github.com/B-danik/SecondTopic/internal/database/postgre/sql"
	"github.com/B-danik/SecondTopic/todo"
	"github.com/jmoiron/sqlx"
)

type IAuthorization interface {
	Create(user todo.User) (int, error)
	Get(email, password string) (todo.User, error)
}

type IBook interface {
	Create(name string) (int, error)
	Get(ID int) (todo.Book, error)
	GetAll() ([]todo.Book, error)
	Delete(ID int) error
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
