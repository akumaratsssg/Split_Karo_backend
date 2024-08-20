package services

import (
	"Expense_Management/backend/models"
	"Expense_Management/backend/repositories"
	"Expense_Management/backend/utils"
	"errors"

	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(input models.RegisterInput) (*models.User, error) {

	// Validate password
	if err := utils.ValidatePassword(input.UserPassword); err != nil {
		return nil, err
	}

	///why ur doing this
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.UserPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		UserName:     input.UserName,
		UserEmail:    input.UserEmail,
		UserPassword: string(hashedPassword),
		UserBalance:  0.0,
	}

	err = repositories.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func LoginUser(input models.LoginInput, c *gin.Context) (string, string, error) {
	user, err := repositories.FindUserByEmail(input.UserEmail)
	if err != nil {
		// Log the error
		log.Printf("Error finding user: %v", err)
		return "", "", errors.New("user not found")
	}

	log.Printf("User found: %+v", user)

	err = bcrypt.CompareHashAndPassword([]byte(user.UserPassword), []byte(input.UserPassword))
	if err != nil {
		log.Printf("Password comparison failed %v", err)
		return "", "", errors.New("incorrect password")
	}

	token, err := utils.GenerateToken(user.UserID)
	if err != nil {
		log.Printf("Error generating token: %v", err)
		return "", "", err
	}

	return token, user.UserName, nil
}
