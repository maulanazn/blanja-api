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
	"regexp"

	"github.com/albrow/forms"
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
		Id:       helper.GenUUID(),
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
		Value:    string(result["id"].([]uint8)),
		Path:     "/",
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(writer, &username)

	writer.Header().Set("Authorization", "Bearer "+response.Token)

	loginResponse := helper.ToWebResponse(200, "Successfully Login")
	fmt.Fprint(writer, loginResponse)
}

func EditCustomer(ctx context.Context, writer http.ResponseWriter, request *http.Request) {
	userdata, formerr := forms.Parse(request)
	helper.PanicIfError(formerr)
	id, cookieerr := request.Cookie("USR_ID")
	helper.PanicIfError(cookieerr)

	userdata.Validator().Require("username")
	userdata.Validator().Require("gender")
	userdata.Validator().Require("dateofbirth")
	userdata.Validator().Match("phone", regexp.MustCompile("^[0-9]{3,40}$"))
	userdata.Validator().AcceptFileExts("userimage", "jpg", "jpeg", "png", "gif")
	userimage, _, err := request.FormFile("userimage")
	helper.PanicIfError(err)
	formatphone, formatphoneerr := helper.ConvertStrInt64(userdata.Get("phone"), 10, 64)
	helper.PanicIfError(formatphoneerr)

	responseimage, err := helper.UploadCloudinary(userimage)
	helper.PanicIfError(err)

	customer := &entity.Customer{
		Userimage:   responseimage.SecureURL,
		Username:    userdata.Get("username"),
		Phone:       formatphone,
		Gender:      userdata.Get("gender"),
		Dateofbirth: userdata.Get("dateofbirth"),
	}

	if err := repository.UpdateCustomer(ctx, *customer, id.Value); err != nil {
		writer.WriteHeader(403)
		failedResponse := helper.ToWebResponse(403, "Duplicate or something, please repeat process")
		fmt.Fprint(writer, failedResponse)

		return
	}

	customerupdate := helper.ToWebResponse(200, "Successfully updated profile")
	fmt.Fprint(writer, customerupdate)
}

func ProfileCustomer(ctx context.Context, writer http.ResponseWriter, request *http.Request) {
	id, cookieerr := request.Cookie("USR_ID")
	helper.PanicIfError(cookieerr)

	result, resulterr := repository.SelectCustomerById(ctx, id.Value)
	helper.PanicIfError(resulterr)

	writer.WriteHeader(200)
	profileresp := helper.ToProfileCustomer(200, "Successfully get customer profile", response.ProfileCustomer{
		Status:  200,
		Message: "Successfully get customer profile",
		Data: response.ProfileCustomerData{
			Userimage:   result.Userimage,
			Username:    result.Username,
			Phone:       result.Phone,
			Gender:      result.Gender,
			Dateofbirth: result.Dateofbirth,
		},
	})
	fmt.Fprint(writer, profileresp)
}
