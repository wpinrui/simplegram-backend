package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"simplegram/internal/controllers"
	"simplegram/internal/db"
	"simplegram/internal/services"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserController(t *testing.T) {
	// Setup test database and server
	testDB, err := db.InitDB()
	if err != nil {
		t.Fatal("Failed to connect to test DB:", err)
	}
	defer testDB.Close()

	// Create a new HTTP server with the user controller
	router := mux.NewRouter()
	router.HandleFunc("/users", controllers.CreateUser(testDB)).Methods("POST")

	// Test case: Successful user creation
	t.Run("Success", func(t *testing.T) {
		username := "testuser_create"
		password := "securepass"

		cleanupUser(username)

		userRequest := map[string]string{"username": username, "password": password}
		body, _ := json.Marshal(userRequest)

		req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)

		var response map[string]string
		json.Unmarshal(recorder.Body.Bytes(), &response)
		assert.NotEmpty(t, response["token"])

		cleanupUser(username)
	})

	// Test case: Duplicate username
	t.Run("DuplicateUsername", func(t *testing.T) {
		username := "testuser_dup"
		password := "securepass"

		cleanupUser(username)

		userRequest := map[string]string{"username": username, "password": password}
		body, _ := json.Marshal(userRequest)

		fmt.Println("Request Body:", string(body)) // Log request body

		// First request
		req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
		if err != nil {
			t.Fatalf("Error creating request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")
		fmt.Println("First request:", req.Method, req.URL, req.Header, string(body)) // Log first request details

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		fmt.Println("First response status:", recorder.Code) // Log first response status
		assert.Equal(t, http.StatusOK, recorder.Code)

		// Second request, recreate the request with the same body
		req, err = http.NewRequest("POST", "/users", bytes.NewBuffer(body))
		if err != nil {
			t.Fatalf("Error creating request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		recorder = httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusUnprocessableEntity, recorder.Code)

		cleanupUser(username)
	})

}

func TestLoginController(t *testing.T) {
	// Setup test database and server
	testDB, err := db.InitDB()
	if err != nil {
		t.Fatal("Failed to connect to test DB:", err)
	}
	defer testDB.Close()

	// Create a new HTTP server with the user controller
	router := mux.NewRouter()
	router.HandleFunc("/login", controllers.Login(testDB)).Methods("POST")

	// Test case: Successful login
	t.Run("Success", func(t *testing.T) {
		username := "testuser_login"
		password := "securepass"

		cleanupUser(username)
		_, _ = services.CreateUser(testDB, username, password)

		userRequest := map[string]string{"username": username, "password": password}
		body, _ := json.Marshal(userRequest)

		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)

		var response map[string]string
		json.Unmarshal(recorder.Body.Bytes(), &response)
		assert.NotEmpty(t, response["token"])

		cleanupUser(username)
	})
}
