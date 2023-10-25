package test

import (
	"belanjabackend/config"
	"belanjabackend/repository"
	"belanjabackend/webserver/controller"
	"belanjabackend/webserver/helper"
	"context"
	"database/sql"
	"fmt"
	"io"
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
	result, resultErr := repository.SelectEmailCustomers(context.Background(), "test11@mail.com")
	helper.PanicIfError(resultErr)

	fmt.Println(result["password"])
}

func TestGetPasswordPlain(t *testing.T) {
	var data map[string]interface{}
	config.GetConnection().Table("customers").Take(&data).Select("*").Where("email = @email", sql.Named("email", "test11@mail.com"))

	fmt.Println(data)
}
