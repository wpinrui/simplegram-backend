package main

import (
	"log"
	"net/http"
	"simplegram/internal/db"
	"simplegram/internal/routes"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbConn, err := db.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.CloseDB()

	router := routes.SetupRoutes(dbConn)

	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("Server failed:", err)
	}
}
