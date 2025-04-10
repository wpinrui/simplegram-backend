package services

import "simplegram/internal/models"

type DBInterface interface {
	InsertUser(username, hashedPassword string) (int, error)
	GetUserByUsername(username string) (*models.User, error)
}

type JwtUtilityInterface interface {
	GenerateJwt(user *models.User) (string, error)
	HashPassword(password string) (string, error)
}

type ErrorInterface interface {
	IsUniqueViolation(err error) bool
}

type UserService struct {
	db        DBInterface
	utilities JwtUtilityInterface
	errors    ErrorInterface
}
