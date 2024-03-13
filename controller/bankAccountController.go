package controller

import (
	"backend-marketplace-openidea/models/payload"
	"backend-marketplace-openidea/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateBankAccountController(c echo.Context) error {

	payloadBankAccount := payload.CreateBankAccountRequest{}
	c.Bind(&payloadBankAccount)

	if err := c.Validate(&payloadBankAccount); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error payload create BankAccount",
			"error":   err.Error(),
		})
	}

	response, err := service.CreateBankAccount(&payloadBankAccount)
	if err != nil {
		return c.JSON(http.StatusConflict, map[string]interface{}{
			"messages": "error create BankAccount",
			"error":    err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, payload.Response{
		Message: "BankAccount registered successfully",
		Data:    response,
	})
}
