package test

import (
	"belanjabackend/config"
	"belanjabackend/entity"
	"belanjabackend/repository"
	"belanjabackend/webserver/helper"
	"belanjabackend/webserver/request"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"testing"
	"time"
)

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
	result, resultErr := repository.SelectCustomerById(context.Background(), 13)
	helper.PanicIfError(resultErr)

	if result != nil {
		customer := &entity.Customer{
			Userimage:   result["userimage"].(string),
			Username:    result["username"].(string),
			Phone:       result["phone"].(int64),
			Gender:      result["gender"].(string),
			Dateofbirth: result["dateofbirth"].(string),
		}
		repository.UpdateCustomer(context.Background(), *customer, 11)

		log.Println(result["userimage"])

		return
	}

	customer := &entity.Customer{
		Userimage:   "https://image.com",
		Username:    "fatih",
		Phone:       2932992,
		Gender:      "male",
		Dateofbirth: "19-10-2004",
	}

	repository.UpdateCustomer(context.Background(), *customer, 11)

	var data map[string]interface{}
	config.GetConnection().Table("customers").Take(&data).Where("email = @email", sql.Named("email", "maulanazn19@mail.com")).Scan(&data)

	log.Println(result["userimage"])
}
