package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jul3x/WebappBoilerplate/models"
	"github.com/jul3x/WebappBoilerplate/tests"
)

func TestRegister_ValidUser(t *testing.T) {
	db := tests.InitializeTestDB(t)

	requestBody, _ := json.Marshal(map[string]string{
		"username": "testuser",
		"email":    "testuser@example.com",
		"password": "testpassword",
	})

	req, err := http.NewRequest("POST", "/api/v1/auth/register", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := Register(db)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Expected status code %v, got %v", http.StatusCreated, status)
	}

        requestBody, err = json.Marshal(map[string]string{
		"email":    "testuser@example.com",
		"password": "testpassword",
        })

        req, err = http.NewRequest("POST", "/api/v1/auth/login", bytes.NewBuffer(requestBody))
        if err != nil {
                t.Fatalf("Could not create request: %v", err)
        }
        req.Header.Set("Content-Type", "application/json")

        rr = httptest.NewRecorder()
        handler = Login(db)
        handler.ServeHTTP(rr, req)

        if status := rr.Code; status != http.StatusOK {
                t.Errorf("Expected status code %v, got %v", http.StatusOK, status)
        }

        var response map[string]interface{}
        json.Unmarshal(rr.Body.Bytes(), &response)
        if response["token"] == nil {
                t.Errorf("Expected token in response, got nil")
        }
}

func TestRegister_DuplicateEmail(t *testing.T) {
	db := tests.InitializeTestDB(t)

	db.Create(&models.User{
		Username: "existinguser",
		Email:    "existinguser@example.com",
		Password: "hashedpassword",
	})

	requestBody, _ := json.Marshal(map[string]string{
		"username": "newuser",
		"email":    "existinguser@example.com", // Duplicate email
		"password": "newpassword",
	})

	req, err := http.NewRequest("POST", "/api/v1/auth/register", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := Register(db)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusConflict {
		t.Errorf("Expected status code %v, got %v", http.StatusConflict, status)
	}

	expected := "Email already in use\n"
	if rr.Body.String() != expected {
		t.Errorf("Expected error message '%v', got '%v'", expected, rr.Body.String())
	}
}

func TestLogin_InvalidUser(t *testing.T) {
        db := tests.InitializeTestDB(t)

        requestBody, _ := json.Marshal(map[string]string{
                "email":    "nonexistent@example.com",
                "password": "invalidpassword",
        })

        req, err := http.NewRequest("POST", "/api/v1/auth/login", bytes.NewBuffer(requestBody))
        if err != nil {
                t.Fatalf("Could not create request: %v", err)
        }
        req.Header.Set("Content-Type", "application/json")

        rr := httptest.NewRecorder()
        handler := Login(db)
        handler.ServeHTTP(rr, req)

        if status := rr.Code; status != http.StatusUnauthorized {
                t.Errorf("Expected status code %v, got %v", http.StatusUnauthorized, status)
        }

        expected := "Invalid credentials\n"
        if rr.Body.String() != expected {
                t.Errorf("Expected error message '%v', got '%v'", expected, rr.Body.String())
        }
}
