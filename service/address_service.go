package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"userboilerplate-api/entity"
	"userboilerplate-api/repository"
	"userboilerplate-api/webserver/helper"
	"userboilerplate-api/webserver/request"
	"userboilerplate-api/webserver/response"
)

func AddAddress(ctx context.Context, req request.AddressCustomerRequest, writer http.ResponseWriter, request *http.Request) {
	if err := req.Validate(); err != nil {
		writer.WriteHeader(400)
		fmt.Fprint(writer, err)

		return
	}
	resultUserCookie, err := request.Cookie("USR_ID")
	helper.PanicIfError(err)

	address := entity.Address{
		Id:             helper.GenUUID(),
		UserId:         resultUserCookie.Value,
		AddressType:    req.AddressType,
		RecipientName:  req.RecipientName,
		RecipientPhone: req.RecipientPhone,
		AddressName:    req.AddressName,
		PostalCode:     req.PostalCode,
		City:           req.City,
	}

	if err := repository.CreateAddress(ctx, &address); err != nil {
		writer.WriteHeader(403)
		failedResponse := helper.ToWebResponse(403, "Duplicate or something, please repeat process")
		fmt.Fprint(writer, failedResponse)

		return
	}

	writer.WriteHeader(201)
	registerResponse := helper.ToWebResponse(201, "Successfully create addresss")
	fmt.Fprint(writer, registerResponse)
}

func EditAddress(ctx context.Context, req request.AddressCustomerRequest, writer http.ResponseWriter, request *http.Request) {
	if err := req.Validate(); err != nil {
		writer.WriteHeader(400)
		fmt.Fprint(writer, err)

		return
	}

	id := request.URL.Query()
	resultUserCookie, err := request.Cookie("USR_ID")
	helper.PanicIfError(err)

	address := &entity.Address{
		UserId:         resultUserCookie.Value,
		AddressType:    req.AddressType,
		RecipientName:  req.RecipientName,
		RecipientPhone: req.RecipientPhone,
		AddressName:    req.AddressName,
		PostalCode:     req.PostalCode,
		City:           req.City,
	}

	if err := repository.UpdateAddress(ctx, *address, id.Get("id")); err != nil {
		writer.WriteHeader(403)
		failedResponse := helper.ToWebResponse(403, "Duplicate or something, please repeat process")
		fmt.Fprint(writer, failedResponse)

		return
	}

	writer.WriteHeader(200)
	registerResponse := helper.ToWebResponse(200, "Successfully updating addresss")
	fmt.Fprint(writer, registerResponse)
}

func AddressDetail(ctx context.Context, writer http.ResponseWriter, request *http.Request) {
	var resultuser []entity.Address
	var resultusererr error
	id := request.URL.Query()
	result, resulterr := repository.AddressById(ctx, id.Get("id"))
	helper.PanicIfError(resulterr)

	usrid, iderror := request.Cookie("USR_ID")
	helper.PanicIfError(iderror)
	resultuser, resultusererr = repository.AddressByUser(ctx, usrid.Value)
	helper.PanicIfError(resultusererr)

	if id.Has("id") {
		writer.WriteHeader(200)
		profileresp := helper.ToDetailAddressById(200, "Successfully get customer address detail", response.DetailAddressById{
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
		helper.PanicIfError(err)
	}

	writer.WriteHeader(200)
	profileresp := helper.ToDetailAddress(200, "Successfully get customer profile", response.DetailAddress{
		Status:  200,
		Message: "Successfully get detail address",
		Data:    resultuser,
	})
	fmt.Fprint(writer, profileresp)
}
