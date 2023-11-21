package repository_test

import (
	"context"
	"fmt"
	"log"
	"repository"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func TestGetUserProduct(t *testing.T) {
	cursor, err := repository.SelectUserProduct(context.Background(), "39a66aa8-9db3-4f3b-9e92-a714308b601a")
	if err != nil {
		log.Println(err)	
	}
	
	var result []bson.M
	if err = cursor.All(context.Background(), &result); err != nil {
		log.Println(err)
	} 

	for _, data := range result {
		fmt.Println(data)
	}
}
