package main

import (
	"log"
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
	defer dbConn.CloseDB()

	router := routes.SetupRoutes()

	log.Println("Server started on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server failed:", err)
	}
}
