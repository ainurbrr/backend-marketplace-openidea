package service

import (
	"backend-marketplace-openidea/models"
	"backend-marketplace-openidea/models/payload"
	"backend-marketplace-openidea/repository/database"
	"errors"
)

func CreateBankAccount(req *payload.CreateBankAccountRequest, userId string) (resp payload.ManageBankAccountResponse, err error) {

	newBankAccount := &models.BankAccount{
		BankName:          req.BankName,
		UserID:            userId,
		BankAccountName:   req.BankAccountName,
		BankAccountNumber: req.BankAccountNumber,
	}

	err, id := database.CreateBankAccount(newBankAccount)
	if err != nil {
		return
	}

	resp = payload.ManageBankAccountResponse{
		BankAccountId:     id,
		BankName:          req.BankName,
		BankAccountName:   req.BankAccountName,
		BankAccountNumber: req.BankAccountNumber,
	}

	return
}

func GetBankAccountById(id string) (*models.BankAccount, error) {
	bank, err := database.GetBankAccountByID(id)
	if err != nil {
		return nil, err
	}

	return bank, nil
}

func GetBankAccountByUserId(userId string) (bankAccounts []payload.ManageBankAccountResponse, err error) {
	bankAccounts, err = database.GetBankAccountByUserId(userId)
	if err != nil {
		return bankAccounts, errors.New("bankAccounts not found")
	}

	return bankAccounts, nil
}

func UpdateBankAccount(bank *models.BankAccount, req *payload.UpdateBankAccountRequest) error {
	bank.BankName = req.BankName
	bank.BankAccountName = req.BankAccountName
	bank.BankAccountNumber = req.BankAccountNumber

	err := database.UpdateBankAccount(bank)
	if err != nil {
		return err
	}
	return nil
}


func DeleteBankAccount(bank *models.BankAccount) error {
	err := database.DeleteBankAccountByID(bank.ID)
	if err != nil {
		return err
	}
	return nil
}
