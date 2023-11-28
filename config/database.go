package config

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"util"
)

func GetConnection() *gorm.DB {
	var viper *viper.Viper = util.LoadConfig(".", "blanja.yaml", "yaml")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s", viper.GetString("database.dbhost"), viper.GetString("database.dbuser"), viper.GetString("database.dbpassword"), viper.GetString("database.dbname"), viper.GetInt("database.dbport"), viper.GetString("database.dbsslmode"), viper.GetString("database.dbtimezone"))

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}
