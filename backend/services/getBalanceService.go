package services

import (
	"Expense_Management/backend/repositories"
	"errors"
)

type getBalancesResponse struct {
	UserName string  `json:"user_name"`
	Balance  float64 `json:"balance"`
}

func GetBalancesForGroup(UserID uint, groupName string) ([]getBalancesResponse, error) {
	// Fetch group ID from group name
	groupID, err := repositories.GetGroupIDByName(groupName)
	if err != nil {
		return nil, errors.New("invalid group name")
	}

	// Fetch debt records for the user within the group
	debts, err := repositories.GetDebtRecordsForGroup(UserID, groupID)
	if err != nil {
		return nil, err
	}

	// Calculate balances with each member
	balances := make(map[uint]float64)
	for _, debt := range debts {
		if debt.DebtLendID == UserID {
			balances[debt.DebtBorrID] -= debt.DebtAmount
		} else if debt.DebtBorrID == UserID {
			balances[debt.DebtLendID] += debt.DebtAmount
		}
	}
	// Prepare the response list
	var response []getBalancesResponse
	for userID, balance := range balances {
		// Fetch the user name based on user ID
		userName, err := repositories.GetUserNameByID(userID)
		if err != nil {
			return nil, err
		}

		// Append to the response list
		response = append(response, getBalancesResponse{
			UserName: userName,
			Balance:  balance,
		})
	}

	return response, nil
}
