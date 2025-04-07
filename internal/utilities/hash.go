package utilities

import (
	"crypto/sha256"
	"encoding/hex"
)

// HashPassword hashes a password using SHA-256 and returns the hash as a string.
func HashPassword(password string) string {
	h := sha256.New()
	h.Write([]byte(password))
	return hex.EncodeToString(h.Sum(nil))
}
