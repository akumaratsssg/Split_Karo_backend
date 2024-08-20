package utils

import (
	"errors"
	"regexp"
)

func ValidatePassword(password string) error {
	// Check if the password length is at least 8 characters
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	// Check if the password contains at least one uppercase letter
	if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
		return errors.New("password must contain at least one uppercase letter")
	}

	// Check if the password contains at least one lowercase letter
	if !regexp.MustCompile(`[a-z]`).MatchString(password) {
		return errors.New("password must contain at least one lowercase letter")
	}

	// Check if the password contains at least one digit
	if !regexp.MustCompile(`\d`).MatchString(password) {
		return errors.New("password must contain at least one digit")
	}

	// Check if the password contains at least one special character
	if !regexp.MustCompile(`[#@!$%^&*]`).MatchString(password) {
		return errors.New("password must contain at least one special character like #,@,!,$,%,^,&,*")
	}

	return nil
}
