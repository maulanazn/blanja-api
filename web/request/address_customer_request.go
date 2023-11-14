package request

type AddressCustomerRequest struct {
	CustomerId     string `validate:"uuid4"`
	AddressType    string `json:"address_type" validate:"alpha"`
	RecipientName  string `json:"recipient_name" validate:"alpha,alphanum"`
	RecipientPhone string `json:"recipient_phone" validate:"numeric"`
	AddressName    string `json:"address_name" validate:"alpha,alphanum"`
	PostalCode     string `json:"postal_code" validate:"numeric"`
	City           string `json:"city" validate:"alpha,alphanum"`
}
