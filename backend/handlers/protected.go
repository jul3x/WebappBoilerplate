package handlers

import (
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
)

func GetProtectedData(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("user_id")

		response := map[string]interface{}{
			"message": "This is protected data",
			"user_id": userID,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
