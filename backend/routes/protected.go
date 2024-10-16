package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jul3x/WebappBoilerplate/handlers"
	"github.com/jul3x/WebappBoilerplate/middlewares"
	"gorm.io/gorm"
)

func RegisterProtectedRoutes(router *gin.Engine, db *gorm.DB) {
	protected := router.Group("/api/v1/protected")
	protected.Use(middleware.JwtMiddleware())
	{
		protected.GET("/data", handlers.GetProtectedData(db))
	}
}
