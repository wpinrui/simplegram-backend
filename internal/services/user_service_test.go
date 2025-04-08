package services_test

import (
	"database/sql"
	"log"
	"os"
	"simplegram/internal/db"
	"simplegram/internal/services"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = db.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to test DB:", err)
	}
	defer testDB.Close()

	code := m.Run()
	os.Exit(code)
}

func cleanupUser(username string) {
	_, _ = testDB.Exec(`DELETE FROM users WHERE username = $1`, username)
}

func TestCreateUser_Success(t *testing.T) {
	username := "testuser_create"
	password := "securepass"

	cleanupUser(username)

	token, err := services.CreateUser(testDB, username, password)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	cleanupUser(username)
}

func TestCreateUser_DuplicateUsername(t *testing.T) {
	username := "testuser_dup"
	password := "securepass"

	cleanupUser(username)

	_, err := services.CreateUser(testDB, username, password)
	assert.NoError(t, err)

	_, err = services.CreateUser(testDB, username, password)
	assert.ErrorIs(t, err, services.ErrUsernameExists)

	cleanupUser(username)
}

func TestLogin_Success(t *testing.T) {
	username := "testuser_login"
	password := "securepass"

	cleanupUser(username)
	_, _ = services.CreateUser(testDB, username, password)

	token, err := services.Login(testDB, username, password)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	cleanupUser(username)
}

func TestLogin_InvalidCredentials(t *testing.T) {
	username := "nonexistent_user"
	password := "wrongpass"

	token, err := services.Login(testDB, username, password)
	assert.ErrorIs(t, err, services.ErrInvalidCredentials)
	assert.Empty(t, token)
}
