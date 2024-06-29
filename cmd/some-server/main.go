package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"some-httpserver/internal/cfg"
)

func main() {
	config := cfg.LoadAndStoreConfig()

	ctx, cancel := context.WithCancel(context.Background())

	ch := make(chan os.Signal, 1) //grace full shutdown

	signal.Notify(ch, os.Interrupt)

	server := app.NewServer(ctx, config)

	go func() {
		osCall := <-ch
		log.Printf("system call: %+v", osCall)

		server.Shutdown()
		cancel()
	}()
	server.Serve()
}
