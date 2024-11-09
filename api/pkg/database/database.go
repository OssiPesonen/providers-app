package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
)

type Database interface {
	Health() map[string]string
	Close() error
	// Returns the connection handle to enable access to database
	Handle() *sql.DB
}

type connection struct {
	db     *sql.DB
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

	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
		config.Schema,
	)

	db, err := sql.Open("pgx", connStr)

	if err != nil {
		log.Fatal(err)
	}

	dbInstance = &connection{
		db:     db,
		logger: logger,
	}

	return dbInstance
}

func (conn *connection) Handle() *sql.DB {
	return conn.db
}

// Health checks the health of the database connection by pinging the database.
// It returns a map with keys indicating various health statistics.
func (conn *connection) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	info := make(map[string]string)

	// Ping the database
	err := conn.db.PingContext(ctx)

	if err != nil {
		info["status"] = "down"
		info["error"] = fmt.Sprintf("db down: %v", err)
		conn.logger.Fatalf("db down: %v", err) // Log the error and terminate the program
		return info
	}

	// Database is up
	info["status"] = "up"

	isDebug := os.Getenv("APP_DEBUG")

	// Don't expose details unless APP_DEBUG enabled
	if isDebug != "" {
		// Get database stats (like open connections, in use, idle, etc.)
		dbStats := conn.db.Stats()
		info["open_connections"] = strconv.Itoa(dbStats.OpenConnections)
		info["in_use"] = strconv.Itoa(dbStats.InUse)
		info["idle"] = strconv.Itoa(dbStats.Idle)
		info["wait_count"] = strconv.FormatInt(dbStats.WaitCount, 10)
		info["wait_duration"] = dbStats.WaitDuration.String()
		info["max_idle_closed"] = strconv.FormatInt(dbStats.MaxIdleClosed, 10)
		info["max_lifetime_closed"] = strconv.FormatInt(dbStats.MaxLifetimeClosed, 10)

		// Evaluate stats to provide a health message
		if dbStats.OpenConnections > 40 { // Assuming 50 is the max for this example
			info["message"] = "The database is experiencing heavy load."
		}

		if dbStats.WaitCount > 1000 {
			info["message"] = "The database has a high number of wait events, indicating potential bottlenecks."
		}

		if dbStats.MaxIdleClosed > int64(dbStats.OpenConnections)/2 {
			info["message"] = "Many idle connections are being closed, consider revising the connection pool settings."
		}

		if dbStats.MaxLifetimeClosed > int64(dbStats.OpenConnections)/2 {
			info["message"] = "Many connections are being closed due to max lifetime, consider increasing max lifetime or revising the connection usage pattern."
		}
	}

	return info
}

// Close closes the database connection.
// If the connection is successfully closed, it returns nil.
// If an error occurs while closing the connection, it returns the error.
func (conn *connection) Close() error {
	conn.logger.Printf("Disconnected from database")
	return conn.db.Close()
}
