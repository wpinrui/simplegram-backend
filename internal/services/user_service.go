package services

import (
	"database/sql"
	"log"
	"simplegram/internal/auth"
	"simplegram/internal/models"
	"simplegram/internal/utilities"

	"github.com/lib/pq"
)

func CreateUser(dbConn *sql.DB, username string, password string) (string, error) {
	hashedPassword := utilities.HashPassword(password)
	user := &models.User{
		Username:       username,
		HashedPassword: hashedPassword,
	}
	query := `INSERT INTO users (username, hashed_password) VALUES ($1, $2) RETURNING id;`
	err := dbConn.QueryRow(query, user.Username, user.HashedPassword).Scan(&user.ID)
	if err != nil {
		if isUniqueViolation(err) {
			return "", ErrUsernameExists
		}
		log.Println("Error inserting user:", err)
		return "", err
	}
	token, err := auth.GenerateJwt(user)
	if err != nil {
		log.Println("Error generating JWT:", err)
		return "", err
	}
	return token, err
}

func Login(dbConn *sql.DB, username string, password string) (string, error) {
	hashedPassword := utilities.HashPassword(password)
	query := `SELECT username, hashed_password FROM users WHERE username = $1 AND hashed_password = $2;`
	user := &models.User{}
	err := dbConn.QueryRow(query, username, hashedPassword).Scan(&user.Username, &user.HashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", ErrInvalidCredentials
		}
		log.Println("Error querying user:", err)
		return "", err
	}
	token, err := auth.GenerateJwt(user)
	if err != nil {
		log.Println("Error generating JWT:", err)
		return "", err
	}
	return token, nil
}

func isUniqueViolation(err error) bool {
	pqErr, ok := err.(*pq.Error)
	return ok && pqErr.Code == "23505"
}
