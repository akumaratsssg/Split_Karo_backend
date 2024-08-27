package repositories

import (
	"errors"

	"Expense_Management/backend/database"
	"Expense_Management/backend/models"

	"gorm.io/gorm"
)

func FindUserByUsername(usernamePrefix string) (*models.User, error) {
	var user models.User
	if err := database.DB.Where("user_name LIKE ?", usernamePrefix+"%").First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // No user found
		}
		return nil, err // Other errors
	}
	return &user, nil
}
