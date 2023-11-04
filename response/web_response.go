package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type WebResponse struct {
	Status  int
	Message string
}

func ToWebResponse(status int, message interface{}, writer http.ResponseWriter) interface{} {
	writer.WriteHeader(status)
	value, err := json.MarshalIndent(&WebResponse{
		Status:  status,
		Message: message.(string),
	}, "", "")
	if err != nil {
		log.Fatal(err)
	}
	return string(value)
}
