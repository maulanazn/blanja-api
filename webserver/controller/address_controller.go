package controller

import (
	"belanjabackend/webserver/helper"
	"belanjabackend/webserver/request"
	"belanjabackend/webserver/service"
	"encoding/json"
	"fmt"
	"net/http"
)

func AddOrEditAddress(writer http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		decoder := json.NewDecoder(req.Body)
		addressRequest := request.AddressCustomerRequest{}
		err := decoder.Decode(&addressRequest)
		helper.PanicIfError(err)

		service.AddAddress(req.Context(), addressRequest, writer, req)
		return
	case http.MethodPut:
		decoder := json.NewDecoder(req.Body)
		var addressRequest request.AddressCustomerRequest
		err := decoder.Decode(&addressRequest)
		helper.PanicIfError(err)

		service.EditAddress(req.Context(), addressRequest, writer, req)
		return
	case http.MethodGet:
		service.AddressDetail(req.Context(), writer, req)
		return
	}

	fmt.Fprint(writer, "Get is not available")
}
