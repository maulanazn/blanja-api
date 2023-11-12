package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CategoryName struct {
	Name        string `json:"name" validate:"alphanum,alpha"`
	Description string `json:"description" validate:"alphanum,alpha"`
}

type BrandName struct {
	Name        string `json:"name" validate:"alphanum,alpha"`
	Description string `json:"description" validate:"alphanum,alpha"`
}

type ColorName struct {
	Name        string `json:"name"`
	Description string `json:"description" validate:"alphanum,alpha"`
}

type SizeName struct {
	Name        string `json:"name" validate:"alphanum,alpha"`
	Description string `json:"description" validate:"alphanum,alpha"`
}

type Products struct {
	ProductId    primitive.ObjectID `json:"id" bson:"_id"`
	UserId       string             `json:"user_id"`
	CategoryName CategoryName       `json:"category"`
	Image        string             `json:"image" validate:"uri"`
	ProductName  string             `json:"product_name" validate:"alphanum,alpha"`
	Brand        BrandName          `json:"brand"`
	Rating       int                `json:"rating" validate:"number"`
	Price        int                `json:"price" validate:"numeric"`
	Color        ColorName          `json:"color"`
	Size         SizeName           `json:"size"`
	Quantity     int                `json:"quantity" validate:"numeric"`
}
