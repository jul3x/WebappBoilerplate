package models

import "gorm.io/gorm"

type Role string

const (
	RoleNormal Role = "normal"
	RoleAdmin  Role = "admin"
)

type User struct {
    gorm.Model
    Username string `gorm:"unique"`
    Password string
    Email    string `gorm:"unique"`
    Role     Role
}
