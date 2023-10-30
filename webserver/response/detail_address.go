package response

import "belanjabackend/entity"

type DetailAddressData struct {
	Id             string
	CustomerId     string
	AddressType    string
	RecipientName  string
	RecipientPhone string
	AddressName    string
	PostalCode     string
	City           string
}

type DetailAddress struct {
	Status  int
	Message string
	Data    []entity.Address
}

type DetailAddressById struct {
	Status  int
	Message string
	Data    DetailAddressData
}
