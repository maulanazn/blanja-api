package helper

import (
	"belanjabackend/webserver/response"
	"encoding/json"
)

func ToWebResponse(status int, message interface{}) interface{} {
	value, err := json.MarshalIndent(&response.WebResponse{
		Status:  status,
		Message: message.(string),
	}, "", "")
	FatalIfError(err)
	return string(value)
}

func ToResponseData(status int, message string, data response.Data) string {
	value, err := json.MarshalIndent(&response.RegisterResponse{
		Status:  status,
		Message: string(message),
		Data:    data,
	}, "", "")
	PanicIfError(err)
	return string(value)
}

func ToProfileCustomer(status int, message string, data response.ProfileCustomer) string {
	value, err := json.MarshalIndent(&response.ProfileCustomer{
		Status:  status,
		Message: string(message),
		Data: response.ProfileCustomerData{
			Userimage:   data.Data.Userimage,
			Username:    data.Data.Username,
			Phone:       data.Data.Phone,
			Gender:      data.Data.Gender,
			Dateofbirth: data.Data.Dateofbirth,
		},
	}, "", "\t")
	PanicIfError(err)
	return string(value)
}

func ToDetailAddress(status int, message string, data response.DetailAddress) string {
	value, err := json.MarshalIndent(&response.DetailAddress{
		Status:  status,
		Message: string(message),
		Data:    data.Data,
	}, "", "\t")
	PanicIfError(err)
	return string(value)
}

func ToDetailAddressById(status int, message string, data response.DetailAddressById) string {
	value, err := json.MarshalIndent(&response.DetailAddressById{
		Status:  status,
		Message: string(message),
		Data: response.DetailAddressData{
			CustomerId:     data.Data.CustomerId,
			AddressType:    data.Data.AddressType,
			RecipientName:  data.Data.RecipientName,
			RecipientPhone: data.Data.RecipientPhone,
			AddressName:    data.Data.AddressName,
			PostalCode:     data.Data.PostalCode,
			City:           data.Data.City,
		},
	}, "", "\t")
	PanicIfError(err)
	return string(value)
}

func ToResponseToken(status int, message string, token response.Token) string {
	value, err := json.MarshalIndent(&response.LoginResponse{
		Status:  status,
		Message: string(message),
		Data:    token,
	}, "", "")
	PanicIfError(err)
	return string(value)
}
