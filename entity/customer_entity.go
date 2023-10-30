package entity

import (
	"time"
)

type Customer struct {
	Id          string `gorm:"primaryKey"`
	Address     []Address
	Userimage   string    `json:"userimage"`
	Username    string    `json:"username"`
	Email       string    `gorm:"unique"`
	Phone       int64     `json:"phone"`
	Gender      string    `json:"gender"`
	Dateofbirth string    `json:"dateofbirth"`
	Password    string    `json:"password"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
