package config

import (
	"context"
	"github.com/spf13/viper"
	"log"
	"os"
	"util"

	"github.com/plutov/paypal/v4"
)

func GetPaypalConfig() *paypal.Client {
	var viper *viper.Viper = util.LoadConfig(".", "blanja.yaml", "yaml")
	conn, connErr := paypal.NewClient(viper.GetString("paypal_clientid"), viper.GetString("paypal_secret"), paypal.APIBaseSandBox)
	if connErr != nil {
		log.Println(connErr)
	}

	conn.SetLog(os.Stdout)

	_, err := conn.GetAccessToken(context.TODO())
	if err != nil {
		log.Println(err)
	}

	return conn
}
