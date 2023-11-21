package util

import (
	"log"

	"github.com/spf13/viper"
)

func LoadConfig(path string, filename string, typefile string) *viper.Viper {
  var viper *viper.Viper = viper.New()
  viper.SetConfigName(filename)
  viper.SetConfigType(typefile)
  viper.AddConfigPath(path)

  err := viper.ReadInConfig()
  if err != nil {
    log.Fatal(err.Error())
  }

  return viper
}
