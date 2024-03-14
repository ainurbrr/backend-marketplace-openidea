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
