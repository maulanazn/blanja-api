package repository

import (
	"config"
	"context"
	"entity"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateProduct(ctx context.Context, product entity.Products) error {
	ctx, timeout := context.WithTimeout(ctx, 8*time.Second)
	defer timeout()

	queryCreateProduct := config.MongoConnection().Database("maulanazn").Collection("products")
	_, err := queryCreateProduct.InsertOne(ctx, entity.Products{
		ProductId:   primitive.NewObjectID(),
		UserId:      product.UserId,
		Image:       product.Image,
		ProductName: product.ProductName,
		Rating:      product.Rating,
		Price:       product.Price,
		Quantity:    product.Quantity,
	})
	if err != nil {
		panic(err)
	}

	return nil
}

func UpdateProduct(ctx context.Context, id byte, product entity.Products) error {
	ctx, timeout := context.WithTimeout(ctx, 8*time.Second)
	defer timeout()

	filterid := bson.D{{"_id", id}}

	queryUpdateProduct := config.MongoConnection().Database("maulanazn").Collection("products")
	_, err := queryUpdateProduct.UpdateOne(ctx, filterid, entity.Products{
		Image:       product.Image,
		ProductName: product.ProductName,
		Rating:      product.Rating,
		Price:       product.Price,
		Quantity:    product.Quantity,
	})
	if err != nil {
		panic(err)
	}
	return nil
}

func SelectUserProduct(ctx context.Context, product entity.Products) error {
	ctx, timeout := context.WithTimeout(ctx, 8*time.Second)
	defer timeout()

	queryUpdateProduct := config.MongoConnection().Database("maulanazn").Collection("products")
	_, err := queryUpdateProduct.Find(ctx, entity.Products{
		Image:       product.Image,
		ProductName: product.ProductName,
		Rating:      product.Rating,
		Price:       product.Price,
		Quantity:    product.Quantity,
	})
	if err != nil {
		panic(err)
	}
	return nil
}
