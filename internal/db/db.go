package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() (*sql.DB, error) {

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbSSLMode := os.Getenv("DB_SSLMODE")
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbName, dbPassword, dbSSLMode,
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	DB = db
	return db, nil
}

func CloseDB() {
	if DB != nil {
		err := DB.Close()
		if err != nil {
			log.Fatal("Failed to close database connection:", err)
		}
	}
}
