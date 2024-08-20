// /backend/controllers/groupController.go

package controllers

import (
	"Expense_Management/backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GroupMemberRequest struct {
	GroupName string `json:"group_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
}

func AddMemberByGroupNameAndEmail(c *gin.Context) {
	var request GroupMemberRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	requestingUserID, exists := c.Get("UserID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized !!!!"})
		return
	}

	if err := services.AddMemberByGroupNameAndEmail(request.GroupName, request.Email, requestingUserID.(uint)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Member added successfully"})
}

func RemoveMemberByGroupNameAndEmail(c *gin.Context) {
	var request GroupMemberRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Assuming you retrieve the requesting user's ID from the context after JWT authentication
	requestingUserID, exists := c.Get("UserID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized !!!!"})
		return
	}

	if err := services.RemoveMemberByGroupNameAndEmail(request.GroupName, request.Email, requestingUserID.(uint)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Member removed successfully"})
}
