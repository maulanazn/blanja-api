package request

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type LoginRequest struct {
	Email    string
	Password string
}

func (lrq LoginRequest) Validate() error {
	return validation.ValidateStruct(&lrq,
		validation.Field(&lrq.Email, validation.Required, is.Email),
		validation.Field(&lrq.Password, validation.Required, validation.Match(regexp.MustCompile("^[a-z0-9@,.!]{5,50}$"))),
	)
}
