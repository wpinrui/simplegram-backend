package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
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

		token, err := services.CreateUser(dbConn, userRequest.Username, userRequest.Password)
		if err != nil {
			if errors.Is(err, services.ErrUsernameExists) {
				http.Error(w, "Username already exists", http.StatusUnprocessableEntity)
				return
			}
			http.Error(w, "Error creating user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		response := map[string]string{"token": token}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
			return
		}
	}
}
