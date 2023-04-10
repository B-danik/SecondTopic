package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func (m *Manager) CreateList(c echo.Context) error {
	id := c.Request().Header.Get(authorizationHeader)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
