package response

import (
	"encoding/json"
	"entity"
)

type DetailAddressData struct {
	CustomerId     string
	AddressType    string
	RecipientName  string
	RecipientPhone string
	AddressName    string
	PostalCode     string
	City           string
}

type DetailAddressById struct {
	Status  int
	Message string
	Data    DetailAddressData
}

type DetailAddress struct {
	Status  int
	Message string
	Data    []entity.Address
}

func ToDetailAddress(status int, message string, data DetailAddress) string {
	value, err := json.MarshalIndent(&DetailAddress{
		Status:  status,
		Message: string(message),
		Data:    data.Data,
	}, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(value)
}

func ToDetailAddressById(status int, message string, data DetailAddressById) string {
	value, err := json.MarshalIndent(&DetailAddressById{
		Status:  status,
		Message: string(message),
		Data: DetailAddressData{
			CustomerId:     data.Data.CustomerId,
			AddressType:    data.Data.AddressType,
			RecipientName:  data.Data.RecipientName,
			RecipientPhone: data.Data.RecipientPhone,
			AddressName:    data.Data.AddressName,
			PostalCode:     data.Data.PostalCode,
			City:           data.Data.City,
		},
	}, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(value)
}
