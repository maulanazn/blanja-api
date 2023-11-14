package repository

import (
	"config"
	"context"
	"entity"
	"time"

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

func CreateCategory(ctx context.Context, category entity.CategoryName) error {
	ctx, timeout := context.WithTimeout(ctx, 8*time.Second)
	defer timeout()

	queryCategoryProduct := config.MongoConnection().Database("maulanazn").Collection("category")
	_, err := queryCategoryProduct.InsertOne(ctx, entity.CategoryName{
		CategoryId:  primitive.NewObjectID(),
		UserId:      category.UserId,
		ProductId:   category.ProductId,
		Name:        category.Name,
		Description: category.Description,
	})
	if err != nil {
		panic(err)
	}

	return nil
}

func CreateBrand(ctx context.Context, brand entity.BrandName) error {
	ctx, timeout := context.WithTimeout(ctx, 8*time.Second)
	defer timeout()

	queryBrandProduct := config.MongoConnection().Database("maulanazn").Collection("brand")
	_, err := queryBrandProduct.InsertOne(ctx, entity.BrandName{
		CategoryId:  primitive.NewObjectID(),
		UserId:      brand.UserId,
		ProductId:   brand.ProductId,
		Name:        brand.Name,
		Description: brand.Description,
	})
	if err != nil {
		panic(err)
	}

	return nil
}

func CreateColor(ctx context.Context, color entity.ColorName) error {
	ctx, timeout := context.WithTimeout(ctx, 8*time.Second)
	defer timeout()

	queryColorProduct := config.MongoConnection().Database("maulanazn").Collection("color")
	_, err := queryColorProduct.InsertOne(ctx, entity.ColorName{
		CategoryId:  primitive.NewObjectID(),
		UserId:      color.UserId,
		ProductId:   color.ProductId,
		Name:        color.Name,
		Description: color.Description,
	})
	if err != nil {
		panic(err)
	}

	return nil
}

func CreateSize(ctx context.Context, size entity.SizeName) error {
	ctx, timeout := context.WithTimeout(ctx, 8*time.Second)
	defer timeout()

	querySizeProduct := config.MongoConnection().Database("maulanazn").Collection("size")
	_, err := querySizeProduct.InsertOne(ctx, entity.SizeName{
		CategoryId:  primitive.NewObjectID(),
		UserId:      size.UserId,
		ProductId:   size.ProductId,
		Name:        size.Name,
		Description: size.Description,
	})
	if err != nil {
		panic(err)
	}

	return nil
}
