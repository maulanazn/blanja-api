package controller

import (
	"fmt"
	"helper"
	"net/http"
	"request"
	"service"
)

func AddOrEditAddress(writer http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		addressRequest := request.AddressCustomerRequest{}
		err := helper.DecodeRequest(req, &addressRequest)
		helper.PanicIfError(err)

		service.AddAddress(req.Context(), addressRequest, writer, req)
		return
	case http.MethodPut:
		addressRequest := request.AddressCustomerRequest{}
		err := helper.DecodeRequest(req, &addressRequest)

		helper.PanicIfError(err)

		service.EditAddress(req.Context(), addressRequest, writer, req)
		return
	case http.MethodGet:
		service.AddressDetail(req.Context(), writer, req)
		return
	}

	fmt.Fprint(writer, "Get is not available")
}
