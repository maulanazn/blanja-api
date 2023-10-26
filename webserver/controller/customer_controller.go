package controller

import (
	"belanjabackend/webserver/helper"
	"belanjabackend/webserver/request"
	"belanjabackend/webserver/service"
	"encoding/json"
	"fmt"
	"net/http"
)

func RegisterCustomer(writer http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		decoder := json.NewDecoder(req.Body)
		registerRequest := request.RegisterRequest{}
		err := decoder.Decode(&registerRequest)
		helper.PanicIfError(err)

		service.CreateCustomer(req.Context(), registerRequest, writer)
		return
	}

	fmt.Fprint(writer, "Get is not available")
}

func LoginCustomer(writer http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		decoder := json.NewDecoder(req.Body)
		loginRequest := request.LoginRequest{}
		err := decoder.Decode(&loginRequest)
		helper.PanicIfError(err)

		service.VerifyCustomer(req.Context(), loginRequest, writer, req)
		return
	}

	fmt.Fprint(writer, "Get is not available")
}

func EditCustomer(writer http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		decoder := json.NewDecoder(req.Body)
		editCustomerRequest := request.EditCustomerRequest{}
		err := decoder.Decode(&editCustomerRequest)
		helper.PanicIfError(err)

		service.EditCustomer(req.Context(), editCustomerRequest, writer)
		return
	}

	fmt.Fprint(writer, "Data customer by id soon")
}
