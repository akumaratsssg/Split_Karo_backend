package repositories

import (
	"Expense_Management/backend/database"
	"Expense_Management/backend/models"
	"errors"
)

func CreateDebtRecord(participantName string, debtAmount float64, expLendID uint, expID uint) (models.DebtRecord, error) {
	var user models.User
	if err := database.DB.Where("user_name = ?", participantName).First(&user).Error; err != nil {
		return models.DebtRecord{}, errors.New("participant not found")
	}

	debtRecord := models.DebtRecord{
		DebtAmount: debtAmount,
		DebtLendID: uint(expLendID),
		DebtBorrID: user.UserID,
		DebtExpID:  uint(expID),
	}

	if err := database.DB.Create(&debtRecord).Error; err != nil {
		return models.DebtRecord{}, err
	}

	return debtRecord, nil
}
