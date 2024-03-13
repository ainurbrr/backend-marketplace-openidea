package service

import (
	"backend-marketplace-openidea/models"
	"backend-marketplace-openidea/models/payload"
	"backend-marketplace-openidea/repository/database"
)

func CreateProduct(req *payload.CreateProductRequest) (resp *payload.ProductResponse, err error) {
	newProduct := &models.Product{
		Name:           req.Name,
		Price:          req.Price,
		ImageURL:       req.ImageURL,
		Stock:          req.Stock,
		Condition:      req.Condition,
		Tags:           req.Tags,
		IsPurchaseable: req.IsPurchaseable,
	}

	err, _ = database.CreateProduct(newProduct)
	if err != nil {
		return nil, err
	}

	resp = &payload.ProductResponse{
		Name:           newProduct.Name,
		Price:          newProduct.Price,
		ImageURL:       newProduct.ImageURL,
		Stock:          newProduct.Stock,
		Condition:      newProduct.Condition,
		Tags:           newProduct.Tags,
		IsPurchaseable: newProduct.IsPurchaseable,
	}

	return resp, nil
}
