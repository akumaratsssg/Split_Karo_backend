package middlewares

import (
	"Expense_Management/backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
			c.Abort()
			return
		}

		claims, err := utils.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("UserID", claims.UserID)
		c.Next()
	}
}
