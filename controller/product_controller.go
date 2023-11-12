package controller

import (
	"context"
	"entity"
	"helper"
	"net/http"
	"service"
)

func AddorEditProduct(writer http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		productReq := entity.Products{}
		err := helper.DecodeRequest(req, &productReq)
		helper.PanicIfError(err)

		service.AddProduct(context.Background(), productReq, writer, req)
		return
	}
}
