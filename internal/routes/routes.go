package routes

import (
	"database/sql"
	"simplegram/internal/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes(dbConn *sql.DB) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/users", controllers.CreateUser(dbConn)).Methods("POST")
	return router
}
