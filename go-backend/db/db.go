package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func Migrate() {
	query := `
	CREATE TABLE IF NOT EXISTS notes (
		id SERIAL PRIMARY KEY,
		text TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatalf("❌ Failed to migrate database: %v", err)
	}

	log.Println("✅ Database migrated (notes table ready)")
}

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  No .env file found, using system environment variables")
	}

	connStr := os.Getenv("DATABASE_URL_LOCAL")
	if connStr == "" {
		log.Fatal("❌ DATABASE_URL not set in environment")
	}

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("❌ Failed to open database:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("❌ Database unreachable:", err)
	}

	log.Println("✅ Connected to PostgreSQL")
}
