package controller

import (
	"fmt"
	"helper"
	"net/http"
	"request"
	"service"
)

func RegisterCustomer(writer http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		req.Header.Add("Content-Type", "application/json")
		registerRequest := request.RegisterRequest{}
		err := helper.DecodeRequest(req, &registerRequest)
		helper.PanicIfError(err)

		service.CreateCustomer(req.Context(), registerRequest, writer)
		return
	}

	fmt.Fprint(writer, "Get is not available")
}

func LoginCustomer(writer http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		req.Header.Add("Content-Type", "application/json")
		loginRequest := request.LoginRequest{}
		err := helper.DecodeRequest(req, &loginRequest)
		helper.PanicIfError(err)

		service.VerifyCustomer(req.Context(), loginRequest, writer, req)
		return
	}

	fmt.Fprint(writer, "Get is not available")
}

func EditCustomer(writer http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPut:
		req.Header.Add("Content-Type", "multipart/form-data")
		service.EditCustomer(req.Context(), writer, req)
		return
	case http.MethodGet:
		service.ProfileCustomer(req.Context(), writer, req)
		return
	}
}
