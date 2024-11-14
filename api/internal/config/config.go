package config

import (
	"log"
	"os"
	"strconv"
)

type (
	Config struct {
		DB     *DB
		Server *Server
		Auth   *Auth
	}

	Auth struct {
		JWTSecret string
	}

	Server struct {
		Port int
	}

	DB struct {
		Database string
		Host     string
		Port     string
		Username string
		Password string
		Schema   string
	}
)

func New() *Config {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("Unable to get port from env vars")
	}

	auth := &Auth{JWTSecret: os.Getenv("APP_JWT_SECRET")}

	server := &Server{Port: port}

	db := &DB{
		Database: os.Getenv("APP_DB_DATABASE"),
		Password: os.Getenv("APP_DB_PASSWORD"),
		Username: os.Getenv("APP_DB_USERNAME"),
		Port:     os.Getenv("APP_DB_PORT"),
		Host:     os.Getenv("APP_DB_HOST"),
	}

	return &Config{
		DB:     db,
		Server: server,
		Auth:   auth,
	}
}
