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
		CustomerId:     resultUserCookie.Value,
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
	resultaddressid, err := repository.AddressById(ctx, id.Get("id"))
	helper.PanicIfError(err)

	if req.AddressType == "" || req.RecipientName == "" || req.RecipientPhone == 0 || req.AddressName == "" || req.PostalCode == "" || req.City == "" {
		address := entity.Address{
			Id:             resultaddressid.Id,
			CustomerId:     resultaddressid.CustomerId,
			AddressType:    resultaddressid.AddressType,
			RecipientName:  resultaddressid.RecipientName,
			RecipientPhone: resultaddressid.RecipientPhone,
			AddressName:    resultaddressid.AddressName,
			PostalCode:     resultaddressid.PostalCode,
			City:           resultaddressid.City,
		}

		address.ValidateUpdate(id.Get("id"))

		if err := repository.UpdateAddress(ctx, address, id.Get("id")); err != nil {
			writer.WriteHeader(403)
			failedResponse := helper.ToWebResponse(403, "Duplicate or something, please repeat process")
			fmt.Fprint(writer, failedResponse)

			return
		}
		writer.WriteHeader(200)
		registerResponse := helper.ToWebResponse(200, "Successfully updating addresss")
		fmt.Fprint(writer, registerResponse)

		return
	} else {
		address := &entity.Address{
			Id:             helper.GenUUID(),
			CustomerId:     resultUserCookie.Value,
			AddressType:    req.AddressType,
			RecipientName:  req.RecipientName,
			RecipientPhone: req.RecipientPhone,
			AddressName:    req.AddressName,
			PostalCode:     req.PostalCode,
			City:           req.City,
		}

		address.ValidateUpdate(resultUserCookie.Value)

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

}

func AddressDetail(ctx context.Context, writer http.ResponseWriter, request *http.Request) {
	id := request.URL.Query()
	result, resulterr := repository.AddressById(ctx, id.Get("id"))
	helper.PanicIfError(resulterr)

	usrid, iderror := request.Cookie("USR_ID")
	helper.PanicIfError(iderror)
	resultuser, resuluserterr := repository.AddressByUser(ctx, usrid.Value)
	helper.PanicIfError(resuluserterr)

	if id.Has("id") {
		writer.WriteHeader(200)
		profileresp := helper.ToDetailAddress(200, "Successfully get customer address detail", response.DetailAddress{
			Status:  200,
			Message: "Successfully get detail address",
			Data: response.DetailAddressData{
				Id:             result.Id,
				CustomerId:     result.CustomerId,
				AddressType:    result.AddressType,
				RecipientName:  result.RecipientName,
				RecipientPhone: int64(result.RecipientPhone),
				AddressName:    result.AddressName,
				PostalCode:     result.PostalCode,
				City:           result.City,
			},
		})
		fmt.Fprint(writer, profileresp)

		return
	}

	for _, data := range resultuser {
		writer.WriteHeader(200)
		profileresp := helper.ToDetailAddress(200, "Successfully get customer profile", response.DetailAddress{
			Status:  200,
			Message: "Successfully get detail address",
			Data: response.DetailAddressData{
				Id:             data.Id,
				CustomerId:     data.CustomerId,
				AddressType:    data.AddressType,
				RecipientName:  data.RecipientName,
				RecipientPhone: int64(data.RecipientPhone),
				AddressName:    data.AddressName,
				PostalCode:     data.PostalCode,
				City:           data.City,
			},
		})
		fmt.Fprint(writer, profileresp)
	}

}
