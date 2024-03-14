package database

import (
	"backend-marketplace-openidea/config"
	"backend-marketplace-openidea/models"
	"fmt"
	"log"

	"github.com/lib/pq"
)

func CreateProduct(product *models.Product) (error, string) {
	var id string

	err := config.DB.QueryRow(`
		INSERT INTO products (name, price, image_url, stock, condition, tags, is_purchaseable)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id`,
		product.Name, product.Price, product.ImageURL, product.Stock,
		product.Condition, pq.Array(product.Tags), product.IsPurchaseable,
	).Scan(&id)

	if err != nil {
		log.Printf("Unable to execute the query. %v", err)
		return err, id
	}

	fmt.Printf("Inserted a single record with ID: %v", id)
	return nil, id
}
func UpdateProduct(productID string, product *models.Product) (error, string) {

	err := config.DB.QueryRow(`
		UPDATE products
		SET name=$1, price=$2, image_url=$3, condition=$4, tags=$5, is_purchaseable=$6
		WHERE id=$7
		RETURNING id`,
		product.Name, product.Price, product.ImageURL,
		product.Condition, pq.Array(product.Tags), product.IsPurchaseable,
		productID,
	).Scan(&productID)

	if err != nil {
		log.Printf("Unable to execute the query. %v", err)
		return err, ""
	}

	fmt.Printf("Updated a single record with ID: %v", productID)
	return nil, productID
}
