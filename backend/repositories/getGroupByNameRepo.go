// /backend/repositories/groupGetByNameRepository.go

package repositories

import (
	"Expense_Management/backend/database"
	"Expense_Management/backend/models"
	"errors"
)

func GetGroupByName(groupName string) (models.Group, error) {
	var group models.Group
	err := database.DB.Where("group_name = ?", groupName).First(&group).Error
	if err != nil {
		return group, errors.New("group not found")
	}
	return group, nil
}
