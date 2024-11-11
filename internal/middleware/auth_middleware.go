package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" || !validateToken(token) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or missing token"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func validateToken(token string) bool {
	secret := []byte(os.Getenv("SECRET_KEY"))
	validatedToken, err := jwt.Parse(token,
		func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		},
	)
	if err != nil || !validatedToken.Valid {
		return false
	}

	return true
}
