package database

import (
	"backend-marketplace-openidea/config"
	"backend-marketplace-openidea/models"
	"backend-marketplace-openidea/models/payload"
	"fmt"
)

func CreateBankAccount(bank *models.BankAccount) (error, string) {
	var id string
	if err := config.DB.QueryRow(`INSERT INTO bank_accounts (user_id, bank_name, bank_account_name, bank_account_number) VALUES ($1, $2, $3, $4) RETURNING id`, bank.UserID, bank.BankName, bank.BankAccountName, bank.BankAccountNumber).Scan(&id); err != nil {
		return err, id
	}
	fmt.Printf("Inserted a single record %v", id)
	return nil, id
}

func GetBankAccountByID(id string) (*models.BankAccount, error) {
	// Execute the SQL query to retrieve bank account data by ID
	var bankAccount models.BankAccount
	if err := config.DB.QueryRow("SELECT id, user_id, bank_name, bank_account_name, bank_account_number FROM bank_accounts WHERE id = $1", id).
		Scan(&bankAccount.ID, &bankAccount.UserID, &bankAccount.BankName, &bankAccount.BankAccountName, &bankAccount.BankAccountNumber); err != nil {
		return nil, err
	}

	return &bankAccount, nil
}

func GetBankAccountByUserId(userId string) (resp []payload.ManageBankAccountResponse, err error) {
	// Execute the SQL query
	rows, err := config.DB.Query("SELECT id, bank_name, bank_account_name, bank_account_number FROM bank_accounts WHERE user_id = $1", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the rows
	for rows.Next() {
		var bankAccount payload.ManageBankAccountResponse
		// Scan the row into the bankAccount struct
		err := rows.Scan(&bankAccount.BankAccountId, &bankAccount.BankName, &bankAccount.BankAccountName, &bankAccount.BankAccountNumber)
		if err != nil {
			return nil, err
		}
		// Append the bankAccount to the slice
		resp = append(resp, bankAccount)
	}
	// Check for errors during rows iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return resp, nil
}

func UpdateBankAccount(bank *models.BankAccount) error {
	_, err := config.DB.Exec("UPDATE bank_accounts SET bank_name = COALESCE($1, bank_name), bank_account_name = COALESCE($2, bank_account_name), bank_account_number = COALESCE($3, bank_account_number) WHERE id = $4",
		bank.BankName, bank.BankAccountName, bank.BankAccountNumber, bank.ID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteBankAccountByID(id string) error {
	// Execute the SQL query to delete bank account data by ID
	_, err := config.DB.Exec("DELETE FROM bank_accounts WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
