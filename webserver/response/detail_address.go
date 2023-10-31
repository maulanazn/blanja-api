package response

import "belanjabackend/entity"

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
