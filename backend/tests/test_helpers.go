package tests

import (
    "testing"
    "github.com/jul3x/webapp-boilerplate/models"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// InitializeTestDB creates an in-memory SQLite database for testing
func InitializeTestDB(t *testing.T) *gorm.DB {
    db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
    if err != nil {
        t.Fatalf("Failed to connect to test database: %v", err)
    }

    db.AutoMigrate(&models.User{})
    return db
}
