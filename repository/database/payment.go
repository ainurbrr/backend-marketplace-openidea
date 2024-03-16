package database

import (
	"backend-marketplace-openidea/config"
	"backend-marketplace-openidea/models"
	"fmt"
)

func CreatePayment(payment *models.Payment) (error, string) {
	var id string
	if err := config.DB.QueryRow(`INSERT INTO payments (user_id, bank_account_id, product_id, payment_proof_image_url, quantity ) VALUES ($1, $2, $3, $4, $5) RETURNING id`, payment.UserID, payment.BankAccountID, 1, payment.PaymentProofImageUrl, payment.Quantity).Scan(&id); err != nil {
		return err, id
	}
	fmt.Printf("Inserted a single record %v", id)
	return nil, id
}
