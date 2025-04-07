package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"simplegram/internal/models"
	"simplegram/internal/services"
)

func CreateUser(dbConn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userRequest struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		hashedPassword := "hashedPassword"

		user := models.User{
			Username:       userRequest.Username,
			HashedPassword: hashedPassword,
		}

		if err := services.CreateUser(dbConn, &user); err != nil {
			http.Error(w, "Error creating user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	}
}
