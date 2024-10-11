package routes

import (
	"github.com/gorilla/mux"
	"github.com/jul3x/WebappBoilerplate/handlers"
	"github.com/jul3x/WebappBoilerplate/middlewares"
	"gorm.io/gorm"
)

func RegisterProtectedRoutes(router *mux.Router, db *gorm.DB) {
	router.HandleFunc("/api/v1/protected/data", middlewares.JwtMiddleware(handlers.GetProtectedData(db))).Methods("GET")
}
