package entity

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	UserImage   string `json:"user_image"`
	UserName    string `json:"username"`
	Email       string `json:"email"`
	Phone       int    `json:"phone"`
	Gender      string `json:"gender"`
	DateofBirth string `json:"date_of_birth"`
	Password    string `json:"password"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
}
