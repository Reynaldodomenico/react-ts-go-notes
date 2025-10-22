package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		connStr = "postgres://credentials_here@localhost:5432/notesdb?sslmode=disable"
	}

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("❌ Failed to open database:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("❌ Database unreachable:", err)
	}

	log.Println("✅ Connected to PostgreSQL")
}
