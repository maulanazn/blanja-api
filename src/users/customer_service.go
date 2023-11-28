package users

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"util"

	"golang.org/x/crypto/bcrypt"
)

func CreateCustomer(ctx context.Context, writer http.ResponseWriter, req *http.Request) {
	registerReq := RegisterRequest{}
	if err := util.DecodeRequestAndValidate(writer, req, &registerReq); err != nil {
		util.PanicIfError(err)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerReq.Password), 10)
	util.PanicIfError(err)

	users := Users{
		Username: registerReq.Username,
		Email:    registerReq.Email,
		Password: string(hashedPassword),
		Roles:    registerReq.Roles,
	}

	if err := InsertCustomer(ctx, &users); err != nil {
		writer.WriteHeader(403)
		failedResponse := util.ToWebResponse(403, "Duplicate or something, please repeat process")
		if _, err := fmt.Fprint(writer, failedResponse); err != nil {
			log.Println(err.Error())
		}

		return
	}

	res := Data{
		Username: users.Username,
		Email:    users.Email,
		Roles:    users.Roles,
	}

	writer.WriteHeader(201)
	registerResponse := ToResponseData(201, "Successfully Registered", res)
	if _, err := fmt.Fprint(writer, registerResponse); err != nil {
		log.Println(err.Error())
	}
}

func VerifyCustomer(ctx context.Context, writer http.ResponseWriter, req *http.Request) {
	loginReq := LoginRequest{}
	if err := util.DecodeRequestAndValidate(writer, req, &loginReq); err != nil {
		util.PanicIfError(err)
		return
	}

	result, resultErr := SelectEmailCustomers(ctx, loginReq.Email)
	util.PanicIfError(resultErr)

	users := Users{
		Id:    result.Id,
		Email: loginReq.Email,
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

	writer.Header().Set("Authorization", "Bearer "+token)

	loginResponse := util.ToWebResponse(200, "Successfully Login")
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
		failedResponse := util.ToWebResponse(400, err.Error())
		if _, err := fmt.Fprint(writer, failedResponse); err != nil {
			log.Println(err.Error())
		}
		return
	}

	responseImage, err := util.UploadCloudinary(userImage)
	util.BadStatusIfError(err, writer)
	userid := util.DecodeToken(req.Header.Get("Authorization"))

	users := &Users{
		UserImage:   responseImage.SecureURL,
		Username:    req.FormValue("username"),
		Roles:       req.FormValue("roles"),
		Phone:       util.ConvertStrInt64(req.FormValue("phone"), 10, 64),
		Gender:      req.FormValue("gender"),
		DateOfBirth: req.FormValue("dateofbirth"),
	}

	if err := UpdateCustomer(ctx, *users, userid); err != nil {
		writer.WriteHeader(403)
		failedResponse := util.ToWebResponse(403, "Duplicate or something, please repeat process")
		if _, err := fmt.Fprint(writer, failedResponse); err != nil {
			log.Println(err.Error())
		}

		return
	}

	customerUpdate := util.ToWebResponse(200, "Successfully updated profile")
	if _, err := fmt.Fprint(writer, customerUpdate); err != nil {
		log.Println(err.Error())
	}
}

func ProfileCustomer(ctx context.Context, writer http.ResponseWriter, request *http.Request) {
	var roles string
	userid := util.DecodeToken(request.Header.Get("Authorization"))

	result, resultErr := SelectCustomerById(ctx, userid)
	util.PanicIfError(resultErr)

	if result.Roles == "superuser" {
		roles = "superuser"
	} else {
		roles = "notsuper"
	}

	writer.WriteHeader(200)
	profileResponse := ToProfileCustomer(200, "Successfully get "+roles+" profile", ProfileCustomerStruct{
		Data: ProfileCustomerData{
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
