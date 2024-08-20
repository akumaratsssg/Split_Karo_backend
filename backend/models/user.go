package models

// User represents a user in the EM_user table.
type User struct {
	UserID       uint    `gorm:"primaryKey;autoIncrement;unsigned" json:"user_id"`
	UserName     string  `gorm:"type:varchar(50);not null" json:"user_name"`
	UserEmail    string  `gorm:"type:varchar(50);unique;not null" json:"user_email"`
	UserPassword string  `gorm:"type:varchar(255);not null" json:"user_password"`
	UserBalance  float64 `gorm:"type:float;not null" json:"user_balance"`
}

// TableName sets the table name for the User struct
func (User) TableName() string {
	return "EM_user"
}

type RegisterInput struct {
	UserName     string `json:"user_name" binding:"required"`
	UserEmail    string `json:"user_email" binding:"required,email"`
	UserPassword string `json:"user_password" binding:"required"`
}

type LoginInput struct {
	UserEmail    string `json:"user_email" binding:"required,email"`
	UserPassword string `json:"user_password" binding:"required"`
}
