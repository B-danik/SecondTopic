package servers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/B-danik/SecondTopic/config"
	"github.com/B-danik/SecondTopic/internal/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Server struct {
	cfg     *config.Config
	App     *echo.Echo
	handler *handlers.Manager
}

func NewServer(cfg *config.Config, h *handlers.Manager) *Server {
	return &Server{cfg: cfg, handler: h}
}

func (s *Server) StartServer(c context.Context) error {
	log.Println("Start Server...")

	s.App = s.EnableCors()
	s.SetupRoutes()

	log.Println("Connect to port: ", fmt.Sprint(":", s.cfg.Port))
	if err := s.App.Start(fmt.Sprint(":", s.cfg.Port)); err != nil && err != http.ErrServerClosed {
		log.Fatal("listen:@{err}\n")
	}
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
