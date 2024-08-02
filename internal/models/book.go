package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	GUID            string `gorm:"unique;not null"`
	Name            string `gorm:"not null"`
	Author          string `gorm:"not null"`
	NumberOfBorrows int
	CurrentBorrower string
	BorrowedAt      *time.Time
}
