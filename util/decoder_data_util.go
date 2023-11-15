package util

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

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

func ConvertStrInt64(data interface{}, base int, bitSize int) (int64, error) {
	format, formatErr := strconv.ParseInt(data.(string), base, bitSize)
	if formatErr != nil {
		return 0, formatErr
	}

	return format, nil
}
