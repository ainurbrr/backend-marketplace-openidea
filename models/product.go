package models

type Product struct {
	ID             string   `json:"Productid"`
	Name           string   `json:"name" validate:"required,min=5,max=60"`
	Price          int      `json:"price" validate:"required,min=0"`
	ImageURL       string   `json:"imageUrl" validate:"required,url"`
	Stock          int      `json:"stock" validate:"required,min=0"`
	Condition      string   `json:"condition" validate:"required,oneof='new' 'second'"`
	Tags           []string `json:"tags" validate:"required,min=0,dive,min=1"`
	IsPurchaseable bool     `json:"isPurchaseable" validate:"required"`
	PurchaseCount  int      `json:"purchaseCount"`
}
