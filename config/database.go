package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnection() *gorm.DB {
	dsn := "host=" + GetConfig().GetString("DB_HOST") + " " + "user=" + GetConfig().GetString("DB_USER") + " " + "password=" + GetConfig().GetString("DB_PASSWORD") + " " + "dbname=" + GetConfig().GetString("DB_NAME") + " " + "port=" + GetConfig().GetString("DB_PORT") + " " + "sslmode=" + GetConfig().GetString("DB_SSLMODE") + " " + "TimeZone=" + GetConfig().GetString("DB_TIMEZONE")

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}
