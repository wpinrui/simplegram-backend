package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() (*sql.DB, error) {
	connStr := "user=postgres dbname=simplegram sslmode=disable password=password"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Ensure the connection is available
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// Assign to the global DB variable for reuse across services
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
