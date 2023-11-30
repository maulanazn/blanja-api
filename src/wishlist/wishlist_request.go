package wishlist

type AddWishlistRequest struct {
  ProductId  string    `json:"product_id" validate:"required,ascii"`
  Quantity   int       `json:"quantity" validate:"required,numeric"`
}

type PutWishlistRequest struct {
  Quantity   int       `json:"quantity" validate:"required,numeric"`
}
