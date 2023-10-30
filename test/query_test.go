package test

import (
	"belanjabackend/config"
	"belanjabackend/entity"
	"belanjabackend/repository"
	"belanjabackend/webserver/helper"
	"belanjabackend/webserver/request"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"testing"
	"time"
)

func TestCreateTableCustomer(t *testing.T) {
	err := config.GetConnection().AutoMigrate(entity.Customer{})
	helper.PanicIfError(err)
}

func TestCreateTableAddress(t *testing.T) {
	err := config.GetConnection().AutoMigrate(entity.Address{})
	helper.PanicIfError(err)
}

func TestGetPassword(t *testing.T) {
	customerRequest := request.LoginRequest{
		Email:    "test11@mail.com",
		Password: "tes123",
	}
	result, resultErr := repository.SelectEmailCustomers(context.Background(), string(customerRequest.Email))
	helper.PanicIfError(resultErr)

	fmt.Println(result["password"].(string))
}

func TestGetAndVerifyPassword(t *testing.T) {
	customerRequest := request.LoginRequest{
		Email:    "test11@mail.com",
		Password: "tes123",
	}
	result, resultErr := repository.SelectEmailCustomers(context.Background(), string(customerRequest.Email))
	helper.PanicIfError(resultErr)

	if err := helper.ComparePasswords([]byte(result["password"].(string)), []byte(customerRequest.Password)); err != nil {
		log.Fatal(err)
		return
	}

	log.Println(result)
}

func TestEmailExist(t *testing.T) {
	customerRequest := request.LoginRequest{
		Email:    "maulanazn20@mail.com",
		Password: "maulanazn123",
	}

	var result map[string]interface{}

	config.GetConnection().WithContext(context.Background()).Table("customers").Take(&result).Where("email = @email", sql.Named("email", customerRequest.Email)).Scan(&result)

	if customerRequest.Email != result["email"] {
		fmt.Println(errors.New("lkjasdf"))
		return
	}

	fmt.Println(nil)
}

func TestGetPasswordPlain(t *testing.T) {
	var data map[string]interface{}
	config.GetConnection().Table("customers").Take(&data).Select("*").Where("email = @email", sql.Named("email", "test11@mail.com"))

	fmt.Println(data["user_name"])
}

func TestUpdateAndGetCustomer(t *testing.T) {
	config.GetConnection().WithContext(context.Background()).Table("customers").Where("id = @id", sql.Named("id", "1")).Updates(map[string]interface{}{
		"userimage":  "https://instagram/akun123/image.png",
		"updated_at": time.Now(),
	})

	var data map[string]interface{}
	config.GetConnection().Table("customers").Take(&data).Where("email = @email", sql.Named("email", "maulanazn19@mail.com")).Scan(&data)

	fmt.Println(data["username"])
}

func TestUpdateCustomer(t *testing.T) {
	result, resultErr := repository.SelectCustomerById(context.Background(), "0151cf30fbd5456aa30a3e5af3ccba18")
	helper.PanicIfError(resultErr)

	if result != nil {
		customer := &entity.Customer{
			Userimage:   result["userimage"].(string),
			Username:    result["username"].(string),
			Phone:       result["phone"].(string),
			Gender:      result["gender"].(string),
			Dateofbirth: result["dateofbirth"].(string),
		}
		repository.UpdateCustomer(context.Background(), *customer, "0151cf30fbd5456aa30a3e5af3ccba18")

		log.Println(result["userimage"])

		return
	}

	customer := &entity.Customer{
		Userimage:   "https://image.com",
		Username:    "fatih",
		Phone:       "2932992",
		Gender:      "male",
		Dateofbirth: "19-10-2004",
	}

	repository.UpdateCustomer(context.Background(), *customer, "0151cf30fbd5456aa30a3e5af3ccba18")

	var data map[string]interface{}
	config.GetConnection().Table("customers").Take(&data).Where("email = @email", sql.Named("email", "maulanazn19@mail.com")).Scan(&data)

	log.Println(result["userimage"])
}

func TestGetAddressByUser(t *testing.T) {
	var result []entity.Address

	result, err := repository.AddressByUser(context.Background(), "0151cf30fbd5456aa30a3e5af3ccba18")
	helper.PanicIfError(err)

	for _, data := range result {
		jsonresult, err := json.MarshalIndent(&data, "", "")
		helper.PanicIfError(err)
		fmt.Println(string(jsonresult))
	}
}

func TestGetAddressById(t *testing.T) {
	var result entity.Address

	result, resulterr := repository.AddressById(context.Background(), "51ac602e02534e6a813b96c509b9b429")
	helper.PanicIfError(resulterr)

	jsonresult, err := json.MarshalIndent(&result, "", "")
	helper.PanicIfError(err)
	fmt.Println(string(jsonresult))
}

func TestUpdateAddress(t *testing.T) {
	var result entity.Address

	repository.UpdateAddress(context.Background(), result, "51ac602e02534e6a813b96c509b9b429")

	fmt.Println(result)
}
