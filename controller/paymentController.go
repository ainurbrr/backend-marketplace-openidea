package controller

import (
	"backend-marketplace-openidea/middleware"
	"backend-marketplace-openidea/models/payload"
	"backend-marketplace-openidea/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreatePaymentController(c echo.Context) error {

	userId, err := middleware.ExtractUserID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"Message": "you must be re login",
		})
	}

	payloadPayment := payload.CreatePaymentRequest{}
	c.Bind(&payloadPayment)

	if err := c.Validate(&payloadPayment); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error payload create Payment",
			"error":   err.Error(),
		})
	}

	err = service.CreatePayment(&payloadPayment, userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "invalid! error create Payment",
			"error":    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, payload.Response{
		Message: "payment processed successfully",
		Data:    200,
	})
}
