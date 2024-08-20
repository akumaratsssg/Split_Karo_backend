package utils

import (
	"Expense_Management/backend/config"
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

func GenerateToken(UserID uint) (string, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return "", err
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: UserID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWTSecret))
}

func ValidateToken(userToken string) (*Claims, error) {

	// Remove the "Bearer " prefix if it exists
	userToken = strings.TrimPrefix(userToken, "Bearer ")

	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(userToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func InvalidateToken(tokenString string) error {
	// Here you would typically use a token blacklist to invalidate the token
	// This is a placeholder implementation
	return nil
}
