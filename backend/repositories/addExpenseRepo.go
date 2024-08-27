package repositories

import (
	"Expense_Management/backend/database"
	"Expense_Management/backend/models"
	"errors"

	"gorm.io/gorm"
)

func GetGroupIDByName(groupName string) (uint, error) {
	var group models.Group

	// Attempt to find the group by its name
	if err := database.DB.Where("group_name = ?", groupName).First(&group).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, errors.New("group not found")
		}
		return 0, err
	}

	// Return the group_id
	return group.GroupID, nil
}

func InsertExpense(expense *models.Expense) error {
	// Insert the expense into the database
	if err := database.DB.Create(expense).Error; err != nil {
		return err
	}

	return nil
}
