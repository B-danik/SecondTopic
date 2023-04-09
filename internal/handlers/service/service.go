package service

import (
	repository "github.com/B-danik/SecondTopic/internal/database/postgre"
	users "github.com/B-danik/SecondTopic/pkg/users"
	"github.com/B-danik/SecondTopic/todo"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: users.NewAuthService(repo.Authorization),
	}
}
