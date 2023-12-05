package users

import (
	"fmt"
	"net/http"
	"util"
)

func RegisterCustomer(writer http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		CreateCustomer(req.Context(), writer, req)
		return
	}

	if _, err := fmt.Fprint(writer, "Only post is allowed"); err != nil {
		util.Log2File(err.Error())
		return
	}
}

func LoginCustomer(writer http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		VerifyCustomer(req.Context(), writer, req)
		return
	}

	if _, err := fmt.Fprint(writer, "Only post is allowed"); err != nil {
		util.Log2File(err.Error())
		return
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
	default:
		if _, err := fmt.Fprint(writer, "The requests that you want is not available"); err != nil {
			util.Log2File(err.Error())
			return
		}
		return
	}
}
