package main

import (
	"context"
	"log"
	"os"
	"os/signal"
)

func main() {
	config := cfg.LoadAndStoreConfig()

	ctx, cancel := context.WithCancel(context.Background())

	ch := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	server := app.NewServer(ctx, config)

	go func() {
		osCall := <-ch
		log.Printf("system call: %+v", osCall)

		server.Shutdown()
		cancel()
	}()
	server.Serve()
}
