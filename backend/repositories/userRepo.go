package repositories

import (
	"Expense_Management/backend/database"
	"Expense_Management/backend/models"
	"errors"

	"gorm.io/gorm"
)

func CreateUser(user *models.User) error {
	return database.DB.Create(user).Error
}

func FindUserByEmail(email string) (models.User, error) {
	var user models.User
	result := database.DB.Table("EM_user").Where("user_email = ?", email).Scan(&user)
	if result.Error != nil || result.Error == gorm.ErrRecordNotFound {
		return user, errors.New("user not found")
	}
	return user, nil
}

func GetUserNameByID(userID uint) (string, error) {
	var user models.User
	err := database.DB.Where("user_id = ?", userID).First(&user).Error
	if err != nil {
		return "", errors.New("user not found")
	}
	return user.UserName, nil
}
