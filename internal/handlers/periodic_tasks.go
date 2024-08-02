package handlers

import (
	"log"
	"time"

	"github.com/ealekseychik/mnemosyne/internal/models"
)

func StartPeriodicTasks() {
	ticker := time.NewTicker(24 * time.Hour)
	go func() {
		for range ticker.C {
			checkOverdueBooks()
		}
	}()
}

func checkOverdueBooks() {
	var books []models.Book
	overdueDuration := 14 * 24 * time.Hour

	if err := models.DB.Where("current_borrower IS NOT NULL AND borrowed_at <= ?", time.Now().Add(-overdueDuration)).Find(&books).Error; err != nil {
		log.Printf("Failed to fetch overdue books: %v", err)
		return
	}

	for _, book := range books {
		subject := "Return Book Reminder"
		body := "Please return the book: " + book.Name + " by " + book.Author
		if err := sendEmail(book.CurrentBorrower, subject, body); err != nil {
			log.Printf("Failed to send email to %s: %v", book.CurrentBorrower, err)
		}
	}
}
