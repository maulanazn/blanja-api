package response

import (
	"encoding/json"
	"log"
)

type WebResponse struct {
	Status  int
	Message string
}

func ToWebResponse(status int, message interface{}) interface{} {
	value, err := json.MarshalIndent(&WebResponse{
		Status:  status,
		Message: message.(string),
	}, "", "")
	if err != nil {
		log.Fatal(err)
	}
	return string(value)
}
