package order_details

type OrderDetailRequest struct {
	AddressId string `json:"address_id" validate:"required,ascii"`
	ProductId string `json:"product_id" validate:"required,ascii"`
	StoreName string `json:"store_name" validate:"required,ascii"`
	Quantity  int    `json:"quantity" validate:"required,numeric"`
	Price     int    `json:"price" validate:"required,numeric"`
}
