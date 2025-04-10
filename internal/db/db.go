package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"simplegram/internal/models"

	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

func InitDB() (*DB, error) {
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

	return &DB{db}, nil
}

func (db *DB) CloseDB() {
	if db != nil && db.DB != nil {
		err := db.DB.Close()
		if err != nil {
			log.Fatal("Failed to close database connection:", err)
		}
	}
}

func (db *DB) InsertUser(username, hashedPassword string) (int, error) {
	query := `INSERT INTO users (username, hashed_password) VALUES ($1, $2) RETURNING id`
	var userID int
	err := db.QueryRow(query, username, hashedPassword).Scan(&userID)
	return userID, err
}

func (db *DB) GetUserByUsername(username string) (*models.User, error) {
	query := `SELECT id, username, hashed_password FROM users WHERE username = $1`
	user := &models.User{}
	err := db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.HashedPassword)
	return user, err
}
