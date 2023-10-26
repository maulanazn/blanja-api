package service

import (
	"belanjabackend/config"
	"belanjabackend/entity"
	"belanjabackend/repository"
	"belanjabackend/webserver/helper"
	"belanjabackend/webserver/request"
	"belanjabackend/webserver/response"
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func CreateCustomer(ctx context.Context, req request.RegisterRequest, writer http.ResponseWriter) {
	if err := req.Validate(); err != nil {
		writer.WriteHeader(400)
		fmt.Fprint(writer, err)

		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	helper.PanicIfError(err)

	customer := entity.Customer{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	if err := repository.CreateCustomer(ctx, &customer); err != nil {
		writer.WriteHeader(403)
		failedResponse := helper.ToWebResponse(403, "Duplicate or something, please repeat process")
		fmt.Fprint(writer, failedResponse)

		return
	}

	response := response.Data{
		Username: customer.Username,
		Email:    customer.Email,
	}

	writer.WriteHeader(201)
	registerResponse := helper.ToResponseData(201, "Successfully Registered", response)
	fmt.Fprint(writer, registerResponse)
}

func VerifyCustomer(ctx context.Context, req request.LoginRequest, writer http.ResponseWriter, request *http.Request) {
	var datamap map[string]interface{}
	if err := req.Validate(); err != nil {
		writer.WriteHeader(400)
		fmt.Fprint(writer, err)

		return
	}

	customer := entity.Customer{
		Email:    req.Email,
		Password: req.Password,
	}

	result, resultErr := repository.SelectEmailCustomers(ctx, string(req.Email))
	helper.PanicIfError(resultErr)

	config.GetConnection().WithContext(context.Background()).Table("customers").Take(&datamap).Where("email = @email", sql.Named("email", req.Email)).Scan(&result)
	if req.Email != result["email"].(string) {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "Wrong Email")
		return
	}

	if err := helper.ComparePasswords([]byte(result["password"].(string)), []byte(req.Password)); err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(writer, "Wrong Password")
		return
	}

	var key []byte = []byte(result["email"].(string) + helper.ReadEnv("../../.env"))
	token := helper.GenerateToken(customer, key)

	response := response.Token{
		Token: token,
	}

	username := http.Cookie{
		Name:     "USR_ID",
		Value:    result["username"].(string),
		Path:     "/",
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}

	usertoken := http.Cookie{
		Name:     "TKN_ID",
		Value:    response.Token,
		Path:     "/",
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(writer, &username)
	http.SetCookie(writer, &usertoken)

	loginResponse := helper.ToWebResponse(200, "Successfully Login")
	fmt.Fprint(writer, loginResponse)
}

func EditCustomer(ctx context.Context, req request.EditCustomerRequest, writer http.ResponseWriter) {
	fmt.Fprint(writer, "Edit Customer")
}
