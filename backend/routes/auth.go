package routes

import (
        "github.com/gin-gonic/gin"
        "github.com/jul3x/WebappBoilerplate/handlers"
        "gorm.io/gorm"
)

// RegisterAuthRoutes sets up the authentication routes
func RegisterAuthRoutes(router *gin.Engine, db *gorm.DB) {
        // Register route
        router.POST("/api/v1/auth/register", handlers.Register(db))

        // Login route
        router.POST("/api/v1/auth/login", handlers.Login(db))
}
