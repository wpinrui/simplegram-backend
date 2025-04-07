package auth

import (
	"fmt"
	"simplegram/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwt(user *models.User) (string, error) {
	var (
		key []byte
		t   *jwt.Token
		s   string
		err error
	)
	key = []byte("KEY")
	t = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		"iat":      jwt.NewNumericDate(time.Now()),
		"nbf":      jwt.NewNumericDate(time.Now()),
		"sub":      user.ID,
		"iss":      "simplegram",
		"aud":      "simplegram",
	})
	s, err = t.SignedString(key)
	if err != nil {
		return "", fmt.Errorf("error signing JWT: %w", err)
	}
	return s, nil
}
