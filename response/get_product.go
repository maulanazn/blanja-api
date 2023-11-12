package response

import (
	"encoding/json"
	"entity"
)

type GetProducts struct {
	Status  int
	Message string
	Data    entity.Products
}

func ToGetProducts(status int, message string, data entity.Products) string {
	value, err := json.MarshalIndent(&GetProducts{
		Status:  status,
		Message: string(message),
		Data:    data,
	}, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(value)
}
