package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/B-danik/SecondTopic/config"
	handlers "github.com/B-danik/SecondTopic/internal/transport/http/handlers"
	"github.com/B-danik/SecondTopic/pkg/service"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Server struct {
	cfg      *config.Config
	App      *echo.Echo
	handler  *handlers.Manager
	services *service.Service
}

func NewServer(cfg *config.Config, h *handlers.Manager, s *service.Service) *Server {
	return &Server{cfg: cfg, handler: h, services: s}
}

func (s *Server) StartServer(c context.Context) error {
	fmt.Printf("s.handler: %v\n", s.handler)
	log.Println("Start Server...")
	s.App = s.EnableCors()
	s.SetupRoutes()
	go func() {
		log.Println("Connect to port: ", fmt.Sprint(":", s.cfg.Port))
		if err := s.App.Start(fmt.Sprint(":", s.cfg.Port)); err != nil && err != http.ErrServerClosed {
			log.Fatal("Error connect: \n", err.Error())
		}
	}()

	<-c.Done()
	ctxShutdown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()
	if err := s.App.Shutdown(ctxShutdown); err != nil {
		log.Fatal("Server shutdown error: ", err.Error())
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
