package controller

import (
	"context"
	"net/http"
	"service"
)

func AddorEditProduct(writer http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		service.AddProduct(context.Background(), writer, req)
		return
	case http.MethodGet:
		service.GetProduct(context.Background(), writer, req)
		return
	case http.MethodPut:
		service.EditProduct(context.Background(), writer, req)
		return
	}
}
