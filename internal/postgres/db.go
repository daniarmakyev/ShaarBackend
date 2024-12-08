package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"myfirstBack/config"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	cfg := config.LoadConfig()

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	log.Println("Successfully connected to the database")
	return db, nil
}

func CloseDB(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Fatalf("Error closing database connection: %v", err)
	}
	log.Println("Database connection closed")
}
