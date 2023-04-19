package service

import (
	repository "github.com/B-danik/SecondTopic/internal/database/postgre"
	"github.com/B-danik/SecondTopic/pkg/book"
	rent "github.com/B-danik/SecondTopic/pkg/rent"
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
	Create(name string, price int) (int, error)
	Get(ID int) (todo.Book, error)
	GetList() ([]todo.Book, error)
	Delete(ID int) error
}

type Rent interface {
	GetRent(id int) ([]todo.Rent, error)
	CraeteRent(user_id int, book_id int) error
}

type Service struct {
	Authorization
	Book
	Rent
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: users.NewAuthService(repo.Authorization),
		Book:          book.NewBook(repo.Book),
		Rent:          rent.NewRent(repo.Rent),
	}
}
