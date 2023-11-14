package util

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"request"
	"strconv"

	"github.com/go-playground/form"
	"github.com/go-playground/validator"
)

const (
	Passphrase = "abcdefghijklmnopqrstuvwx"
)

func DecodeRequestAndValidate(writer http.ResponseWriter, req *http.Request, data interface{}) error {
	validate := validator.New()

	if err := json.NewDecoder(req.Body).Decode(data); err != nil {
		return errors.New("failed to decode")
	}

	if err := validate.Struct(data); err != nil {
		writer.WriteHeader(400)
		writer.Write([]byte(err.Error()))
		return errors.New(err.Error())
	}

	return nil
}

func DecodeFormRequestAndValidate(writer http.ResponseWriter, req *http.Request, data interface{}) error {
	validate := validator.New()

	var decoder *form.Decoder = form.NewDecoder()

	values := parseForm()

	err := decoder.Decode(&data, values)
	if err != nil {
		log.Panic(err)
	}

	if err := validate.Struct(data); err != nil {
		writer.WriteHeader(400)
		writer.Write([]byte(err.Error()))
		return errors.New(err.Error())
	}

	return nil
}

func parseForm() url.Values {
	editCustomerRequest := request.EditCustomerRequest{}
	return url.Values{
		// FIXME: FIXING ERROR
		"user_name":   []string(editCustomerRequest.Username.([]string)),
		"phone":       []string(editCustomerRequest.Phone.([]string)),
		"gender":      []string(editCustomerRequest.Gender.([]string)),
		"dateofbirth": []string(editCustomerRequest.Dateofbirth.([]string)),
		"roles":       []string(editCustomerRequest.Roles.([]string)),
	}
}

func ConvertStrInt64(data interface{}, base int, bitSize int) (int64, error) {
	format, formatErr := strconv.ParseInt(data.(string), base, bitSize)
	if formatErr != nil {
		return 0, formatErr
	}

	return format, nil
}
