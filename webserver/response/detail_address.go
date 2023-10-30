package response

type DetailAddressData struct {
	Id             string `json:"id"`
	CustomerId     string `json:"customer_id"`
	AddressType    string `json:"address_type"`
	RecipientName  string `json:"recipient_name"`
	RecipientPhone int64  `json:"recipient_phone"`
	AddressName    string `json:"address_name"`
	PostalCode     string `json:"postal_code"`
	City           string `json:"city"`
}

type DetailAddress struct {
	Status  int
	Message string
	Data    DetailAddressData
}
