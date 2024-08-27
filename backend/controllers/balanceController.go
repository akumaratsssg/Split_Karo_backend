package controllers

import (
	"Expense_Management/backend/models"
	"Expense_Management/backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBalance(c *gin.Context) {
	var req models.CreateBalanceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	UserID, exists := c.Get("UserID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized !!!!"})
		return
	}

	req.ExpLendID = UserID.(uint)

	debtRecords, err := services.CreateBalance(req.Participants, req.ExpAmount, req.ExpLendID, req.ExpID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, debtRecords)
}
