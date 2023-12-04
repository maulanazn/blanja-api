package wishlist

import "time"

type Wishlist struct {
	WishlistId string    `gorm:"type:varchar;unique;notNull;primaryKey;column:wishlist_id;default:uuid_generate_v4()"`
	UserId     string    `gorm:"type:varchar;column:user_id"`
	ProductId  string    `gorm:"type:varchar;column:product_id"`
	StoreName  string    `gorm:"type:varchar;column:store_name"`
	Quantity   int       `gorm:"type:int;column:quantity"`
	CreatedAt  time.Time `gorm:"type:timestamp;autoCreateTime;column:created_at"`
	UpdatedAt  time.Time `gorm:"type:timestamp;column:updated_at"`
	DeletedAt  time.Time `gorm:"type:timestamp;column:deleted_at"`
}
