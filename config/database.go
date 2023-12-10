package config

import (
	"fmt"
	"util"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnection() *gorm.DB {
	var viper *viper.Viper = util.LoadConfig(".", "blanja.yaml", "yaml")
	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s", viper.GetString("database.dbhost"), viper.GetString("database.dbuser"), viper.GetString("database.dbpassword"), viper.GetString("database.dbname"), viper.GetInt("database.dbport"), viper.GetString("database.dbsslmode"), viper.GetString("database.dbtimezone"))
	dsn := fmt.Sprintf("%s", viper.GetString("database.url"))

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		PrepareStmt: true,
	})

	if err != nil {
		panic(err)
	}

	return db
}
