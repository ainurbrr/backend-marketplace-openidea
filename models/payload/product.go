package payload

type CreateProductRequest struct {
	Name           string   `json:"name"`
	Price          int      `json:"price"`
	ImageURL       string   `json:"imageUrl"`
	Stock          int      `json:"stock"`
	Condition      string   `json:"condition"`
	Tags           []string `json:"tags"`
	IsPurchaseable bool     `json:"isPurchaseable"`
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
