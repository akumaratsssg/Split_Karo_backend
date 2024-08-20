package services

import (
	"Expense_Management/backend/models"
	"Expense_Management/backend/repositories"
)

func CreateGroup(input models.GroupInput, adminID uint) (*models.Group, error) {
	group := &models.Group{
		GroupName:  input.GroupName,
		GroupDesc:  input.GroupDesc,
		GroupAdmin: adminID, // Assigning the logged-in user as the admin
	}

	err := repositories.CreateGroup(group)
	if err != nil {
		return nil, err
	}

	return group, nil
}
