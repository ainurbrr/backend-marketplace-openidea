package payload

type CreatePaymentRequest struct {
	BankAccountID        string `json:"bankAccountId"`
	PaymentProofImageUrl string `json:"paymentProofImageUrl" validate:"required,url"`
	Quantity             int    `json:"quantity"`
}
