package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       int8
	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
	Username string `gorm:"not null"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
