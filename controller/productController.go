package controller

import (
	"backend-marketplace-openidea/models/payload"
	"backend-marketplace-openidea/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateProduct(c echo.Context) error {
	payloadProduct := payload.CreateProductRequest{}
	c.Bind(&payloadProduct)

	if err := c.Validate(&payloadProduct); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error payload create Product",
			"error":   err.Error(),
		})
	}
	response, err := service.CreateProduct(&payloadProduct)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "error create Product",
			"error":    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, payload.Response{
		Message: "success Create Product",
		Data:    response,
	})
}
