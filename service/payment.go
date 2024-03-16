package service

import (
	"backend-marketplace-openidea/models"
	"backend-marketplace-openidea/models/payload"
	"backend-marketplace-openidea/repository/database"
)

func CreatePayment(req *payload.CreatePaymentRequest, userId string) (err error) {

	_, err = database.GetBankAccountByID(req.BankAccountID)
	if err != nil {
		return
	}

	newPayment := &models.Payment{
		UserID:               userId,
		BankAccountID:        req.BankAccountID,
		PaymentProofImageUrl: req.PaymentProofImageUrl,
		Quantity:             req.Quantity,
	}

	err, _ = database.CreatePayment(newPayment)
	if err != nil {
		return
	}

	return
}
