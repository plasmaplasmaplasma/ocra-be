package entities

import (
	"time"

	"gorm.io/gorm"
)

// BaseEntity centralizes common persistence fields for all entities.
type BaseEntity struct {
	ID        int64          `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
