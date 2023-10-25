package config

import (
	"belanjabackend/entity"
	_ "embed"
	"fmt"
	"os"
	"testing"
)

func TestGetConnection(t *testing.T) {
	GetConnection()

	fmt.Println("connected")
}

func TestCreateCustomerTable(t *testing.T) {
	GetConnection().AutoMigrate(entity.Customer{})
}

func TestEmbedEnv(t *testing.T) {
	file, _ := os.ReadFile("../.env")

	fmt.Println(string(file[8:]))
}
