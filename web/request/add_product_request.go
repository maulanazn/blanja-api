package request

type AddProductRequest struct {
	Image       string `json:"image" validate:"uri"`
	ProductName string `json:"product_name"`
	Rating      int    `json:"rating" validate:"number"`
	Price       int    `json:"price" validate:"numeric"`
	Quantity    int    `json:"quantity" validate:"numeric"`
}
