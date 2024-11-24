package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/ossipesonen/providers-app/internal/config"
	"github.com/ossipesonen/providers-app/internal/server"
)

func main() {
	ctx := context.Background()
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	logger.Println("starting...")

	config := config.New()
	s := server.New(config, logger)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for range c {
			logger.Println("shutting down server...")

			s.GracefulStop()
			<-ctx.Done()
		}
	}()
}
