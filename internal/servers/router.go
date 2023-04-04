package servers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/B-danik/SecondTopic/pkg/users"
	"github.com/labstack/echo"
)

type logUser struct {
	Email    string `json:"email`
	Password string `json:password`
}

func (s *Server) SetupRoutes() {
	v1 := s.App.Group("/api/v1")

	s.App.GET("/users", func(c echo.Context) error {
		user, err := users.New()
		if err != nil {
			return nil
		}
		read, err := user.Get(c)
		log.Print(read)
		return c.NoContent(http.StatusOK)
	})

	s.App.GET("/hello", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	s.App.POST("/login", loginUser)

	s.App.POST("/auth", authUser)

	v1.GET("/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello world")
	})
}

func loginUser(c echo.Context) error {
	user := logUser{}
	defer c.Request().Body.Close()
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(b, &user)
	if err != nil {
		return nil
	}
	log.Print(user)
	return c.NoContent(http.StatusOK)
}

func authUser(c echo.Context) error {
	user, err := users.New()
	if err != nil {
		return nil
	}
	user.Create(c)
	return c.NoContent(http.StatusOK)
}
