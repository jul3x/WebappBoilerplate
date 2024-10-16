package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jul3x/WebappBoilerplate/models"
	"github.com/jul3x/WebappBoilerplate/tests"
	"github.com/stretchr/testify/assert"
)

func TestRegister_ValidUser(t *testing.T) {
    // Initialize a test database
    db := tests.InitializeTestDB(t)

    // Set Gin to Test Mode
    gin.SetMode(gin.TestMode)

    // Create a new Gin router
    router := gin.Default()

    // Register routes for testing
    router.POST("/api/v1/auth/register", Register(db))
    router.POST("/api/v1/auth/login", Login(db))

    // Step 1: Test Registration

    // Mock request body for registration
    requestBody, _ := json.Marshal(map[string]string{
        "username": "testuser",
        "email":    "testuser@example.com",
        "password": "testpassword",
    })

    // Create a new HTTP POST request to the registration endpoint
    req, err := http.NewRequest(http.MethodPost, "/api/v1/auth/register", bytes.NewBuffer(requestBody))
    if err != nil {
        t.Fatalf("Could not create request: %v", err)
    }
    req.Header.Set("Content-Type", "application/json")

    // Create a ResponseRecorder to record the response
    rr := httptest.NewRecorder()

    // Serve the request to the router
    router.ServeHTTP(rr, req)

    // Assert the status code
    assert.Equal(t, http.StatusCreated, rr.Code)

    // Step 2: Test Login

    // Mock request body for login
    loginRequestBody, _ := json.Marshal(map[string]string{
        "email":    "testuser@example.com",
        "password": "testpassword",
    })

    // Create a new HTTP POST request to the login endpoint
    req, err = http.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewBuffer(loginRequestBody))
    if err != nil {
        t.Fatalf("Could not create request: %v", err)
    }
    req.Header.Set("Content-Type", "application/json")

    // Reset the ResponseRecorder to record the login response
    rr = httptest.NewRecorder()

    // Serve the request to the router
    router.ServeHTTP(rr, req)

    // Assert the status code for the login response
    assert.Equal(t, http.StatusOK, rr.Code)

    // Parse the JSON response body
    var response map[string]interface{}
    err = json.Unmarshal(rr.Body.Bytes(), &response)
    if err != nil {
        t.Fatalf("Could not parse JSON response: %v", err)
    }

    // Assert that the token is present in the response
    assert.NotNil(t, response["token"], "Expected token in the response")
}

func TestRegister_DuplicateEmail(t *testing.T) {
    // Initialize a test database
    db := tests.InitializeTestDB(t)

    // Create an existing user to simulate duplicate email
    db.Create(&models.User{
        Username: "existinguser",
        Email:    "existinguser@example.com",
        Password: "hashedpassword", // Ideally this should be hashed
    })

    // Set Gin to Test Mode
    gin.SetMode(gin.TestMode)

    // Create a new Gin router
    router := gin.Default()

    // Register routes for testing
    router.POST("/api/v1/auth/register", Register(db))

    // Mock request body for registration with duplicate email
    requestBody, _ := json.Marshal(map[string]string{
        "username": "newuser",
        "email":    "existinguser@example.com", // Duplicate email
        "password": "newpassword",
    })

    // Create a new HTTP POST request to the registration endpoint
    req, err := http.NewRequest(http.MethodPost, "/api/v1/auth/register", bytes.NewBuffer(requestBody))
    if err != nil {
        t.Fatalf("Could not create request: %v", err)
    }
    req.Header.Set("Content-Type", "application/json")

    // Create a ResponseRecorder to record the response
    rr := httptest.NewRecorder()

    // Serve the request to the router
    router.ServeHTTP(rr, req)

    // Assert the status code
    assert.Equal(t, http.StatusConflict, rr.Code)

    // Assert the expected error message
    var response map[string]interface{}
    err = json.Unmarshal(rr.Body.Bytes(), &response)
    if err != nil {
        t.Fatalf("Could not parse JSON response: %v", err)
    }

    expectedMessage := "Email already in use"
    assert.Equal(t, expectedMessage, response["error"], "Expected error message does not match")
}


func TestLogin_InvalidUser(t *testing.T) {
    // Initialize a test database
    db := tests.InitializeTestDB(t)

    // Set Gin to Test Mode
    gin.SetMode(gin.TestMode)

    // Create a new Gin router
    router := gin.Default()

    // Register routes for testing
    router.POST("/api/v1/auth/login", Login(db))

    // Mock request body for logging in with invalid credentials
    requestBody, _ := json.Marshal(map[string]string{
        "email":    "nonexistent@example.com", // Non-existing user
        "password": "invalidpassword",          // Invalid password
    })

    // Create a new HTTP POST request to the login endpoint
    req, err := http.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewBuffer(requestBody))
    if err != nil {
        t.Fatalf("Could not create request: %v", err)
    }
    req.Header.Set("Content-Type", "application/json")

    // Create a ResponseRecorder to record the response
    rr := httptest.NewRecorder()

    // Serve the request to the router
    router.ServeHTTP(rr, req)

    // Assert the status code
    assert.Equal(t, http.StatusUnauthorized, rr.Code)

    // Assert the expected error message
    var response map[string]interface{}
    err = json.Unmarshal(rr.Body.Bytes(), &response)
    if err != nil {
        t.Fatalf("Could not parse JSON response: %v", err)
    }

    expectedMessage := "Invalid credentials"
    assert.Equal(t, expectedMessage, response["error"], "Expected error message does not match")
}
