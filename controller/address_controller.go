package controller

import (
	"fmt"
	"log"
	"net/http"
	"service"
)

func AddOrEditAddress(writer http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		service.AddAddress(req.Context(), writer, req)
		return
	case http.MethodPut:
		service.EditAddress(req.Context(), writer, req)
		return
	case http.MethodGet:
		service.AddressDetail(req.Context(), writer, req)
		return
	}
	
	if _, err := fmt.Fprint(writer, "Get is not available"); err != nil {
		log.Println(err.Error())
	}
}
