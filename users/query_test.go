package users_test

import (
	"address"
	"config"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"testing"
	"users"
	"util"
)

func TestCreateTableCustomer(t *testing.T) {
	err := config.GetConnection().AutoMigrate(users.Users{})
	util.PanicIfError(err)
}

func TestCreateTableAddress(t *testing.T) {
	err := config.GetConnection().AutoMigrate(address.Address{})
	util.PanicIfError(err)
}

func TestGetPassword(t *testing.T) {
	customerRequest := users.LoginRequest{
		Email:    "test11@mail.com",
		Password: "tes123",
	}
	result, resultErr := users.SelectEmailCustomers(context.Background(), customerRequest.Email)
	util.PanicIfError(resultErr)

	fmt.Println(result.Password)
}

func TestSelectUserById(t *testing.T) {
	id := "3fa38705ba084a93b299f387eccf89ab"

	result, err := users.SelectCustomerById(context.Background(), id)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestGetAndVerifyPassword(t *testing.T) {
	customerRequest := users.LoginRequest{
		Email:    "test11@mail.com",
		Password: "tes123",
	}
	result, resultErr := users.SelectEmailCustomers(context.Background(), customerRequest.Email)
	util.PanicIfError(resultErr)

	if err := util.ComparePasswords([]byte(result.Password), []byte(customerRequest.Password)); err != nil {
		log.Fatal(err)
		return
	}

	log.Println(result)
}

func TestEmailExist(t *testing.T) {
	customerRequest := users.LoginRequest{
		Email: "maulanazn19@mail.com",
	}

	customerData := users.Users{
		Email: "maulanazn19@mail.com",
	}

	result, _ := users.SelectEmailCustomers(context.Background(), customerRequest.Email)

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
	var userEntity users.Users
	users.UpdateCustomer(context.Background(), userEntity, "12312").Error()

	fmt.Println("ok")
}

func TestGetAddressByUser(t *testing.T) {
	var result []address.Address

	result, err := address.AddressByUser(context.Background(), "0151cf30fbd5456aa30a3e5af3ccba18")
	util.PanicIfError(err)

	for _, data := range result {
		jsonresult, err := json.MarshalIndent(&data, "", "")
		util.PanicIfError(err)
		fmt.Println(string(jsonresult))
	}
}

func TestGetAddressById(t *testing.T) {
	var result address.Address

	result, resulterr := address.AddressById(context.Background(), "51ac602e02534e6a813b96c509b9b429")
	util.PanicIfError(resulterr)

	jsonresult, err := json.MarshalIndent(&result, "", "")
	util.PanicIfError(err)
	fmt.Println(string(jsonresult))
}

func TestUpdateAddress(t *testing.T) {
	var result address.Address

	address.UpdateAddress(context.Background(), result, "51ac602e02534e6a813b96c509b9b429").Error()

	fmt.Println(result)
}
