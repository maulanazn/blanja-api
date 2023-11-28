package product

import (
	"encoding/json"
)

type GetProducts struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Products
}

type GetProductStruct struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Products
}

func ToGetProducts(status int, message string, product GetProducts) string {
	value, err := json.MarshalIndent(&GetProducts{
		Status:  status,
		Message: message,
		Data:    product.Data,
	}, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(value)
}

func ToGetProduct(status int, message string, product GetProductStruct) string {
	value, err := json.MarshalIndent(&GetProductStruct{
		Status:  status,
		Message: message,
		Data:    product.Data,
	}, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(value)
}
