package config

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"util"
)

func MongoConnection() *mongo.Client {
	var viper *viper.Viper = util.LoadConfig(".", "blanja.yaml", "yaml")
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d/?maxPoolSize=40", viper.GetString("database.mongodbuser"), viper.GetString("database.mongodbpassword"), viper.GetString("database.mongodbhost"), viper.GetInt("database.mongodbport"))

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	pingErr := client.Ping(context.TODO(), nil)
	if pingErr != nil {
		log.Println(pingErr.Error())
	}

	return client
}
