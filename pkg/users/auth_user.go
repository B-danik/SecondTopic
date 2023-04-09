package users

import (
	"crypto/sha1"
	"fmt"

	repository "github.com/B-danik/SecondTopic/internal/database/postgre"
	"github.com/B-danik/SecondTopic/todo"
)

const salt = "tedsa1235das"

type AuthService struct {
	repo repository.IAuthorization
}

func NewAuthService(repo repository.IAuthorization) *AuthService {
	return &AuthService{repo: repo}
}

func (a *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return a.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
