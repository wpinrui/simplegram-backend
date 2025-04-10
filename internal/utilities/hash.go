package utilities

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes the password using bcrypt.
func (u *Utilities) HashPassword(password string) (string, error) {
	// bcrypt.DefaultCost is the default cost for bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
