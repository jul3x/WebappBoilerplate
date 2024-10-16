package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/jul3x/WebappBoilerplate/config"
	"github.com/jul3x/WebappBoilerplate/models"
	"github.com/jul3x/WebappBoilerplate/routes"
)

var db *gorm.DB
var err error

func init() {
}

func main() {
	configPath := "./"
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
		os.Exit(1)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db.AutoMigrate(&models.User{})

	router := gin.Default()

	// Add CORS middleware to allow requests from the frontend
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			fmt.Sprintf("https://%s:%d", cfg.Server.Host, cfg.Server.Port),
			fmt.Sprintf("http://%s:%d", cfg.Server.Host, cfg.Server.Port)},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routes.RegisterAuthRoutes(router, db)
	routes.RegisterProtectedRoutes(router, db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server is running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
