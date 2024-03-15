package controller

import (
	"backend-marketplace-openidea/middleware"
	"backend-marketplace-openidea/models/payload"
	"backend-marketplace-openidea/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateProductController(c echo.Context) error {
	_, err := middleware.ExtractUserID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"Message": "you must be re login",
		})
	}
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
	_, err := middleware.ExtractUserID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"Message": "you must be re login",
		})
	}
	productID := c.Param("productId")

	productId, err := service.GetProductByID(productID)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	if productId.ID != productID {
		return c.JSON(http.StatusForbidden, "forbidden access")
	}

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

func GetProductById(c echo.Context) error {
	_, err := middleware.ExtractUserID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "You must be re-login",
		})
	}

	id := c.Param("productId")
	response, err := service.GetProductByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, payload.Response{
		Message: "success",
		Data:    response,
	})
}

func DeleteProductController(c echo.Context) error {
	_, err := middleware.ExtractUserID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"Message": "you must be re-login",
		})
	}
	productID := c.Param("productId")

	productId, err := service.GetProductByID(productID)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	if productId.ID != productID {
		return c.JSON(http.StatusForbidden, "forbidden access")
	}
	err = service.DeleteProduct(productID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error deleting product",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, payload.Response{
		Message: "success Delete Product",
		Data:    http.StatusOK,
	})
}

func GetAllProductsController(c echo.Context) error {
	_, err := middleware.ExtractUserID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "You must be re-login",
		})
	}
	products, err := service.GetAllProducts()
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, payload.Response{
		Message: "success",
		Data:    products,
	})
}
func UpdateProductStockController(c echo.Context) error {
	_, err := middleware.ExtractUserID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"Message": "you must be re login",
		})
	}

	payloadStock := payload.UpdateProductStockRequest{}
	c.Bind(&payloadStock)

	id := c.Param("productId")

	product, err := service.GetProductByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	if err := c.Validate(payloadStock); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error payload update product",
			"error":   err.Error(),
		})
	}

	err = service.UpdateProductStock(product, &payloadStock)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, payload.Response{
		Message: "success update product data",
		Data:    http.StatusOK,
	})

}
