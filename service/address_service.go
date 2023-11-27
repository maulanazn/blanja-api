package service

import (
	"context"
	"encoding/json"
	"entity"
	"fmt"
	"log"
	"net/http"
	"repository"
	"request"
	"response"
	"util"
)

func AddAddress(ctx context.Context, writer http.ResponseWriter, req *http.Request) {
	addressRequest := request.AddressCustomerRequest{}
	if err := util.DecodeRequestAndValidate(writer, req, &addressRequest); err != nil {
		util.PanicIfError(err)
		return
	}

	resultUserCookie, err := req.Cookie("USR_ID")
	util.PanicIfError(err)

	address := entity.Address{
		UserId:         resultUserCookie.Value,
		AddressType:    addressRequest.AddressType,
		RecipientName:  addressRequest.RecipientName,
		RecipientPhone: addressRequest.RecipientPhone,
		AddressName:    addressRequest.AddressName,
		PostalCode:     addressRequest.PostalCode,
		City:           addressRequest.City,
	}

	if err := repository.CreateAddress(ctx, &address); err != nil {
		writer.WriteHeader(403)
		failedResponse := response.ToWebResponse(403, "Duplicate or something, please repeat process")
		if _, err := fmt.Fprint(writer, failedResponse); err != nil {
			log.Println(err.Error())
		}

		return
	}

	writer.WriteHeader(201)
	registerResponse := response.ToWebResponse(201, "Successfully create addresss")
	if _, err := fmt.Fprint(writer, registerResponse); err != nil {
		log.Println(err.Error())
	}
}

func EditAddress(ctx context.Context, writer http.ResponseWriter, req *http.Request) {
	addressRequest := request.AddressCustomerRequest{}
	if err := util.DecodeRequestAndValidate(writer, req, &addressRequest); err != nil {
		util.PanicIfError(err)
		return
	}

	id := req.URL.Query()
	resultUserCookie, err := req.Cookie("USR_ID")
	util.PanicIfError(err)

	address := &entity.Address{
		UserId:         resultUserCookie.Value,
		AddressType:    addressRequest.AddressType,
		RecipientName:  addressRequest.RecipientName,
		RecipientPhone: addressRequest.RecipientPhone,
		AddressName:    addressRequest.AddressName,
		PostalCode:     addressRequest.PostalCode,
		City:           addressRequest.City,
	}

	if err := repository.UpdateAddress(ctx, *address, id.Get("id")); err != nil {
		writer.WriteHeader(403)
		failedResponse := response.ToWebResponse(403, "Duplicate or something, please repeat process")
		if _, err := fmt.Fprint(writer, failedResponse); err != nil {
			log.Println(err.Error())
		}

		return
	}

	writer.WriteHeader(200)
	registerResponse := response.ToWebResponse(200, "Successfully updating addresss")
	if _, err := fmt.Fprint(writer, registerResponse); err != nil {
		log.Println(err.Error())
	}
}

func AddressDetail(ctx context.Context, writer http.ResponseWriter, request *http.Request) {
	var resultUser []entity.Address
	var resultUserErr error
	id := request.URL.Query()
	result, resultErr := repository.AddressById(ctx, id.Get("id"))
	util.PanicIfError(resultErr)

	userId, idError := request.Cookie("USR_ID")
	util.PanicIfError(idError)
	resultUser, resultUserErr = repository.AddressByUser(ctx, userId.Value)
	util.PanicIfError(resultUserErr)

	if id.Has("id") {
		writer.WriteHeader(200)
		profileResponse := response.ToDetailAddressById(200, "Successfully get customer address detail", response.DetailAddressById{
			Status:  200,
			Message: "Successfully get detail address",
			Data: response.DetailAddressData{
				CustomerId:     result.UserId,
				AddressType:    result.AddressType,
				RecipientName:  result.RecipientName,
				RecipientPhone: result.RecipientPhone,
				AddressName:    result.AddressName,
				PostalCode:     result.PostalCode,
				City:           result.City,
			},
		})
		if _, err := fmt.Fprint(writer, profileResponse); err != nil {
			log.Println(err.Error())
		}

		return
	}

	for _, data := range resultUser {
		_, err := json.MarshalIndent(&data, "", "")
		util.PanicIfError(err)
	}

	writer.WriteHeader(200)
	profileResponse := response.ToDetailAddress(200, "Successfully get customer profile", response.DetailAddress{
		Status:  200,
		Message: "Successfully get detail address",
		Data:    resultUser,
	})
	if _, err := fmt.Fprint(writer, profileResponse); err != nil {
		log.Println(err.Error())
	}
}
