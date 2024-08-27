package repositories

import (
	"Expense_Management/backend/database"
	"Expense_Management/backend/models"
)

// Fetch the group ID based on the group name
// todo: already present in addExpenseRepo.go file

// Fetch all debt records for the group where the user is involved
func GetDebtRecordsForGroup(userID uint, groupID uint) ([]models.DebtRecord, error) {
	var debts []models.DebtRecord

	if err := database.DB.Where("debt_exp_id IN (SELECT exp_id FROM EM_expense WHERE exp_group_id = ?)", groupID).
		Where("debt_lend_id = ? OR debt_borr_id = ?", userID, userID).
		Find(&debts).Error; err != nil {
		return nil, err
	}

	return debts, nil
}
