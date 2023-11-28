package address

import (
	"time"
)

type Address struct {
	Id             string    `gorm:"type:varchar;unique;notNull;primaryKey;column:id;default:uuid_generate_v4()"`
	UserId         string    `gorm:"type:varchar;column:user_id"`
	AddressType    string    `gorm:"type:varchar;column:address_type"`
	RecipientName  string    `gorm:"type:varchar;column:recipient_name"`
	RecipientPhone string    `gorm:"type:varchar;column:recipient_phone"`
	AddressName    string    `gorm:"type:varchar;column:address_name"`
	PostalCode     string    `gorm:"type:varchar;column:postal_code"`
	City           string    `gorm:"type:varchar;column:city"`
	CreatedAt      time.Time `gorm:"type:timestamp;autoCreateTime;column:created_at"`
	UpdatedAt      time.Time `gorm:"type:timestamp;column:updated_at"`
	DeletedAt      time.Time `gorm:"type:timestamp;column:deleted_at"`
}
