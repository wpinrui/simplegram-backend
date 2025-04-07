package services

import (
	"database/sql"
	"log"

	"simplegram/internal/models"

	"github.com/lib/pq"
)

func CreateUser(dbConn *sql.DB, user *models.User) error {
	query := `INSERT INTO users (username, hashed_password) VALUES ($1, $2) RETURNING id;`
	err := dbConn.QueryRow(query, user.Username, user.HashedPassword).Scan(&user.ID)
	if err != nil {
		if isUniqueViolation(err) {
			return ErrUsernameExists
		}
		log.Println("Error inserting user:", err)
		return err
	}
	return nil
}

func isUniqueViolation(err error) bool {
	pqErr, ok := err.(*pq.Error)
	return ok && pqErr.Code == "23505"
}
