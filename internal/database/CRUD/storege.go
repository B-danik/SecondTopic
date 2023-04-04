package database

import (
	"context"

	"github.com/B-danik/SecondTopic/config"
)

type ICRUD interface {
	Create()
	Read()
	Update()
	Delete()
}

type CRUD struct {
	context *context.Context
	config  *config.Config
	crud    ICRUD
}

func NewCRUD(ctx context.Context, cfg config.Config) (*CRUD, error) {
	return &CRUD{
		context: &ctx,
		config:  &cfg,
	}, nil
}
