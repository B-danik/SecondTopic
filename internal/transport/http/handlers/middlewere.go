package handlers

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (m *Manager) UserIdentity(c echo.Context) error {
	header := c.Request().Header.Get(authorizationHeader)

	if header == "" {
		return c.JSON(http.StatusOK, "empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return c.JSON(http.StatusOK, "invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return c.JSON(http.StatusOK, "token is empty")
	}

	userId, err := m.srv.Authorization.ParseToken(headerParts[1])
	if err != nil {
		c.JSON(http.StatusOK, err.Error())
	}

	c.Set(userCtx, userId)

	return m.CreateList(c)
}
