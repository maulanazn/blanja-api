package request

import (
	"regexp"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type EditCustomerRequest struct {
	Userimage   string
	Username    string
	Phone       int
	Gender      string
	Dateofbirth string
	UpdatedAt   string
	DeletedAt   string
}

func (ecq EditCustomerRequest) Validate() error {
	return validation.ValidateStruct(&ecq,
		validation.Field(&ecq.Userimage, is.DataURI),
		validation.Field(&ecq.Username, validation.Match(regexp.MustCompile("^[a-zA-Z0-9@,.!]{5,50}$"))),
		validation.Field(&ecq.Phone, validation.Match(regexp.MustCompile("^[0-9]{3,50}$"))),
		validation.Field(&ecq.Gender, validation.In("male", "female")),
		validation.Field(&ecq.Dateofbirth, validation.Date(time.DateOnly)),
	)
}
