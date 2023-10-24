package controller

import (
	"encoding/json"
	"net/http"
)

func RegisterCustomer(writer http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		decoder := json.NewDecoder(req.Body)
		decoder.Decode(writer)

		return
	}
}
