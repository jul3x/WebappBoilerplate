package routes

import (
    "github.com/gorilla/mux"
    "gorm.io/gorm"
    "github.com/jul3x/WebappBoilerplate/handlers"
)

func RegisterAuthRoutes(router *mux.Router, db *gorm.DB) {
    router.HandleFunc("/api/v1/auth/register", handlers.Register(db)).Methods("POST")
    router.HandleFunc("/api/v1/auth/login", handlers.Login(db)).Methods("POST")
}
