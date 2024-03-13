package database

import (
	"backend-marketplace-openidea/config"
	"backend-marketplace-openidea/models"
	"fmt"
	"log"

	"github.com/lib/pq"
)

func CreateProduct(product *models.Product) (error, uint) {
	var id uint

	err := config.DB.QueryRow(`
		INSERT INTO Product (name, price, imageUrl, stock, condition, tags, isPurchaseable)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id`,
		product.Name, product.Price, product.ImageURL, product.Stock,
		product.Condition, pq.Array(product.Tags), product.IsPurchaseable,
	).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
		return err, 0
	}

	fmt.Printf("Inserted a single record with ID: %v", id)
	return nil, id
}
