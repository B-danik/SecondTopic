package service

import (
	repository "github.com/B-danik/SecondTopic/internal/database/postgre"
	"github.com/B-danik/SecondTopic/pkg/book"
	users "github.com/B-danik/SecondTopic/pkg/users"
	"github.com/B-danik/SecondTopic/todo"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type Book interface {
	CreateBook(name string) (int, error)
	GetBook()
}

type Service struct {
	Authorization
	Book
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: users.NewAuthService(repo.Authorization),
		Book:          book.NewBook(repo.Book),
	}
}
