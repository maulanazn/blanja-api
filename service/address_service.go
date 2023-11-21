package service

import (
	"context"
	"encoding/json"
	"entity"
	"fmt"
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
		fmt.Fprint(writer, failedResponse)

		return
	}

	writer.WriteHeader(201)
	registerResponse := response.ToWebResponse(201, "Successfully create addresss")
	fmt.Fprint(writer, registerResponse)
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
		fmt.Fprint(writer, failedResponse)

		return
	}

	writer.WriteHeader(200)
	registerResponse := response.ToWebResponse(200, "Successfully updating addresss")
	fmt.Fprint(writer, registerResponse)
}

func AddressDetail(ctx context.Context, writer http.ResponseWriter, request *http.Request) {
	var resultuser []entity.Address
	var resultusererr error
	id := request.URL.Query()
	result, resulterr := repository.AddressById(ctx, id.Get("id"))
	util.PanicIfError(resulterr)

	usrid, iderror := request.Cookie("USR_ID")
	util.PanicIfError(iderror)
	resultuser, resultusererr = repository.AddressByUser(ctx, usrid.Value)
	util.PanicIfError(resultusererr)

	if id.Has("id") {
		writer.WriteHeader(200)
		profileresp := response.ToDetailAddressById(200, "Successfully get customer address detail", response.DetailAddressById{
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
		fmt.Fprint(writer, profileresp)

		return
	}

	for _, data := range resultuser {
		_, err := json.MarshalIndent(&data, "", "")
		util.PanicIfError(err)
	}

	writer.WriteHeader(200)
	profileresp := response.ToDetailAddress(200, "Successfully get customer profile", response.DetailAddress{
		Status:  200,
		Message: "Successfully get detail address",
		Data:    resultuser,
	})
	fmt.Fprint(writer, profileresp)
}
