package handlers

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/golang-jwt/jwt/v4"
    "github.com/jul3x/webapp-boilerplate/middlewares"
    "github.com/jul3x/webapp-boilerplate/tests"
    "os"
    "time"
)

func TestProtected_ValidJWT(t *testing.T) {
    db := tests.InitializeTestDB(t)

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": 1,
        "exp":     time.Now().Add(time.Hour * 1).Unix(),
    })
    tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

    req, err := http.NewRequest("GET", "/api/v1/protected/data", nil)
    if err != nil {
        t.Fatalf("Could not create request: %v", err)
    }
    req.Header.Set("Authorization", "Bearer "+tokenString)

    rr := httptest.NewRecorder()
    handler := middlewares.JwtMiddleware(GetProtectedData(db))
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("Expected status code %v, got %v", http.StatusOK, status)
    }
}

func TestProtected_InvalidJWT(t *testing.T) {
    db := tests.InitializeTestDB(t)

    req, err := http.NewRequest("GET", "/api/v1/protected/data", nil)
    if err != nil {
        t.Fatalf("Could not create request: %v", err)
    }
    req.Header.Set("Authorization", "Bearer invalidtoken")

    rr := httptest.NewRecorder()
    handler := middlewares.JwtMiddleware(GetProtectedData(db))
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusUnauthorized {
        t.Errorf("Expected status code %v, got %v", http.StatusUnauthorized, status)
    }

    expected := "Invalid token\n"
    if rr.Body.String() != expected {
        t.Errorf("Expected error message '%v', got '%v'", expected, rr.Body.String())
    }
}
