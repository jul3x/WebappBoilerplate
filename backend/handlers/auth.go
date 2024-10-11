package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/jul3x/WebappBoilerplate/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"os"
)

type AuthResponse struct {
	Token string `json:"token"`
}

func Register(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		json.NewDecoder(r.Body).Decode(&user)

                var existingUser models.User
                db.Where("email = ?", user.Email).First(&existingUser)
                if existingUser.ID != 0 {
                        http.Error(w, "Email already in use", http.StatusConflict)
                        return
                }

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}
		user.Password = string(hashedPassword)

		if result := db.Create(&user); result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
                user.Password = "<hidden>"
		json.NewEncoder(w).Encode(user)
	}
}

func Login(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		var foundUser models.User

		json.NewDecoder(r.Body).Decode(&user)

		db.Where("email = ?", user.Email).First(&foundUser)
		if foundUser.ID == 0 {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))
		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": foundUser.ID,
			"exp":     time.Now().Add(time.Hour * 24).Unix(),
		})

		tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
		if err != nil {
			http.Error(w, "Error generating token", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(AuthResponse{Token: tokenString})
	}
}
