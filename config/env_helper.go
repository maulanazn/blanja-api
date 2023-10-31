package config

import "github.com/spf13/viper"

func GetConfig() *viper.Viper {
	config := viper.New()

	config.SetConfigFile("config.env")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	if err != nil {
		panic(err)
	}

	return config
}
