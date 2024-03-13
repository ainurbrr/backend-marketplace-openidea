package database

import (
	"backend-marketplace-openidea/config"
	"backend-marketplace-openidea/models"
	"fmt"
)

func CreateBankAccount(bank *models.BankAccount) (error, string) {
	var id string
	if err := config.DB.QueryRow(`INSERT INTO bank_accounts (bank_name, bank_account_name, bank_account_number) VALUES ($1, $2, $3) RETURNING id`, bank.BankName, bank.BankAccountName, bank.BankAccountNumber).Scan(&id); err != nil {
		return err, id
	}
	fmt.Printf("Inserted a single record %v", id)
	return nil, id
}
