package controllers

import (
	"Expense_Management/backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BalanceRequest struct {
	GroupName string `json:"group_name" binding:"required"`
}

func GetGroupBalances(c *gin.Context) {
	// Bind JSON request to struct
	var req BalanceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get current user ID from context
	currentUserID, exists := c.Get("UserID")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized !!!!"})
		return
	}

	// Call the service to get balances
	balances, err := services.GetBalancesForGroup(currentUserID.(uint), req.GroupName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, balances)
}
