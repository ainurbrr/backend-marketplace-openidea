package payload

type CreateProductRequest struct {
	Name           string   `json:"name" validate:"required,min=5,max=60"`
	Price          int      `json:"price" validate:"required,min=0"`
	ImageURL       string   `json:"imageUrl" validate:"required,url"`
	Stock          int      `json:"stock" validate:"required,min=0"`
	Condition      string   `json:"condition" validate:"required,eq=new|eq=second"`
	Tags           []string `json:"tags" validate:"required,min=0,dive,min=1"`
	IsPurchaseable bool     `json:"isPurchaseable" validate:"required"`
}
type ProductResponse struct {
	Name           string   `json:"name"`
	Price          int      `json:"price"`
	ImageURL       string   `json:"imageUrl"`
	Stock          int      `json:"stock"`
	Condition      string   `json:"condition"`
	Tags           []string `json:"tags"`
	IsPurchaseable bool     `json:"isPurchaseable"`
}

type UpdateProductRequest struct {
	Name           string   `json:"name" validate:"required,min=5,max=60"`
	Price          int      `json:"price" validate:"required,min=0"`
	ImageURL       string   `json:"imageUrl" validate:"required,url"`
	Condition      string   `json:"condition" validate:"required,eq=new|eq=second"`
	Tags           []string `json:"tags" validate:"required,min=0,dive,min=1"`
	IsPurchaseable bool     `json:"isPurchaseable" validate:"required"`
}

type ProductUpdateResponse struct {
	Name           string   `json:"name"`
	Price          int      `json:"price"`
	ImageURL       string   `json:"imageUrl"`
	Condition      string   `json:"condition"`
	Tags           []string `json:"tags"`
	IsPurchaseable bool     `json:"isPurchaseable"`
}

type GetProduct struct {
	ProductID     string   `json:"productId"`
	Name          string   `json:"name"`
	Price         int      `json:"price"`
	ImageURL      string   `json:"imageUrl"`
	Stock         int      `json:"stock"`
	Condition     string   `json:"condition"`
	Tags          []string `json:"tags"`
	IsPurchasable bool     `json:"isPurchasable"`
	PurchaseCount int      `json:"purchaseCount"`
}

type UpdateProductStockRequest struct {
	Stock int `json:"stock" validate:required"`
}
