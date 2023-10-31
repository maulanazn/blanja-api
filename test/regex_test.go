package test

import (
	"fmt"
	"regexp"
	"testing"
	"userboilerplate-api/webserver/helper"
)

type RegisterRequest struct {
	Username string
	Email    string
	Password string
}

func TestMatches(t *testing.T) {
	var email RegisterRequest = RegisterRequest{
		Email: "maulanazn@mail.my",
	}

	validateEmail, err := regexp.Match("[a-zA-Z_0-9\\s][@]", []byte(email.Email))
	helper.FatalIfError(err)

	if validateEmail {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
