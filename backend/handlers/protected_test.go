package handlers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jul3x/WebappBoilerplate/middlewares"
	"github.com/jul3x/WebappBoilerplate/tests"
	"github.com/stretchr/testify/assert"
)

// TestRegisterProtectedRoutes tests the registration of protected routes.
func TestRegisterProtectedRoutes(t *testing.T) {
	// Initialize test database and router
	db := tests.InitializeTestDB(t)
	router := gin.New()

	protected := router.Group("/api/v1/protected")
	protected.Use(middleware.JwtMiddleware())
	{
		protected.GET("/data", GetProtectedData(db))
	}

	// Create a valid JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": 1,
		"exp":     time.Now().Add(time.Hour).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		t.Fatalf("Could not sign token: %v", err)
	}

	// Test valid JWT
	req, err := http.NewRequest(http.MethodGet, "/api/v1/protected/data", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+tokenString)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Check response for valid JWT
	assert.Equal(t, http.StatusOK, rr.Code)

	// Test invalid JWT
	req, err = http.NewRequest(http.MethodGet, "/api/v1/protected/data", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer invalidtoken")

	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Check response for invalid JWT
	assert.Equal(t, http.StatusUnauthorized, rr.Code)

	expected := "{\"error\":\"Invalid token\"}"
	assert.Equal(t, expected, rr.Body.String())
}
