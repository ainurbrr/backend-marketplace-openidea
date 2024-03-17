package service

import (
	"backend-marketplace-openidea/models"
	"backend-marketplace-openidea/models/payload"
	"backend-marketplace-openidea/repository/database"
	"errors"
)

func CreatePayment(req *payload.CreatePaymentRequest, userId string, productId string) (err error) {
	product, err := database.GetProductByID(productId)
	if err != nil {
		return
	}

	if !product.IsPurchaseable {
		return errors.New("Stock Habis")
	}

	_, err = database.GetBankAccountByID(req.BankAccountID)
	if err != nil {
		return
	}

	if product.Stock < req.Quantity {
		return errors.New("Pesanan anda melebihi stock")
	}

	product.Stock -= req.Quantity
	product.PurchaseCount += 1

	if product.Stock <= 0 {
		product.IsPurchaseable = false
		database.UpdateProductAfterPay(product)
	} else {
		product.IsPurchaseable = true
		database.UpdateProductAfterPay(product)
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
