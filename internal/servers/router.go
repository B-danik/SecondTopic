package servers

import (
	"net/http"

	"github.com/labstack/echo"
)

func (s *Server) SetupRoutes() {
	v1 := s.App.Group("/api/v1")
	s.App.GET("/ready", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
	s.App.GET("/hello", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
	v1.GET("/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello world")
	})
}
