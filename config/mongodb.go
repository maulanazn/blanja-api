package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func MongoConnection() *mongo.Client {
  // var viper *viper.Viper = util.LoadConfig(".", "blanja.yaml", "yaml")
	// uri := fmt.Sprintf("mongodb://%s:%s@%s:%d/?maxPoolSize=40", viper.GetString("database.mongodbuser"), viper.GetString("mongodbpassword"), viper.GetString("mongodbhost"), viper.GetInt("mongodbport"))
  uri := "mongodb://maulanazn:maulanazn123@localhost:27017/?maxPoolSize=40"
	if false {
		log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
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
