package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/m7md2323/JoTech-Directory/internal/database"
	"github.com/m7md2323/JoTech-Directory/internal/models"
	"gorm.io/gorm"
)

func setupTestDB() {
	// Initialize an in-memory database for testing
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to test database")
	}

	database.DB = db
	database.DB.AutoMigrate(&models.Company{})
	database.DB.AutoMigrate(&models.Tag{})
	database.DB.AutoMigrate(&models.Location{})
}

func TestHomeHandler(t *testing.T) {
	setupTestDB()

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", Home)

	mux.ServeHTTP(rr, req)

	// We expect the home page to return a 200 OK status
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
