package handlers

import (
	"github.com/B-danik/SecondTopic/config"
	service "github.com/B-danik/SecondTopic/internal/handlers/service"
)

type Manager struct {
	srv *service.Service
}

func NewManager(con *config.Config, srv *service.Service) *Manager {
	return &Manager{
		srv: srv,
	}
}
