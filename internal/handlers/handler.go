package handlers

import (
	"github.com/B-danik/SecondTopic/config"
	service "github.com/B-danik/SecondTopic/pkg/handlers"
)

type Manager struct {
	srv *service.Manager
}

func NewManager(con *config.Config, srv *service.Manager) *Manager {
	return &Manager{
		srv: srv,
	}
}
