package product_test

import (
	"context"
	"fmt"
	"log"
	"product"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func TestGetUserProduct(t *testing.T) {
	cursor := product.SelectUserProduct(context.Background(), "39a66aa8-9db3-4f3b-9e92-a714308b601a")

	var result []bson.M
	if err := cursor.All(context.Background(), &result); err != nil {
		log.Println(err.Error())
	}

	for _, data := range result {
		fmt.Println(data)
	}
}
