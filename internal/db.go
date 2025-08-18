package internal

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
	var db *sql.DB
	var err error

	connStr := "postgres://[db-name]:@localhost:5432/eventsdb?sslmode=disable" // Move to environment variables
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to open DB: %v", err)
		return db, err
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to connect: %v", err)
		return db, err
	}

	fmt.Println("Connected to PostgreSQL!")
	return db, nil
}
