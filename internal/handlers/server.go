package handlers

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Server struct {
	App *echo.Echo
}

func NewServer() {

}

func (s *Server) StartServer() error {
	s.App = s.EnableCors()
	return nil
}

func (s *Server) EnableCors() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))
	return e
}
