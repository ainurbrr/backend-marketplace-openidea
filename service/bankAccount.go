package service

import (
	"backend-marketplace-openidea/models"
	"backend-marketplace-openidea/models/payload"
	"backend-marketplace-openidea/repository/database"
)

func CreateBankAccount(req *payload.CreateBankAccountRequest) (resp payload.ManageBankAccountResponse, err error) {

	newBankAccount := &models.BankAccount{
		BankName: req.BankName,
		BankAccountName: req.BankAccountName,
		BankAccountNumber: req.BankAccountNumber,
	}

	err, id := database.CreateBankAccount(newBankAccount)
	if err != nil {
		return
	}

	resp = payload.ManageBankAccountResponse{
		BankAccountId: id,
		BankName: req.BankName,
		BankAccountName: req.BankAccountName,
		BankAccountNumber: req.BankAccountNumber,
	}

	return
}
