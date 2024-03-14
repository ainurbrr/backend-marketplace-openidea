package models

type BankAccount struct {
	ID                string `json:"id"`
	UserID            string `json:"userID"`
	BankName          string `json:"bankName" validate:"required,min=5,max=15"`
	BankAccountName   string `json:"bankAccountName" validate:"required,min=5,max=15"`
	BankAccountNumber string `json:"bankAccountNumber" validate:"required,min=5,max=15"`
	User User
}
