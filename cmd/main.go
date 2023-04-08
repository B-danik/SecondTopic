package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/B-danik/SecondTopic/config"
	database "github.com/B-danik/SecondTopic/internal/database/postgre/connect"
	handlers "github.com/B-danik/SecondTopic/internal/handlers/manager"
	service "github.com/B-danik/SecondTopic/internal/handlers/service"
	"github.com/B-danik/SecondTopic/internal/servers"
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
	database.ConnectDB(con)

	svc, err := service.NewService()

	handler := handlers.NewManager(con, svc)

	server := servers.NewServer(con, handler)

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
