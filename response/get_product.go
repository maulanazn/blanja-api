package response

import (
	"encoding/json"
	"entity"
)

type GetProducts struct {
	Status  int `json:"status"`
	Message string `json:"message"` 
	Data []entity.Products
}

type GetProduct struct {
	Status  int `json:"status"`
	Message string `json:"message"` 
	Data entity.Products
}

func ToGetProducts(status int, message string, product GetProducts) string {
	value, err := json.MarshalIndent(&GetProducts{
		Status:  status,
		Message: string(message),
		Data: product.Data,
	}, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(value)
}

func ToGetProduct(status int, message string, product GetProduct) string {
	value, err := json.MarshalIndent(&GetProduct{
		Status:  status,
		Message: string(message),
		Data: product.Data,
	}, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(value)
}
