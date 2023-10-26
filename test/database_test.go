package test

import (
	"belanjabackend/config"
	"belanjabackend/entity"
	"belanjabackend/webserver/helper"
	_ "embed"
	"fmt"
	"testing"
)

func TestGetConnection(t *testing.T) {
	config.GetConnection()

	fmt.Println("connected")
}

func TestCreateCustomerTable(t *testing.T) {
	config.GetConnection().AutoMigrate(entity.Customer{})
}

func TestEmbedEnv(t *testing.T) {
	fmt.Println(helper.ReadEnv("../.env"))
}
