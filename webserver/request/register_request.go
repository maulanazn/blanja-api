package request

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type RegisterRequest struct {
	Username string
	Email    string
	Password string
}

func (rrq RegisterRequest) Validate() error {
	return validation.ValidateStruct(&rrq,
		validation.Field(&rrq.Username, validation.Required, validation.Match(regexp.MustCompile("^[a-z0-9]{5,50}$"))),
		validation.Field(&rrq.Email, validation.Required, is.Email),
		validation.Field(&rrq.Password, validation.Required, validation.Match(regexp.MustCompile("^[a-z0-9@,.!]{5,50}$"))),
	)
}
