package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	IsAdmin   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
