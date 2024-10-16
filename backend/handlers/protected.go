package handlers

import (
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
	"net/http"
)

func GetProtectedData(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		response := map[string]interface{}{
			"message": "This is protected data",
			"user_id": userID,
		}

		c.JSON(http.StatusOK, response)
	}
}
