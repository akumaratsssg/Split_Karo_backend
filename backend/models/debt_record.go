package models

type DebtRecord struct {
	DebtID     uint    `gorm:"primaryKey;autoIncrement;not null" json:"debt_id"`
	DebtAmount float64 `gorm:"not null" json:"debt_amount"`
	DebtLendID uint    `gorm:"not null" json:"debt_lend_id"`
	DebtBorrID uint    `gorm:"not null" json:"debt_borr_id"`
	DebtExpID  uint    `gorm:"not null" json:"debt_exp_id"`

	Lender   User    `gorm:"foreignKey:DebtLendID;references:UserID" `
	Borrower User    `gorm:"foreignKey:DebtBorrID;references:UserID" `
	ExpID    Expense `gorm:"foreignKey:DebtExpID;references:ExpID" `
}

func (DebtRecord) TableName() string {
	return "EM_debt_record"
}

type CreateBalanceRequest struct {
	Participants []string `json:"participants" binding:"required"`
	ExpAmount    float64  `json:"exp_amount" binding:"required"`
	ExpLendID    uint     `json:"exp_lend_id" binding:"required"`
	ExpID        uint     `json:"exp_id" binding:"required"`
}

type DebtRecordResponse struct {
	DebtID     uint    `json:"debt_id"`
	DebtAmount float64 `json:"debt_amount"`
	DebtLendID uint    `json:"debt_lend_id"`
	DebtBorrID uint    `json:"debt_borr_id"`
	DebtExpID  uint    `json:"debt_exp_id"`
}
