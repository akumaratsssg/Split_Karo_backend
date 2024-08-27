package controllers

import (
	"Expense_Management/backend/models"
	"Expense_Management/backend/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func AddExpenseToGroup(c *gin.Context) {
	var input models.AddExpenseInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	UserID, exists := c.Get("UserID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized !!!!"})
		return
	}

	expense, err := services.AddExpense(input, UserID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Define a custom response struct excluding the Admin field
	type ExpenseResponse struct {
		ExpID       uint      `json:"exp_id" `
		ExpAmount   float32   `json:"exp_amount" `
		ExpDate     time.Time `json:"exp_date"`
		ExpCategory string    `json:"exp_category" `
		ExpDesc     string    `json:"exp_desc" `
		ExpGroupID  uint      `json:"exp_group_id"`
		ExpLendID   uint      `json:"exp_lend_id"`
	}

	// Create the response data
	response := ExpenseResponse{
		ExpID:       expense.ExpID,
		ExpAmount:   expense.ExpAmount,
		ExpDate:     expense.ExpDate,
		ExpCategory: expense.ExpCategory,
		ExpDesc:     expense.ExpDesc,
		ExpGroupID:  expense.ExpGroupID,
		ExpLendID:   expense.ExpLendID,
	}

	c.JSON(http.StatusOK, gin.H{"data": response})

}
