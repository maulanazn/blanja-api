package repository_test

import (
	"config"
	"context"
	"database/sql"
	"encoding/json"
	"entity"
	"errors"
	"fmt"
	"log"
	"repository"
	"request"
	"testing"
	"util"
)

func TestCreateTableCustomer(t *testing.T) {
	err := config.GetConnection().AutoMigrate(entity.Users{})
	util.PanicIfError(err)
}

func TestCreateTableAddress(t *testing.T) {
	err := config.GetConnection().AutoMigrate(entity.Address{})
	util.PanicIfError(err)
}

func TestGetPassword(t *testing.T) {
	customerRequest := request.LoginRequest{
		Email:    "test11@mail.com",
		Password: "tes123",
	}
	result, resultErr := repository.SelectEmailCustomers(context.Background(), customerRequest.Email)
	util.PanicIfError(resultErr)

	fmt.Println(result.Password)
}

func TestSelectUserById(t *testing.T) {
	id := "3fa38705ba084a93b299f387eccf89ab"

	result, err := repository.SelectCustomerById(context.Background(), id)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestGetAndVerifyPassword(t *testing.T) {
	customerRequest := request.LoginRequest{
		Email:    "test11@mail.com",
		Password: "tes123",
	}
	result, resultErr := repository.SelectEmailCustomers(context.Background(), customerRequest.Email)
	util.PanicIfError(resultErr)

	if err := util.ComparePasswords([]byte(result.Password), []byte(customerRequest.Password)); err != nil {
		log.Fatal(err)
		return
	}

	log.Println(result)
}

func TestEmailExist(t *testing.T) {
	customerRequest := request.LoginRequest{
		Email: "maulanazn19@mail.com",
	}

	customerData := entity.Users{
		Email: "maulanazn19@mail.com",
	}

	result, _ := repository.SelectEmailCustomers(context.Background(), customerRequest.Email)

	if customerRequest.Email != customerData.Email {
		fmt.Println(errors.New("lkjasdf"))
		return
	}

	fmt.Println(result.Roles)
}

func TestGetPasswordPlain(t *testing.T) {
	var data map[string]interface{}
	config.GetConnection().Table("customers").Take(&data).Select("*").Where("email = @email", sql.Named("email", "test11@mail.com"))

	fmt.Println(data["user_name"])
}

func TestUpdateAndGetCustomer(t *testing.T) {
	var users entity.Users
	repository.UpdateCustomer(context.Background(), users, "12312").Error()

	fmt.Println("ok")
}

func TestGetAddressByUser(t *testing.T) {
	var result []entity.Address

	result, err := repository.AddressByUser(context.Background(), "0151cf30fbd5456aa30a3e5af3ccba18")
	util.PanicIfError(err)

	for _, data := range result {
		jsonresult, err := json.MarshalIndent(&data, "", "")
		util.PanicIfError(err)
		fmt.Println(string(jsonresult))
	}
}

func TestGetAddressById(t *testing.T) {
	var result entity.Address

	result, resulterr := repository.AddressById(context.Background(), "51ac602e02534e6a813b96c509b9b429")
	util.PanicIfError(resulterr)

	jsonresult, err := json.MarshalIndent(&result, "", "")
	util.PanicIfError(err)
	fmt.Println(string(jsonresult))
}

func TestUpdateAddress(t *testing.T) {
	var result entity.Address

	repository.UpdateAddress(context.Background(), result, "51ac602e02534e6a813b96c509b9b429").Error()

	fmt.Println(result)
}
