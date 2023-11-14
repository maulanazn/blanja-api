package controller

import (
	"fmt"
	"net/http"
	"service"
)

func RegisterCustomer(writer http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		service.CreateCustomer(req.Context(), writer, req)
		return
	}

	fmt.Fprint(writer, "Get is not available")
}

func LoginCustomer(writer http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		service.VerifyCustomer(req.Context(), writer, req)
		return
	}

	fmt.Fprint(writer, "Get is not available")
}

func EditCustomer(writer http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPut:
		service.EditCustomer(req.Context(), writer, req)
		return
	case http.MethodGet:
		service.ProfileCustomer(req.Context(), writer, req)
		return
	}
}
