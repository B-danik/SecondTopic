package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/B-danik/SecondTopic/todo"
	"github.com/labstack/echo"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Manager) signUp(c echo.Context) error {
	var user todo.User
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	err = json.Unmarshal(b, &user)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	id, err := h.srv.CreateUser(user)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
	return c.NoContent(http.StatusOK)
}

func (m *Manager) signIn(c echo.Context) error {
	var user todo.User
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	err = json.Unmarshal(b, &user)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	token, err := m.srv.Authorization.GenerateToken(user.Email, user.Password)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
	return c.NoContent(http.StatusOK)
}
