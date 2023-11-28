package address

type AddressCustomerRequest struct {
	CustomerId     string `validate:"required,uuid4"`
	AddressType    string `json:"address_type" validate:"required,alpha"`
	RecipientName  string `json:"recipient_name" validate:"required,alpha|alphanum"`
	RecipientPhone string `json:"recipient_phone" validate:"required,numeric"`
	AddressName    string `json:"address_name" validate:"required,alpha|alphanum,ascii"`
	PostalCode     string `json:"postal_code" validate:"required,numeric"`
	City           string `json:"city" validate:"required,alpha|alphanum"`
}
