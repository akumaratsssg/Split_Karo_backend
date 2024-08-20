package repositories

import (
	"Expense_Management/backend/database"
	"Expense_Management/backend/models"
)

func CreateGroup(group *models.Group) error {
	// Start a database transaction
	tx := database.DB.Begin()

	// Create the group
	if err := tx.Create(group).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Add the group admin to the group participants table
	groupParticipant := &models.GroupParticipant{
		PartUserID:  group.GroupAdmin,
		PartGroupID: group.GroupID,
	}

	if err := tx.Create(groupParticipant).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Commit the transaction
	return tx.Commit().Error
}
