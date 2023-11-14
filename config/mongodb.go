package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnection() *mongo.Client {
	// uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/?maxPoolSize=15", os.Getenv("MONGO_DB_USER"), os.Getenv("MONGO_DB_PASS"), os.Getenv("MONGO_DB_HOST"), os.Getenv("MONGO_DB_PORT"))
	uri := "mongodb://maulanazn:maulanazn123@localhost:27017/?maxPoolSize=15"
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	pingerr := client.Ping(context.TODO(), nil)
	if pingerr != nil {
		log.Fatal(pingerr)
	}

	fmt.Println("Connected")
	return client
}
