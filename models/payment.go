package models

type Payment struct {
	ID                   string `json:"id"`
	UserID               string `json:"userId"`
	BankAccountID        string `json:"bankAccountId"`
	ProductID        string `json:"productId"`
	PaymentProofImageUrl string `json:"paymentProofImageUrl" validate:"required,url"`
	Quantity             int    `json:"quantity"`
	User                 User
	BankAccount          BankAccount
}
