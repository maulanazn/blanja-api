package config

import (
	"fmt"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestGetConnection(t *testing.T) {
	_, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=maulanazn password=t00r123 dbname=paybook port=5432 sslmode=disable TimeZone=Asia/Jakarta",
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("connected")
}
