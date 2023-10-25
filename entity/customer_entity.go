package entity

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Userimage   string `json:"userimage"`
	Username    string `json:"username"`
	Email       string `gorm:"unique"`
	Phone       int    `json:"phone"`
	Gender      string `json:"gender"`
	Dateofbirth string `json:"dateofbirth"`
	Password    string `json:"password"`
	CreatedAt   string `json:"createdat"`
	UpdatedAt   string `json:"updatedat"`
	DeletedAt   string `json:"deletedat"`
}
