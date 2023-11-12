package repository

import (
	"config"
	"context"
	"entity"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateProduct(ctx context.Context, product entity.Products) error {
	var categoryName entity.CategoryName
	var brandName entity.BrandName
	var colorName entity.ColorName
	var sizeName entity.SizeName

	ctx, timeout := context.WithTimeout(ctx, 8*time.Second)
	defer timeout()

	queryCreateProduct := config.MongoConnection().Database("maulanazn").Collection("products")
	_, err := queryCreateProduct.InsertOne(ctx, entity.Products{
		ProductId: primitive.NewObjectID(),
		UserId:    product.UserId,
		CategoryName: entity.CategoryName{
			Name:        categoryName.Name,
			Description: categoryName.Description,
		},
		Image:       product.Image,
		ProductName: product.ProductName,
		Brand: entity.BrandName{
			Name:        brandName.Name,
			Description: brandName.Description,
		},
		Rating: product.Rating,
		Price:  product.Price,
		Color: entity.ColorName{
			Name:        colorName.Name,
			Description: colorName.Description,
		},
		Size: entity.SizeName{
			Name:        sizeName.Name,
			Description: sizeName.Description,
		},
		Quantity: product.Quantity,
	})
	if err != nil {
		panic(err)
	}

	return nil
}
