package handlers

import (
	"errors"
	"net/http"
	"net/mail"
	"time"

	"github.com/ealekseychik/mnemosyne/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func BorrowBook(c *gin.Context) {
	bookGUID := c.Param("bookGUID")

	var book models.Book
	if err := models.DB.Where("guid = ?", bookGUID).First(&book).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	if book.CurrentBorrower != "" {
		c.JSON(http.StatusOK, gin.H{"message": "This book is borrowed by " + book.CurrentBorrower + ". Do you want to return it?", "continue": true})
	} else {
		email := c.PostForm("email")
		if email == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
			return
		}

		if !isValidEmail(email) {
			c.JSON(http.StatusTeapot, gin.H{"error": "Invalid email format"})
			return
		}

		book.CurrentBorrower = email
		now := time.Now()
		book.BorrowedAt = &now
		if err := models.DB.Save(&book).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to borrow book"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "You have successfully borrowed the book"})
	}
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
