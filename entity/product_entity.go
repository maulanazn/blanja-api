package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CategoryName struct {
	CategoryId  primitive.ObjectID `bson:"_id"`
	UserId      string
	ProductId   string
	Name        string
	Description string
}

type BrandName struct {
	CategoryId  primitive.ObjectID `bson:"_id"`
	UserId      string
	ProductId   string
	Name        string
	Description string
}

type ColorName struct {
	CategoryId  primitive.ObjectID `bson:"_id"`
	UserId      string
	ProductId   string
	Name        string
	Description string
}

type SizeName struct {
	CategoryId  primitive.ObjectID `bson:"_id"`
	UserId      string
	ProductId   string
	Name        string
	Description string
}

type Products struct {
	ProductId   primitive.ObjectID `bson:"_id"`
	UserId      string
	Image       string
	ProductName string
	Rating      int
	Price       int
	Quantity    int
}
