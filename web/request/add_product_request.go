package request

type AddProductRequest struct {
	Image       string `json:"image" validate:"required,uri|image"`
	ProductName string `json:"product_name" validate:"required,alphanum|alpha,ascii"`
	Rating      int    `json:"rating" validate:"required,number"`
	Price       int    `json:"price" validate:"required,numeric"`
	Quantity    int    `json:"quantity" validate:"required,numeric"`
}
