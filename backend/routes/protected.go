package routes

import (
	"github.com/gorilla/mux"
	"github.com/jul3x/webapp-boilerplate/handlers"
	"github.com/jul3x/webapp-boilerplate/middlewares"
	"gorm.io/gorm"
)

func RegisterProtectedRoutes(router *mux.Router, db *gorm.DB) {
	router.HandleFunc("/api/v1/protected/data", middlewares.JwtMiddleware(handlers.GetProtectedData(db))).Methods("GET")
}
