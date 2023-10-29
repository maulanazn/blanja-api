package response

type DetailAddressData struct {
	Id             string
	CustomerId     string
	AddressType    string
	RecipientName  string
	RecipientPhone int64
	AddressName    string
	PostalCode     string
	City           string
}

type DetailAddress struct {
	Status  int
	Message string
	Data    DetailAddressData
}
