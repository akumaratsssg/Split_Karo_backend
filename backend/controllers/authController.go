package controllers

import (
	"Expense_Management/backend/models"
	"Expense_Management/backend/services"
	"Expense_Management/backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "welcome to App Home",
	})
}

func Register(c *gin.Context) {
	var input models.RegisterInput

	if err := c.BindJSON(&input); err != nil {
		print(err.Error(), "error is this")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}
	user, err := services.CreateUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func Login(c *gin.Context) {
	var input models.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, userName, err := services.LoginUser(input, c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Return the token and the user_name in the response
	c.JSON(http.StatusOK, gin.H{
		"token":     token,
		"user_name": userName,
	})
}

func Logout(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No token provided"})
		return
	}

	err := utils.InvalidateToken(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to logout"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
