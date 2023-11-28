package users

import (
	"fmt"
	"log"
	"net/http"
)

func RegisterCustomer(writer http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		CreateCustomer(req.Context(), writer, req)
		return
	}

	if _, err := fmt.Fprint(writer, "Get is not available"); err != nil {
		log.Println(err.Error())
	}
}

func LoginCustomer(writer http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		VerifyCustomer(req.Context(), writer, req)
		return
	}

	if _, err := fmt.Fprint(writer, "Get is not available"); err != nil {
		log.Println(err.Error())
	}
}

func PutCustomer(writer http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPut:
		EditCustomer(req.Context(), writer, req)
		return
	case http.MethodGet:
		ProfileCustomer(req.Context(), writer, req)
		return
	}
}
