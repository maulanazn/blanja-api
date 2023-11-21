package repository

import (
	"config"
	"context"
	"entity"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func UpdateProduct(ctx context.Context, id string, product entity.Products) error {
  productid, productiderr := primitive.ObjectIDFromHex(id)
  if productiderr != nil {
    log.Println(productiderr)
  }
	ctx, timeout := context.WithTimeout(ctx, 8*time.Second)
	defer timeout()

  filterid := bson.M{"_id": productid}
  productdata := bson.D{
    {"$set", bson.D{
      {"image", product.Image}, 
      {"productname", product.ProductName},  
      {"rating", product.Rating}, 
      {"price", product.Price}, 
      {"quantity", product.Quantity},
  }}} 

	queryUpdateProduct := config.MongoConnection().Database("maulanazn").Collection("products")
  _, err := queryUpdateProduct.UpdateOne(ctx, filterid, productdata)
	if err != nil {
		panic(err)
	}
	return nil
}

func SelectUserProduct(ctx context.Context, user_id string) *mongo.Cursor {
	ctx, timeout := context.WithTimeout(ctx, 8*time.Second)
	defer timeout()

	filter := bson.D{{"userid", user_id}}

	queryUpdateProduct := config.MongoConnection().Database("maulanazn").Collection("products")
	cursor, err := queryUpdateProduct.Find(ctx, filter)
  if err != nil {
		panic(err)
	}

	return cursor
}

func SelectProduct(ctx context.Context, id string) entity.Products {
  productid, err := primitive.ObjectIDFromHex(id)
  if err != nil {
    log.Println(err)
  }

	ctx, timeout := context.WithTimeout(ctx, 8*time.Second)
	defer timeout()

	filter := bson.D{{"_id", productid}}
  var result entity.Products

	queryGetSingle := config.MongoConnection().Database("maulanazn").Collection("products")
	queryGetSingle.FindOne(ctx, filter).Decode(&result)

  return result
}
