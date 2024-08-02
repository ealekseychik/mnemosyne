package handlers

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("static/html/*")
	router.Static("/static", "./static")

	// User endpoints
	router.GET("/ping", Ping)
	router.GET("/borrow/:bookGUID", BorrowBook)

	// Admin endpoints
	admin := router.Group("/admin")
	{
		admin.POST("/login", AdminLogin)
		admin.GET("/dashboard", AuthMiddleware(), AdminDashboard)
		admin.POST("/book", AuthMiddleware(), AddBook)
		admin.PUT("/book/:bookGUID", AuthMiddleware(), EditBook)
		admin.DELETE("/book/:bookGUID", AuthMiddleware(), DeleteBook)
	}

	return router
}
