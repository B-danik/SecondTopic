package servers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/B-danik/SecondTopic/todo"
	"github.com/labstack/echo"
)

func (s *Server) SetupRoutes() {
	auth := s.App.Group("/auth")
	{
		auth.POST("/sign-up", s.signUp)
		auth.POST("/sign-in", s.signIn)
	}
}

func (s *Server) signUp(c echo.Context) error {
	var user todo.User
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	err = json.Unmarshal(b, &user)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	log.Print(user)

	id, err := s.services.Authorization.CreateUser(user)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
	return c.NoContent(http.StatusOK)
}

func (s *Server) signIn(c echo.Context) error {
	var user todo.User
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(b, &user)
	if err != nil {
		return nil
	}
	token, err := s.services.Authorization.GenerateToken(user.Email, user.Password)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
	return c.NoContent(http.StatusOK)
}
