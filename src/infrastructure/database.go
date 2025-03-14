package infrastructure

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func NewPostgresDB() *sql.DB {
	dsn := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL", err)
	}
	return db
}
