package repositories

import (
	"errors"

	"Expense_Management/backend/database"
	"Expense_Management/backend/models"

	"gorm.io/gorm"
)

func GetMembersByGroupName(groupName string) ([]models.User, error) {
	var users []models.User
	err := database.DB.
		// Start with the EM_group_participants table
		Table("EM_group_participants").
		// Join the EM_user table to get user details
		Joins("JOIN EM_user ON EM_user.user_id = EM_group_participants.part_user_id").
		// Join the EM_group table to filter by group name
		Joins("JOIN EM_group ON EM_group.group_id = EM_group_participants.part_group_id").
		// Filter the results by the group name
		Where("EM_group.group_name = ?", groupName).
		// Select the fields from the EM_user table
		Select("EM_user.*").
		// Order the results if needed
		Find(&users).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // No users found
		}
		return nil, err // Other errors
	}

	return users, nil
}
