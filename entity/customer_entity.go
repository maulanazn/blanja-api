package entity

import (
	"time"
)

type Users struct {
	Id          string    `gorm:"primaryKey"`
	Address     []Address `gorm:"foreignKey:user_id;references:id"`
	Userimage   string    `json:"userimage"`
	Username    string    `json:"username"`
	Email       string    `gorm:"unique"`
	Phone       int64     `json:"phone"`
	Gender      string    `json:"gender"`
	Dateofbirth string    `json:"dateofbirth"`
	Password    string    `json:"password"`
	Roles       string    `json:"roles"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
