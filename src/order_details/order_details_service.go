package order_details

import (
	"address"
	"context"
	"fmt"
	"net/http"
	"product"
	"strconv"
	"users"
	"util"
)

func PostOrderDetail(ctx context.Context, writer http.ResponseWriter, req *http.Request) {
	insertOrderRequest := OrderDetailRequest{}
	util.DecodeRequestAndValidate(writer, req, &insertOrderRequest)

	userid := util.DecodeToken(req.Header.Get("Authorization"))

	userIdData, _ := users.SelectCustomerById(ctx, userid)
	isAddressId, _ := address.AddressById(ctx, insertOrderRequest.AddressId)
	productAndStoreData, _ := product.SelectProductById(ctx, insertOrderRequest.ProductId)

	if userIdData.Id != userid {
		writer.WriteHeader(400)
		writer.Write([]byte("Invalid user"))
		return
	}

	if isAddressId.Id != insertOrderRequest.AddressId {
		writer.WriteHeader(400)
		writer.Write([]byte("Invalid address"))
		return
	}

	if productAndStoreData.ProductId.String() != insertOrderRequest.ProductId {
		writer.WriteHeader(400)
		writer.Write([]byte("Invalid product"))
		return
	}

	if productAndStoreData.StoreName != insertOrderRequest.StoreName {
		writer.WriteHeader(400)
		writer.Write([]byte("Invalid store"))
		return
	}

	orderDetail := OrderDetail{
		UserId:    userIdData.Id,
		AddressId: isAddressId.Id,
		ProductId: productAndStoreData.ProductId.Hex(),
		StoreName: productAndStoreData.StoreName,
		Quantity:  insertOrderRequest.Quantity,
		Price:     insertOrderRequest.Price,
		SubTotal:  insertOrderRequest.Quantity * insertOrderRequest.Price,
	}

	if err := InsertOrderDetail(ctx, &orderDetail); err != nil {
		writer.WriteHeader(500)
		failedResponse := util.ToWebResponse(500, http.StatusInternalServerError)
		if _, err := fmt.Fprint(writer, failedResponse); err != nil {
			util.Log2File(err.Error())
			return
		}

		return
	}

	writer.WriteHeader(201)
	successResponse := util.ToWebResponse(200, "Success adding new order")
	if _, err := fmt.Fprint(writer, successResponse); err != nil {
		util.Log2File(err.Error())
		return
	}
}

func GetOrderDetailById(ctx context.Context, writer http.ResponseWriter, req *http.Request) {
	order_id, _ := strconv.Atoi(req.URL.Query().Get("order_id"))
	result, err := SelectOrderDetailById(ctx, order_id)
	if err != nil {
		writer.WriteHeader(500)
		failedResponse := util.ToWebResponse(500, http.StatusInternalServerError)
		if _, err := fmt.Fprint(writer, failedResponse); err != nil {
			util.Log2File(err.Error())
			return
		}

		return
	}

	writer.WriteHeader(200)
	successResponse := ToOrderDetail(200, "Success getting order id", OrderDetailResponse{
		Data: result,
	})

	if _, err := fmt.Fprint(writer, successResponse); err != nil {
		util.Log2File(err.Error())
		return
	}
}
