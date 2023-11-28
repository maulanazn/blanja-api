package users

import (
	"address"
	"time"
	"wishlist"
)

type Users struct {
	Id          string            `gorm:"type:varchar;unique;notNull;primaryKey;column:id;default:uuid_generate_v4()"`
	Address     []address.Address `gorm:"foreignKey:user_id;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Wishlist    wishlist.Wishlist `gorm:"foreignKey:wishlist_id;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserImage   string            `gorm:"type:varchar;column:user_image"`
	Username    string            `gorm:"type:varchar;column:user_name"`
	Email       string            `gorm:"type:varchar;unique;column:email"`
	Phone       int64             `gorm:"type:varchar;column:phone"`
	Gender      string            `gorm:"type:varchar;column:gender"`
	DateOfBirth string            `gorm:"type:varchar;column:dateofbirth"`
	Password    string            `gorm:"type:varchar;column:password"`
	Roles       string            `gorm:"type:varchar;column:roles"`
	CreatedAt   time.Time         `gorm:"type:timestamp;autoCreateTime;column:created_at"`
	UpdatedAt   time.Time         `gorm:"type:timestamp;column:updated_at"`
	DeletedAt   time.Time         `gorm:"type:timestamp;column:deleted_at"`
}
