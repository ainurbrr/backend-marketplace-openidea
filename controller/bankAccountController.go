package controller

import (
	"backend-marketplace-openidea/middleware"
	"backend-marketplace-openidea/models/payload"
	"backend-marketplace-openidea/service"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateBankAccountController(c echo.Context) error {

	userId, err := middleware.ExtractUserID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"Message": "you must be re login",
		})
	}

	payloadBankAccount := payload.CreateBankAccountRequest{}
	c.Bind(&payloadBankAccount)

	if err := c.Validate(&payloadBankAccount); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error payload create BankAccount",
			"error":   err.Error(),
		})
	}

	response, err := service.CreateBankAccount(&payloadBankAccount, userId)
	if err != nil {
		return c.JSON(http.StatusConflict, map[string]interface{}{
			"messages": "error create BankAccount",
			"error":    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, payload.Response{
		Message: "BankAccount successfully added",
		Data:    response,
	})
}

func GetBankAccountByUserIdController(c echo.Context) error {
	userId, err := middleware.ExtractUserID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"Message": "you must be re login",
		})
	}

	response, err := service.GetBankAccountByUserId(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, payload.Response{
		Message: "success",
		Data:    response,
	})
}

func UpdateBankAccountController(c echo.Context) error {
	userId, err := middleware.ExtractUserID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"Message": "you must be re login",
		})
	}

	payloadBank := payload.UpdateBankAccountRequest{}
	c.Bind(&payloadBank)

	id := c.Param("bankAccountId")

	bank, err := service.GetBankAccountById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	if bank.UserID != userId {
		return c.JSON(http.StatusForbidden, errors.New("forbidden"))
	}

	if err := c.Validate(payloadBank); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error payload update Bank",
			"error":   err.Error(),
		})
	}

	err = service.UpdateBankAccount(bank, &payloadBank)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, payload.Response{
		Message: "success update Bank Account data",
		Data:    http.StatusOK,
	})

}

func DeleteBankAccountByUserIdController(c echo.Context) error {
	userId, err := middleware.ExtractUserID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"Message": "you must be re login",
		})
	}

	id := c.Param("bankAccountId")

	bank, err := service.GetBankAccountById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	if bank.UserID != userId {
		return c.JSON(http.StatusForbidden, errors.New("forbidden"))
	}

	err = service.DeleteBankAccount(bank)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, payload.Response{
		Message: "success",
		Data:    http.StatusOK,
	})
}
