package service

import (
	"context"
	"entity"
	"fmt"
	"log"
	"net/http"
	"os"
	"repository"
	"request"
	"response"
	"util"

	"golang.org/x/crypto/bcrypt"
)

func CreateCustomer(ctx context.Context, writer http.ResponseWriter, req *http.Request) {
	registerReq := request.RegisterRequest{}
	if err := util.DecodeRequestAndValidate(writer, req, &registerReq); err != nil {
		util.PanicIfError(err)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerReq.Password), 10)
	util.PanicIfError(err)

	users := entity.Users{
		Username: registerReq.Username,
		Email:    registerReq.Email,
		Password: string(hashedPassword),
		Roles:    registerReq.Roles,
	}

	if err := repository.CreateCustomer(ctx, &users); err != nil {
		writer.WriteHeader(403)
		failedResponse := response.ToWebResponse(403, "Duplicate or something, please repeat process")
		if _, err := fmt.Fprint(writer, failedResponse); err != nil {
			log.Println(err.Error())
		}

		return
	}

	res := response.Data{
		Username: users.Username,
		Email:    users.Email,
		Roles:    users.Roles,
	}

	writer.WriteHeader(201)
	registerResponse := response.ToResponseData(201, "Successfully Registered", res)
	if _, err := fmt.Fprint(writer, registerResponse); err != nil {
		log.Println(err.Error())
	}
}

func VerifyCustomer(ctx context.Context, writer http.ResponseWriter, req *http.Request) {
	loginReq := request.LoginRequest{}
	if err := util.DecodeRequestAndValidate(writer, req, &loginReq); err != nil {
		util.PanicIfError(err)
		return
	}

	result, resultErr := repository.SelectEmailCustomers(ctx, loginReq.Email)
	util.PanicIfError(resultErr)

	users := entity.Users{
    Id: result.Id,
		Email:    loginReq.Email,
	}

	if loginReq.Email != result.Email {
		writer.WriteHeader(http.StatusBadRequest)
		if _, err := fmt.Fprint(writer, "Wrong Email"); err != nil {
			log.Println(err.Error())
		}
		return
	}

	if err := util.ComparePasswords([]byte(result.Password), []byte(loginReq.Password)); err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		if _, err := fmt.Fprint(writer, "Wrong Password"); err != nil {
			log.Println(err.Error())
		}
		return
	}

	token := util.GenerateToken(users.Id, os.Getenv("JWT_KEY"))
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

	loginResponse := response.ToWebResponse(200, "Successfully Login")
	if _, err := fmt.Fprint(writer, loginResponse); err != nil {
		log.Println(err.Error())
	}
}

func EditCustomer(ctx context.Context, writer http.ResponseWriter, req *http.Request) {
	userImage, userImageHeader, err := req.FormFile("userimage")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	if err := util.ValidateImage(userImage, userImageHeader, writer); err != nil {
		failedResponse := response.ToWebResponse(400, err.Error())
		if _, err := fmt.Fprint(writer, failedResponse); err != nil {
			log.Println(err.Error())
		}
		return
	}

	formatPhone, formatPhoneErr := util.ConvertStrInt64(req.FormValue("phone"), 10, 64)
	util.PanicIfError(formatPhoneErr)
	responseImage, err := util.UploadCloudinary(userImage)
	util.BadStatusIfError(err, writer)
	id, cookieErr := req.Cookie("USR_ID")
	util.PanicIfError(cookieErr)

	users := &entity.Users{
		UserImage:   responseImage.SecureURL,
		Username:    req.FormValue("username"),
		Roles:       req.FormValue("roles"),
		Phone:       formatPhone,
		Gender:      req.FormValue("gender"),
		DateOfBirth: req.FormValue("dateofbirth"),
	}

	if err := repository.UpdateCustomer(ctx, *users, id.Value); err != nil {
		writer.WriteHeader(403)
		failedResponse := response.ToWebResponse(403, "Duplicate or something, please repeat process")
		if _, err := fmt.Fprint(writer, failedResponse); err != nil {
			log.Println(err.Error())
		}

		return
	}

	customerUpdate := response.ToWebResponse(200, "Successfully updated profile")
	if _, err := fmt.Fprint(writer, customerUpdate); err != nil {
		log.Println(err.Error())
	}
}

func ProfileCustomer(ctx context.Context, writer http.ResponseWriter, request *http.Request) {
	var roles string
	id, cookieErr := request.Cookie("USR_ID")
	util.PanicIfError(cookieErr)

	result, resultErr := repository.SelectCustomerById(ctx, id.Value)
	util.PanicIfError(resultErr)

	if result.Roles == "superuser" {
		roles = "superuser"
	} else {
		roles = "notsuper"
	}

	writer.WriteHeader(200)
	profileResponse := response.ToProfileCustomer(200, "Successfully get "+roles+" profile", response.ProfileCustomer{
		Data: response.ProfileCustomerData{
			Userimage:   result.UserImage,
			Username:    result.Username,
			Phone:       result.Phone,
			Gender:      result.Gender,
			Dateofbirth: result.DateOfBirth,
		},
	})
	
	if _, err := fmt.Fprint(writer, profileResponse); err != nil {
		log.Println(err.Error())
	}
}
