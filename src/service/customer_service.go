package service

import (
	"context"
	"entity"
	"fmt"
	"net/http"
	"os"
	"repository"
	"request"
	"response"
	"util"

	"golang.org/x/crypto/bcrypt"
)

func CreateCustomer(ctx context.Context, writer http.ResponseWriter, req *http.Request) {
	registerreq := request.RegisterRequest{}
	if err := util.DecodeRequestAndValidate(writer, req, &registerreq); err != nil {
		util.PanicIfError(err)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerreq.Password), 10)
	util.PanicIfError(err)

	users := entity.Users{
		Username: registerreq.Username,
		Email:    registerreq.Email,
		Password: string(hashedPassword),
		Roles:    registerreq.Roles,
	}

	if err := repository.CreateCustomer(ctx, &users); err != nil {
		writer.WriteHeader(403)
		failedResponse := response.ToWebResponse(403, "Duplicate or something, please repeat process", writer)
		fmt.Fprint(writer, failedResponse)

		return
	}

	res := response.Data{
		Username: users.Username,
		Email:    users.Email,
		Roles:    users.Roles,
	}

	writer.WriteHeader(201)
	registerResponse := response.ToResponseData(201, "Successfully Registered", res)
	fmt.Fprint(writer, registerResponse)
}

func VerifyCustomer(ctx context.Context, writer http.ResponseWriter, req *http.Request) {
	loginreq := request.LoginRequest{}
	if err := util.DecodeRequestAndValidate(writer, req, &loginreq); err != nil {
		util.PanicIfError(err)
		return
	}

	users := entity.Users{
		Email:    loginreq.Email,
		Password: loginreq.Password,
	}

	result, resultErr := repository.SelectEmailCustomers(ctx, loginreq.Email)
	util.PanicIfError(resultErr)
	if loginreq.Email != result.Email {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "Wrong Email")
		return
	}

	if err := util.ComparePasswords([]byte(result.Password), []byte(loginreq.Password)); err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(writer, "Wrong Password")
		return
	}

	token := util.GenerateToken(users, os.Getenv("JWT_KEY"))
	res := response.Token{
		Token: token,
	}

	userid := http.Cookie{
		Name:     "USR_ID",
		Value:    result.Id,
		Path:     "/",
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}
	http.SetCookie(writer, &userid)
	writer.Header().Set("Authorization", "Bearer "+res.Token)

	loginResponse := response.ToWebResponse(200, "Successfully Login", writer)
	fmt.Fprint(writer, loginResponse)
}

func EditCustomer(ctx context.Context, writer http.ResponseWriter, req *http.Request) {
	editCustomerRequest := request.EditCustomerRequest{}
	if err := util.DecodeFormRequestAndValidate(writer, req, editCustomerRequest); err != nil {
		writer.WriteHeader(500)
		writer.Write([]byte("Failed to decode"))
		return
	}

	userimage, _, err := req.FormFile("userimage")
	util.PanicIfError(err)
	formatphone, formatphoneerr := util.ConvertStrInt64(editCustomerRequest.Phone, 10, 64)
	util.PanicIfError(formatphoneerr)
	responseimage, err := util.UploadCloudinary(userimage)
	util.PanicIfError(err)
	id, cookieerr := req.Cookie("USR_ID")
	util.PanicIfError(cookieerr)

	users := &entity.Users{
		Userimage:   responseimage.SecureURL,
		Username:    string(editCustomerRequest.Username.(string)),
		Roles:       string(editCustomerRequest.Roles.(string)),
		Phone:       formatphone,
		Gender:      string(editCustomerRequest.Gender.(string)),
		Dateofbirth: string(editCustomerRequest.Dateofbirth.(string)),
	}

	if err := repository.UpdateCustomer(ctx, *users, id.Value); err != nil {
		writer.WriteHeader(403)
		failedResponse := response.ToWebResponse(403, "Duplicate or something, please repeat process", writer)
		fmt.Fprint(writer, failedResponse)

		return
	}

	customerupdate := response.ToWebResponse(200, "Successfully updated profile", writer)
	fmt.Fprint(writer, customerupdate)
}

func ProfileCustomer(ctx context.Context, writer http.ResponseWriter, request *http.Request) {
	var roles string
	id, cookieerr := request.Cookie("USR_ID")
	util.PanicIfError(cookieerr)

	result, resulterr := repository.SelectCustomerById(ctx, id.Value)
	util.PanicIfError(resulterr)

	if result.Roles == "superuser" {
		roles = "superuser"
	} else {
		roles = "notsuper"
	}

	writer.WriteHeader(200)
	profileresp := response.ToProfileCustomer(200, "Successfully get "+roles+" profile", response.ProfileCustomer{
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
