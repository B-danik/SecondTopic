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
	Create(name string, price int) (int, error)
	Get(ID int) (todo.Book, error)
	GetList() ([]todo.Book, error)
	Delete(ID int) error
}

type IRent interface {
	GetRent(id int) ([]todo.Rent, error)
	CraeteRent(user_id int, book_id int) error
}

type Repository struct {
	Authorization IAuthorization
	Book          IBook
	Rent          IRent
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: sql.NewAuthPostgres(db),
		Book:          sql.NewBook(db),
		Rent:          sql.NewRent(db),
	}
}
