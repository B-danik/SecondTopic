package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/B-danik/SecondTopic/config"
	repository "github.com/B-danik/SecondTopic/internal/database/postgre"
	database "github.com/B-danik/SecondTopic/internal/database/postgre/connect"
	servers "github.com/B-danik/SecondTopic/internal/transport/http"
	"github.com/B-danik/SecondTopic/internal/transport/http/handlers"
	"github.com/B-danik/SecondTopic/pkg/service"
)

func main() {
	log.Println("Stop application", run())
}

func run() error {
	log.Println("Start application...")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	grecefulShurdown(cancel)
	con, err := config.Read()
	if err != nil {
		log.Fatal(err.Error())
	}

	db, err := database.NewPostgresDB(con)
	if err != nil {
		log.Fatal(err.Error())
	}

	repos := repository.NewRepository(db)

	svc := service.NewService(repos)

	manager := handlers.NewManager(svc)

	server := servers.NewServer(con, manager, svc)

	return server.StartServer(ctx)
}

func grecefulShurdown(c context.CancelFunc) {
	gS := make(chan os.Signal, 1)
	signal.Notify(gS, os.Interrupt)
	go func() {
		log.Print(<-gS)
		c()
	}()
}
