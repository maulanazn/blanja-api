package order_details

import (
	"encoding/json"
	"util"
)

type OrderDetailResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    OrderDetail
}

func ToOrderDetail(code int, message string, data OrderDetailResponse) string {
	value, err := json.MarshalIndent(&OrderDetailResponse{
		Code:    code,
		Message: message,
		Data:    data.Data,
	}, "", "\t")
	if err != nil {
		util.Log2File(err.Error())
	}

	return string(value)
}
