package handlers

import (
	"github.com/B-danik/SecondTopic/pkg/service"
)

type Manager struct {
	srv *service.Service
}

func NewManager(srv *service.Service) *Manager {
	return &Manager{
		srv: srv,
	}
}
