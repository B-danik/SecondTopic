package users

import (
	"encoding/json"
	"io/ioutil"

	"github.com/B-danik/SecondTopic/internal/database/CRUD/sql"
	"github.com/labstack/echo"
)

type authUsers struct {
	Name     string `json:name`
	LastName string `json:lastname`
	Email    string `json:email`
	Password string `json:password`
}

type BookService struct {
}

func New() (*BookService, error) {
	return &BookService{}, nil
}

func (b *BookService) Get(c echo.Context) (string, error) {

	users, err := sql.ReadFile("xxx@mail.ru")
	if err != nil {
		return "", nil
	}
	return users, nil
}

func (b *BookService) Create(c echo.Context) error {
	user := authUsers{}
	defer c.Request().Body.Close()
	read, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(read, &user)
	if err != nil {
		return nil
	}
	sql := sql.NewFile(user.Name, user.LastName, user.Email, user.Password)
	sql.CreateFile()
	return nil
}
