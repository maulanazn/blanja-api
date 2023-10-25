package service

import (
	"belanjabackend/entity"
	"belanjabackend/repository"
	"belanjabackend/webserver/helper"
	"belanjabackend/webserver/request"
	"belanjabackend/webserver/response"
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func CreateCustomer(ctx context.Context, req request.RegisterRequest, writer http.ResponseWriter) {
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

func VerifyCustomer(ctx context.Context, req request.LoginRequest, writer http.ResponseWriter) {
	customer := entity.Customer{
		Email:    req.Email,
		Password: req.Password,
	}

	result, resultErr := repository.SelectEmailCustomers(ctx, &req.Email)
	helper.PanicIfError(resultErr)

	password := result["password"].(string)
	email := result["email"].(string)
	file, _ := os.ReadFile("../../.env")

	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(req.Password)); err != nil {
		helper.BadStatusIfError(err, writer)
		return
	}

	var key []byte = []byte(email + string(file))
	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": customer.Username,
		"email":    customer.Email,
		"exp":      time.Now().Add(time.Duration(time.Now().UTC().Day())),
	})
	token, err := generateToken.SignedString(key)
	helper.PanicIfError(err)

	response := response.Token{
		Value: token,
	}

	writer.WriteHeader(200)
	loginResponse := helper.ToResponseToken(200, "Successfully Login", response)
	fmt.Fprint(writer, loginResponse)
}
