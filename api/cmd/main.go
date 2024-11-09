package main

import (
	"log"
	"os"

	"github.com/ossipesonen/go-traffic-lights/internal/config"
	"github.com/ossipesonen/go-traffic-lights/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	logger.Println("Server starting")

	config := config.New()
	server.New(config, logger)

	logger.Println("Graceful shutdown complete.")
}
