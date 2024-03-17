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
		INSERT INTO products (name, price, image_url, stock, condition, tags, is_purchaseable, purchase_count)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id`,
		product.Name, product.Price, product.ImageURL, product.Stock,
		product.Condition, pq.Array(product.Tags), product.IsPurchaseable, product.PurchaseCount,
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
		SET name=$1, price=$2, image_url=$3, condition=$4, tags=$5, is_purchaseable=$6,
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
func GetProductByID(id string) (*models.Product, error) {
	var product models.Product
	err := config.DB.QueryRow(`
        SELECT id, name, price, image_url, stock, condition, tags, is_purchaseable, purchase_count
        FROM products
        WHERE id = $1`, id).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
		&product.ImageURL,
		&product.Stock,
		&product.Condition,
		pq.Array(&product.Tags),
		&product.IsPurchaseable,
		&product.PurchaseCount,
	)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func DeleteProduct(productID string) error {
	_, err := config.DB.Exec(`
        DELETE FROM products
        WHERE id = $1
    `, productID)

	if err != nil {
		return err
	}

	return nil
}
func GetAllProducts() ([]*models.Product, error) {
	rows, err := config.DB.Query(`
        SELECT id, name, price, image_url, stock, condition, tags, is_purchaseable, purchase_count
        FROM products
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*models.Product

	for rows.Next() {
		var product models.Product
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.ImageURL,
			&product.Stock,
			&product.Condition,

			pq.Array(&product.Tags),
			&product.IsPurchaseable,
			&product.PurchaseCount,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
func UpdateProductStock(product *models.Product) error {
	_, err := config.DB.Exec("UPDATE products SET stock = COALESCE($1, stock) WHERE id = $2", product.Stock, product.ID)
	if err != nil {
		return err
	}

	return nil
}

func UpdateProductAfterPay(product *models.Product) error {
	_, err := config.DB.Exec("UPDATE products SET stock = COALESCE($1, stock), is_purchaseable = COALESCE($2, is_purchaseable), purchase_count = $3 WHERE id = $4", product.Stock, product.IsPurchaseable, product.PurchaseCount, product.ID)
	if err != nil {
		return err
	}

	return nil
}
