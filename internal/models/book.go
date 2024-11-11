package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	GUID            string  `gorm:"unique;not null"`
	Name            string  `gorm:"not null"`
	Author          string  `gorm:"not null"`
	Note            *string // Invisible for user
	NumberOfBorrows int
	CurrentBorrower string
	BorrowedAt      *time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
