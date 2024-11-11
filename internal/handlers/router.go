package handlers

import (
	"github.com/ealekseychik/mnemosyne/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("static/html/*")
	router.Static("/static", "./static")

	// User endpoints
	router.GET("/ping", Ping)
	router.GET("/borrow/:bookGUID", GetBookStatus)
	router.POST("/borrow/:bookGUID", BorrowBook)

	// Admin endpoints
	admin := router.Group("/admin")
	{
		admin.POST("/login", AdminLogin)
		admin.GET("/dashboard", middleware.AuthMiddleware(), AdminDashboard)
		admin.POST("/book", middleware.AuthMiddleware(), AddBook)
		admin.PUT("/book/:bookGUID", middleware.AuthMiddleware(), EditBook)
		admin.DELETE("/book/:bookGUID", middleware.AuthMiddleware(), DeleteBook)
		admin.GET("/book/:bookGUID/ping", middleware.AuthMiddleware(), NotifyBookBorrower)
	}

	return router
}
