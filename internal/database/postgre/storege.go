package repository

import (
	authUser "github.com/B-danik/SecondTopic/internal/database/postgre/sql"
	"github.com/B-danik/SecondTopic/todo"
	"github.com/jmoiron/sqlx"
)

type IAuthorization interface {
	CreateUser(user todo.User) (int, error)
}

type Repository struct {
	Authorization IAuthorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: authUser.NewAuthPostgres(db),
	}
}
