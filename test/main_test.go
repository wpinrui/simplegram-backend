package test

import (
	"database/sql"
	"log"
	"os"
	"simplegram/internal/db"
	"testing"
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
