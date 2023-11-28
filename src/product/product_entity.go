package product

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Products struct {
	ProductId    primitive.ObjectID `bson:"_id"`
	UserId       string
	CategoryName string
	BrandName    string
	ColorName    string
	SizeName     string
	Image        string
	ProductName  string
	Rating       int
	Price        int
	Quantity     int
}
