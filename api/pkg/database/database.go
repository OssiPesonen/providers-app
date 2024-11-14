package database

import (
	"fmt"
	"log"

	_ "github.com/joho/godotenv/autoload"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
)

type Database interface {
	Health() map[string]string
	Close() error
	// Returns the connection handle to enable access to database
	Handle() db.Session
}

type connection struct {
	db     db.Session
	logger *log.Logger
}

type DBConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
	Schema   string
}

var (
	dbInstance *connection
)

func New(config *DBConfig, logger *log.Logger) Database {
	// Reuse Connection
	if dbInstance != nil {
		return dbInstance
	}

	var settings = postgresql.ConnectionURL{
		Host:     config.Host,
		Database: config.Database,
		User:     config.Username,
		Password: config.Password,
	}

	db, err := postgresql.Open(settings)

	if err != nil {
		log.Fatal(err)
	}

	dbInstance = &connection{
		db:     db,
		logger: logger,
	}

	return dbInstance
}

func (conn *connection) Handle() db.Session {
	return conn.db
}

// Health checks the health of the database connection by pinging the database.
// It returns a map with keys indicating various health statistics.
func (conn *connection) Health() map[string]string {
	info := make(map[string]string)

	// Ping the database
	err := conn.db.Ping()

	if err != nil {
		info["status"] = "down"
		info["error"] = fmt.Sprintf("db down: %v", err)
		conn.logger.Fatalf("db down: %v", err) // Log the error and terminate the program
		return info
	}

	// Database is up
	info["status"] = "up"
	return info
}

// Close closes the database connection.
// If the connection is successfully closed, it returns nil.
// If an error occurs while closing the connection, it returns the error.
func (conn *connection) Close() error {
	conn.logger.Printf("Disconnected from database")
	return conn.db.Close()
}
