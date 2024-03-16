package service

import (
	"backend-marketplace-openidea/models"
	"backend-marketplace-openidea/models/payload"
	"backend-marketplace-openidea/repository/database"
	"errors"
)

func CreateProduct(req *payload.CreateProductRequest) (resp *payload.ProductResponse, err error) {
	updateProduct := &models.Product{
		Name:           req.Name,
		Price:          req.Price,
		ImageURL:       req.ImageURL,
		Stock:          req.Stock,
		Condition:      req.Condition,
		Tags:           req.Tags,
		IsPurchaseable: req.IsPurchaseable,
	}

	err, _ = database.CreateProduct(updateProduct)
	if err != nil {
		return nil, err
	}

	resp = &payload.ProductResponse{
		Name:           updateProduct.Name,
		Price:          updateProduct.Price,
		ImageURL:       updateProduct.ImageURL,
		Stock:          updateProduct.Stock,
		Condition:      updateProduct.Condition,
		Tags:           updateProduct.Tags,
		IsPurchaseable: updateProduct.IsPurchaseable,
	}

	return resp, nil
}

func UpdateProduct(productID string, req *payload.UpdateProductRequest) (resp *payload.ProductUpdateResponse, err error) {
	updateProduct := &models.Product{
		Name:           req.Name,
		Price:          req.Price,
		ImageURL:       req.ImageURL,
		Condition:      req.Condition,
		Tags:           req.Tags,
		IsPurchaseable: req.IsPurchaseable,
	}

	err, _ = database.UpdateProduct(productID, updateProduct)
	if err != nil {
		return nil, err
	}

	resp = &payload.ProductUpdateResponse{
		Name:           updateProduct.Name,
		Price:          updateProduct.Price,
		ImageURL:       updateProduct.ImageURL,
		Condition:      updateProduct.Condition,
		Tags:           updateProduct.Tags,
		IsPurchaseable: updateProduct.IsPurchaseable,
	}

	return resp, nil
}

func GetProductByID(id string) (*models.Product, error) {
	product, err := database.GetProductByID(id)
	if err != nil {
		return nil, errors.New("Product not found")
	}
	return product, nil
}
func DeleteProduct(productID string) error {
	err := database.DeleteProduct(productID)
	if err != nil {
		return err
	}
	return nil
}
func GetAllProducts() ([]*models.Product, error) {
	products, err := database.GetAllProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}
func UpdateProductStock(product *models.Product, req *payload.UpdateProductStockRequest) error {
	product.Stock = req.Stock

	err := database.UpdateProductStock(product)
	if err != nil {
		return err
	}
	return nil
}
