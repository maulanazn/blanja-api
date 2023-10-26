package helper

import (
	"belanjabackend/webserver/response"
	"encoding/json"
)

func ToWebResponse(status int, message interface{}) interface{} {
	value, err := json.MarshalIndent(&response.WebResponse{
		Status:  status,
		Message: message.(string),
	}, "", "")
	FatalIfError(err)
	return string(value)
}

func ToResponseData(status int, message string, data response.Data) string {
	value, err := json.MarshalIndent(&response.RegisterResponse{
		Status:  status,
		Message: string(message),
		Data:    data,
	}, "", "")
	PanicIfError(err)
	return string(value)
}

func ToResponseToken(status int, message string, token response.Token) string {
	value, err := json.MarshalIndent(&response.LoginResponse{
		Status:  status,
		Message: string(message),
		Data:    token,
	}, "", "")
	PanicIfError(err)
	return string(value)
}
