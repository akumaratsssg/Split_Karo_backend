package controllers

import (
	"Expense_Management/backend/models"
	"Expense_Management/backend/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Group(c *gin.Context) {
	var input models.GroupInput
	fmt.Println("input: ", input.GroupName, " ", input.GroupDesc)
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Extract the user ID from the context

	UserID, exists := c.Get("UserID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized !!!!"})
		return
	}

	// Call the service to create the group
	group, err := services.CreateGroup(input, UserID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Define a custom response struct excluding the Admin field
	type GroupResponse struct {
		GroupID   uint   `json:"group_id"`
		GroupName string `json:"group_name"`
		GroupDesc string `json:"group_desc"`
	}

	// Create the response data
	response := GroupResponse{
		GroupID:   group.GroupID,
		GroupName: group.GroupName,
		GroupDesc: group.GroupDesc,
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}
