package entity

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	UserImage   string `json:"userimage"`
	UserName    string `json:"username"`
	Email       string `gorm:"unique"`
	Phone       int    `json:"phone"`
	Gender      string `json:"gender"`
	DateofBirth string `json:"dateofbirth"`
	Password    string `json:"password"`
	CreatedAt   string `json:"createdat"`
	UpdatedAt   string `json:"updatedat"`
	DeletedAt   string `json:"deletedat"`
}
