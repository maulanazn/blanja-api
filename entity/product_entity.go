package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type CategoryName struct {
	Name        string
	Description string
}

type BrandName struct {
	Name        string
	Description string
}

type ColorName struct {
	Name        string
	Description string
}

type SizeName struct {
	Name        string
	Description string
}

type Products struct {
	ProductId    primitive.ObjectID `bson:"_id"`
	UserId       string
	CategoryName []CategoryName
	Image        string
	ProductName  string
	Brand        []BrandName
	Rating       int
	Price        int
	Color        []ColorName
	Size         []SizeName
	Quantity     int
}
