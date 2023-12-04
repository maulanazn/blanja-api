package address

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"util"
)

func AddAddress(ctx context.Context, writer http.ResponseWriter, req *http.Request) {
	addressRequest := AddressCustomerRequest{}
	util.DecodeRequestAndValidate(writer, req, &addressRequest)

	userid := util.DecodeToken(req.Header.Get("Authorization"))

	address := Address{
		UserId:         userid,
		AddressType:    addressRequest.AddressType,
		RecipientName:  addressRequest.RecipientName,
		RecipientPhone: addressRequest.RecipientPhone,
		AddressName:    addressRequest.AddressName,
		PostalCode:     addressRequest.PostalCode,
		City:           addressRequest.City,
	}

	if err := CreateAddress(ctx, &address); err != nil {
		writer.WriteHeader(403)
		failedResponse := util.ToWebResponse(403, "Duplicate or something, please repeat process")
		if _, err := fmt.Fprint(writer, failedResponse); err != nil {
			log.Println(err.Error())
		}

		return
	}

	writer.WriteHeader(201)
	registerResponse := util.ToWebResponse(201, "Successfully create addresss")
	if _, err := fmt.Fprint(writer, registerResponse); err != nil {
		log.Println(err.Error())
	}
}

func EditAddress(ctx context.Context, writer http.ResponseWriter, req *http.Request) {
	addressRequest := AddressCustomerRequest{}
	util.DecodeRequestAndValidate(writer, req, &addressRequest)

	id := req.URL.Query()
	userid := util.DecodeToken(req.Header.Get("Authorization"))

	address := &Address{
		UserId:         userid,
		AddressType:    addressRequest.AddressType,
		RecipientName:  addressRequest.RecipientName,
		RecipientPhone: addressRequest.RecipientPhone,
		AddressName:    addressRequest.AddressName,
		PostalCode:     addressRequest.PostalCode,
		City:           addressRequest.City,
	}

	if err := UpdateAddress(ctx, *address, id.Get("id")); err != nil {
		writer.WriteHeader(403)
		failedResponse := util.ToWebResponse(403, "Duplicate or something, please repeat process")
		if _, err := fmt.Fprint(writer, failedResponse); err != nil {
			log.Println(err.Error())
		}

		return
	}

	writer.WriteHeader(200)
	registerResponse := util.ToWebResponse(200, "Successfully updating addresss")
	if _, err := fmt.Fprint(writer, registerResponse); err != nil {
		log.Println(err.Error())
	}
}

func AddressDetail(ctx context.Context, writer http.ResponseWriter, request *http.Request) {
	var resultUser []Address
	var resultUserErr error
	id := request.URL.Query()
	result, resultErr := AddressById(ctx, id.Get("id"))
	util.PanicIfError(resultErr)

	userid := util.DecodeToken(request.Header.Get("Authorization"))
	resultUser, resultUserErr = AddressByUser(ctx, userid)
	util.PanicIfError(resultUserErr)

	if id.Has("id") {
		writer.WriteHeader(200)
		profileResponse := ToDetailAddressById(200, "Successfully get customer address detail", DetailAddressById{
			Status:  200,
			Message: "Successfully get detail address",
			Data: DetailAddressData{
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
	profileResponse := ToDetailAddress(200, "Successfully get customer profile", DetailAddress{
		Status:  200,
		Message: "Successfully get detail address",
		Data:    resultUser,
	})
	if _, err := fmt.Fprint(writer, profileResponse); err != nil {
		log.Println(err.Error())
	}
}
