package models

import (
	"time"
)

type Expense struct {
	ExpID       uint      `gorm:"primaryKey;autoIncrement;not null;unique" json:"exp_id"`
	ExpAmount   float32   `gorm:"not null" json:"exp_amount"`
	ExpDate     time.Time `gorm:"not null;default:current_date + interval 1 year" json:"exp_date"`
	ExpCategory string    `gorm:"type:varchar(45);not null" json:"exp_category"`
	ExpDesc     string    `gorm:"type:varchar(45);default:null" json:"exp_desc"`
	ExpGroupID  uint      `gorm:"not null" json:"exp_group_id"`
	ExpLendID   uint      `gorm:"not null" json:"exp_lend_id"`

	Group  Group `gorm:"foreignKey:ExpGroupID;references:GroupID" `
	Lender User  `gorm:"foreignKey:ExpLendID;references:UserID" `
}

func (Expense) TableName() string {
	return "EM_expense"
}

type AddExpenseInput struct {
	ExpAmount    float32 `json:"exp_amount" binding:"required"`
	ExpCategory  string  `json:"exp_category" binding:"required"`
	ExpDesc      string  `json:"exp_desc" binding:"required"`
	ExpGroupName string  `json:"group_name" binding:"required"`
}
