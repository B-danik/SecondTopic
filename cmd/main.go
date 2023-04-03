package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/B-danik/SecondTopic/config"
)

func main() {
	log.Println("Stop application...", run())
}

func run() error {
	log.Println("Start application...")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	grecefulShurdown(cancel)

	config, err := config.Read()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println(config)

	return ctx.Err()
}

func grecefulShurdown(c context.CancelFunc) {
	gS := make(chan os.Signal, 1)
	signal.Notify(gS, os.Interrupt)
	go func() {
		log.Print(<-gS)
		c()
	}()
}
