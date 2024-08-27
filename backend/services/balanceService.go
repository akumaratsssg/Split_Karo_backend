package services

import (
	"Expense_Management/backend/models"
	"Expense_Management/backend/repositories"
)

func CreateBalance(participants []string, expAmount float64, expLendID uint, expID uint) ([]models.DebtRecordResponse, error) {
	debtAmount := expAmount / float64(len(participants)+1)

	var debtRecords []models.DebtRecord
	for _, participant := range participants {
		debtRecord, err := repositories.CreateDebtRecord(participant, debtAmount, expLendID, expID)
		if err != nil {
			return nil, err
		}
		debtRecords = append(debtRecords, debtRecord)
	}

	// Transform debtRecords to DebtRecordResponse
	var debtRecordResponses []models.DebtRecordResponse
	for _, record := range debtRecords {
		debtRecordResponses = append(debtRecordResponses, models.DebtRecordResponse{
			DebtID:     record.DebtID,
			DebtAmount: record.DebtAmount,
			DebtLendID: record.DebtLendID,
			DebtBorrID: record.DebtBorrID,
			DebtExpID:  record.DebtExpID,
		})
	}

	return debtRecordResponses, nil
}
