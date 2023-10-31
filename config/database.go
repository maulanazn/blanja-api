package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnection() *gorm.DB {
	dsn := "host=" + GetDBHost() + "user=" + GetDBUser() + "password=" + GetDBPassword() + "dbname=" + GetDBName() + "port=" + GetDBPort() + "sslmode=" + GetDBSSLMode() + "TimeZone=" + GetDBTimezone()
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}
