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

func (m *Manager) MessageHandler(message string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(message))
	})
}

func (m *Manager) UserIdentity(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
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
			return c.JSON(http.StatusOK, err.Error())
		}
		c.Set(userCtx, userId)
		return next(c)
	}
}
