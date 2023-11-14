package request

import (
	"regexp"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type EditCustomerRequest struct {
	Username    interface{}
	Phone       interface{}
	Gender      interface{}
	Dateofbirth interface{}
	Roles       interface{}
}

func (ecq EditCustomerRequest) Validate() error {
	return validation.ValidateStruct(&ecq,
		validation.Field(&ecq.Username, validation.Match(regexp.MustCompile("^[a-zA-Z0-9@,.!]{5,50}$"))),
		validation.Field(&ecq.Phone, is.Digit),
		validation.Field(&ecq.Gender, validation.In("male", "female")),
		validation.Field(&ecq.Dateofbirth, validation.Date(time.DateOnly)),
		validation.Field(&ecq.Roles, validation.In("notsuper", "superuser")),
	)
}
