package entity

import (
	"time"
)

type Address struct {
	Id             string `gorm:"primaryKey"`
	CustomerId     string
	AddressType    string
	RecipientName  string
	RecipientPhone string
	AddressName    string
	PostalCode     string
	City           string
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time
	DeletedAt      time.Time
}
