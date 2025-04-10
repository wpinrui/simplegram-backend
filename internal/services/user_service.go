package services

import (
	"database/sql"
	"simplegram/internal/errors"
	"simplegram/internal/models"

	"golang.org/x/crypto/bcrypt"
)

func NewUserService(db DBInterface, utilities JwtUtilityInterface, errors ErrorInterface) *UserService {
	return &UserService{
		db:        db,
		utilities: utilities,
		errors:    errors,
	}
}

func (us *UserService) CreateUser(username string, password string) (string, error) {
	hashedPassword, err := us.utilities.HashPassword(password)
	if err != nil {
		return "", err
	}

	user := &models.User{
		Username:       username,
		HashedPassword: string(hashedPassword),
	}

	userID, err := us.db.InsertUser(user.Username, user.HashedPassword)
	if err != nil {
		if us.errors.IsUniqueViolation(err) {
			return "", errors.ErrUsernameExists
		}
		return "", err
	}

	user.ID = userID
	token, err := us.utilities.GenerateJwt(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (us *UserService) Login(username string, password string) (string, error) {
	user, err := us.db.GetUserByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.ErrInvalidCredentials
		}
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	if err != nil {
		return "", errors.ErrInvalidCredentials
	}

	token, err := us.utilities.GenerateJwt(user)
	if err != nil {
		return "", err
	}

	return token, nil
}
