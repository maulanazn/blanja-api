package order_details

import "time"

type OrderDetail struct {
	OrderId   int       `gorm:"type:serial;unique;notNull;primaryKey;column:order_id"`
	UserId    string    `gorm:"type:varchar;column:user_id"`
	AddressId string    `gorm:"type:varchar;column:address_id"`
	ProductId string    `gorm:"type:varchar;column:product_id"`
	StoreName string    `gorm:"type:varchar;column:store_name"`
	Quantity  int       `gorm:"type:int;column:quantity"`
	Price     int       `gorm:"type:int;column:price"`
	SubTotal  int       `gorm:"type:int;column:subtotal"`
	CreatedAt time.Time `gorm:"type:timestamp;column:created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp;column:updated_at"`
	DeletedAt time.Time `gorm:"type:timestamp;column:deleted_at"`
}
