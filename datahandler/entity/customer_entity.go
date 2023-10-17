package entity

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	CustomerId  int       `gorm:"primaryKey;autoIncrement;unique" json:"customer_id"`
	UserImage   string    `json:"user_image"`
	UserName    string    `json:"user_name"`
	Email       string    `json:"email"`
	Phone       int       `json:"phone"`
	Gender      string    `json:"gender"`
	DateofBirth time.Time `json:"date_of_birth" gorm:"autoUpdateTime"`
	Password    string    `json:"password"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoUpdateTime"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}
