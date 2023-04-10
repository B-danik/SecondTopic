package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/B-danik/SecondTopic/todo"
	"github.com/labstack/echo"
)

func (m *Manager) SignUp(c echo.Context) error {
	var user todo.User
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	err = json.Unmarshal(b, &user)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	id, err := m.srv.CreateUser(user)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (m *Manager) SignIn(c echo.Context) error {
	var user todo.User
	if err := c.Bind(&user); err != nil {
		return err
	}
	token, err := m.srv.Authorization.GenerateToken(user.Email, user.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
