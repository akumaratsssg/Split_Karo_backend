package controllers

import (
	"Expense_Management/backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

/// GET GROUP CONTROLLER STARTS

type GroupController struct {
	groupService *services.GroupService
}

func NewGroupController(groupService *services.GroupService) *GroupController {
	return &GroupController{groupService: groupService}
}

// GetGroups fetches groups created by the logged-in user.
func (ctrl *GroupController) GetGroups(c *gin.Context) {
	userID, exists := c.Get("UserID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}

	groups, err := ctrl.groupService.GetGroupsByUser(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	groupData := make([]map[string]interface{}, len(groups))
	for i, group := range groups {
		groupData[i] = map[string]interface{}{
			"group_name":       group.GroupName,
			"group_desc":       group.GroupDesc,
			"group_admin_name": group.Admin.UserName, // Extract only the `UserName` field.
		}
	}

	c.JSON(http.StatusOK, groupData)
}
