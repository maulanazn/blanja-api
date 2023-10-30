package request

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type AddressCustomerRequest struct {
	CustomerId     string
	AddressType    string
	RecipientName  string
	RecipientPhone int
	AddressName    string
	PostalCode     string
	City           string
}

func (ecq AddressCustomerRequest) Validate() error {
	return validation.ValidateStruct(&ecq,
		validation.Field(&ecq.AddressType, validation.In("home", "office")),
		validation.Field(&ecq.RecipientName, is.Alpha),
		validation.Field(&ecq.RecipientPhone, is.Digit),
		validation.Field(&ecq.AddressName, is.Alphanumeric),
		validation.Field(&ecq.PostalCode, is.Alphanumeric),
		validation.Field(&ecq.City, is.Alphanumeric),
	)
}
