package controllers

import (
	"fmt"
	"net/http"

	"Expense_Management/backend/repositories"

	"github.com/gin-gonic/gin"
)

// SearchUserRequest is the request payload for searching a user by username.
type SearchUserRequest struct {
	UserName string `json:"user_name" binding:"required"`
}

// SearchUserByUsername handles the API request to search for a user by their username.
func SearchUserByUsername(c *gin.Context) {
	var req SearchUserRequest

	// Bind the incoming JSON request to the req struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User name is required"})
		return
	}

	// Log the received username
	fmt.Println("input username:", req.UserName)

	// Search for the user by username using the repository
	user, err := repositories.FindUserByUsername(req.UserName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Check if the user was found
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "No user found"})
		return
	}

	// Return the user information if found
	c.JSON(http.StatusOK, gin.H{
		"user_name":  user.UserName,
		"user_email": user.UserEmail,
	})
}
