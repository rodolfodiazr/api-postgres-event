package databases

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Cannot load the environment variables: ", err)
	}

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return db, err
	}

	if err := db.Ping(); err != nil {
		return db, err
	}

	log.Println("Connected to PostgreSQL!")
	return db, nil
}
