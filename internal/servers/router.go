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
	v1 := s.App.Group("/api/v1")

	s.App.GET("/users", func(c echo.Context) error {

		return c.NoContent(http.StatusOK)
	})

	s.App.GET("/hello", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	s.App.POST("/login", s.loginUser)

	s.App.POST("/auth", s.authUser)

	v1.GET("/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello world")
	})
}

func (s *Server) loginUser(c echo.Context) error {
	var user todo.User
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
	// id, err := service.IAuthorization.CreateUser(user{})
	// if err != nil {
	// 	return nil
	// }
	// c.JSON(http.StatusOK, map[string]interface{}{
	// 	"id": id,
	// })
	return c.NoContent(http.StatusOK)
}

func (s *Server) authUser(c echo.Context) error {
	var user todo.User
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(b, &user)
	if err != nil {
		return nil
	}
	log.Print(user)

	id, err := s.services.Authorization.CreateUser(user)
	if err != nil {
		return nil
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
	return c.NoContent(http.StatusOK)
}
