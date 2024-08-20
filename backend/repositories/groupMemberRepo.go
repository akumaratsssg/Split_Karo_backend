// /backend/repositories/groupParticipantRepository.go

package repositories

import (
	"Expense_Management/backend/database"
	"Expense_Management/backend/models"
	"errors"
)

func AddParticipant(participant *models.GroupParticipant) error {
	return database.DB.Create(participant).Error
}

func RemoveParticipant(userID, groupID uint) error {
	result := database.DB.Where("part_user_id = ? AND part_group_id = ?", userID, groupID).Delete(&models.GroupParticipant{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("user already absent in group")
	}

	return nil
}
