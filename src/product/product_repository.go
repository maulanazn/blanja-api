package product

import (
	"config"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateProduct(ctx context.Context, product Products) error {
	ctx, timeout := context.WithTimeout(ctx, 8*time.Second)
	defer timeout()

	queryCreateProduct := config.MongoConnection().Database("maulanazn").Collection("products")
	_, err := queryCreateProduct.InsertOne(ctx, Products{
		ProductId:    primitive.NewObjectID(),
		UserId:       product.UserId,
		CategoryName: product.CategoryName,
		BrandName:    product.BrandName,
		ColorName:    product.ColorName,
		SizeName:     product.SizeName,
		Image:        product.Image,
		ProductName:  product.ProductName,
		StoreName: product.StoreName,
		Rating:       product.Rating,
		Price:        product.Price,
		Quantity:     product.Quantity,
	})
	if err != nil {
		panic(err)
	}

	return nil
}

func UpdateProduct(ctx context.Context, id string, product Products) error {
	productid, productiderr := primitive.ObjectIDFromHex(id)
	if productiderr != nil {
		log.Println(productiderr)
	}
	ctx, timeout := context.WithTimeout(ctx, 8*time.Second)
	defer timeout()

	filterid := bson.M{"_id": productid}
	productdata := bson.D{
		{"$set", bson.D{
			{"category_name", product.CategoryName},
			{"brand_name", product.BrandName},
			{"color_name", product.ColorName},
			{"size_name", product.SizeName},
			{"image", product.Image},
			{"product_name", product.ProductName},
			{"store_name", product.StoreName},
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

func SelectUserProduct(ctx context.Context, userId string) *mongo.Cursor {
	ctx, timeout := context.WithTimeout(ctx, 8*time.Second)
	defer timeout()

	filter := bson.D{{"userid", userId}}

	queryUpdateProduct := config.MongoConnection().Database("maulanazn").Collection("products")
	cursor, err := queryUpdateProduct.Find(ctx, filter)
	if err != nil {
		panic(err)
	}

	return cursor
}

func SelectProduct(ctx context.Context, id string) *Products {
	productId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
	}

	ctx, timeout := context.WithTimeout(ctx, 8*time.Second)
	defer timeout()

	filter := bson.D{{"_id", productId}}

	var result *Products

	queryGetSingle := config.MongoConnection().Database("maulanazn").Collection("products")
	if err := queryGetSingle.FindOne(ctx, filter).Decode(&result); err != nil {
		log.Println(err)
	}

	return result
}
