package services

import (
	"Expense_Management/backend/models"
	"Expense_Management/backend/repositories"
)

func AddExpense(input models.AddExpenseInput, lendID uint) (*models.Expense, error) {
	var expGroupID, err = repositories.GetGroupIDByName(input.ExpGroupName)
	if err != nil {
		return nil, err
	}
	expense := &models.Expense{
		ExpAmount:   input.ExpAmount,
		ExpCategory: input.ExpCategory,
		ExpDesc:     input.ExpDesc,
		ExpGroupID:  expGroupID,
		ExpLendID:   lendID,
	}

	err = repositories.InsertExpense(expense)
	if err != nil {
		return nil, err
	}

	return expense, nil
}
