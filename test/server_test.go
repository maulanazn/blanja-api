package test

import (
	"belanjabackend/config"
	"belanjabackend/repository"
	"belanjabackend/webserver/controller"
	"belanjabackend/webserver/helper"
	"belanjabackend/webserver/request"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootFailed(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	controller.RootHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	formatted := string(body)

	assert.NotEqual(t, formatted, "Hello world")
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
