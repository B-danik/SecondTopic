package service

import (
	repository "github.com/B-danik/SecondTopic/internal/database/postgre"
	"github.com/B-danik/SecondTopic/pkg/book"
	"github.com/B-danik/SecondTopic/pkg/transactions"
	users "github.com/B-danik/SecondTopic/pkg/users"
	"github.com/B-danik/SecondTopic/todo"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	Create(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type Book interface {
	Create(name string) (int, error)
	Get(ID int) (todo.Book, error)
	GetAll() ([]todo.Book, error)
	Delete(ID int) error
}

type Transactions interface {
}

type Service struct {
	Authorization
	Book
	Transactions
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: users.NewAuthService(repo.Authorization),
		Book:          book.NewBook(repo.Book),
		Transactions:  transactions.NewBook(repo.Transactions),
	}
}
