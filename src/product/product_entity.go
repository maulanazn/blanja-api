package product

import "wishlist"

type Products struct {
	ProductId         string              `gorm:"type:varchar;unique;notNull;primaryKey;column:product_id;default:uuid_generate_v4()"`
	UserId            string              `gorm:"type:varchar;column:user_id"`
	Wishlist          []wishlist.Wishlist `gorm:"foreignKey:product_id;references:product_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	StoreNameWishlist []wishlist.Wishlist `gorm:"foreignKey:store_name;references:store_name;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Image             string              `gorm:"type:varchar;column:image"`
	ProductName       string              `gorm:"type:varchar;column:product_name"`
	StoreName         string              `gorm:"type:varchar;column:store_name;unique;notNull"`
	Rating            int                 `gorm:"type:int;column:rating"`
	Price             int                 `gorm:"type:int;column:price"`
	Quantity          int                 `gorm:"type:int;column:quantity"`
}
