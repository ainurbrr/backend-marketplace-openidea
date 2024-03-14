package controller

import (
	"backend-marketplace-openidea/models/payload"
	"backend-marketplace-openidea/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateProductController(c echo.Context) error {
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
		Message: "product added successfully",
		Data:    response,
	})
}

func UpdateProductController(c echo.Context) error {
	productID := c.Param("productId")
	payloadProduct := payload.UpdateProductRequest{}
	if err := c.Bind(&payloadProduct); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error payload update Product",
			"error":   err.Error(),
		})
	}

	if err := c.Validate(&payloadProduct); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error payload update Product",
			"error":   err.Error(),
		})
	}
	response, err := service.UpdateProduct(productID, &payloadProduct)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error update Product",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, payload.Response{
		Message: "product updated successfully",
		Data:    response,
	})
}
