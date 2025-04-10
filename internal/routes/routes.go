package routes

import (
	"log"
	"simplegram/internal/controllers"
	"simplegram/internal/db"
	internalErrors "simplegram/internal/errors"
	"simplegram/internal/services"
	"simplegram/internal/utilities"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	// Initialize the database connection
	dbConn, err := db.InitDB()
	if err != nil {
		log.Fatal("Failed to initialize the database: ", err)
	}

	userService := services.NewUserService(dbConn, utilities.NewUtilities(), internalErrors.NewError())
	userController := controllers.NewUserController(userService)

	router := gin.Default()

	router.POST("/users", userController.CreateUser)
	router.POST("/users/login", userController.Login)

	return router
}
