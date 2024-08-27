package controllers

import (
	"Expense_Management/backend/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetMemberRequest struct {
	GroupName string `json:"group_name" binding:"required"`
}

func GetMembersByGroupName(c *gin.Context) {
	var req GetMemberRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Group name is required"})
		return
	}

	users, err := repositories.GetMembersByGroupName(req.GroupName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Check if the user was found
	if users == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "No members in group"})
		return
	}

	memberData := make([]map[string]interface{}, len(users))
	for i, user := range users {
		memberData[i] = map[string]interface{}{
			"user_name":  user.UserName,
			"user_email": user.UserEmail,
		}
	}

	c.JSON(http.StatusOK, memberData)

}
