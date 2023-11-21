package config

import (
	"context"
	"fmt"
	"log"
	"util"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnection() *mongo.Client {
  var viper *viper.Viper = util.LoadConfig(".", "blanja.yaml", "yaml")
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d/?maxPoolSize=40", viper.GetString("database.mongodbuser"), viper.GetString("mongodbpassword"), viper.GetString("mongodbhost"), viper.GetInt("mongodbport"))
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

	return client
}
