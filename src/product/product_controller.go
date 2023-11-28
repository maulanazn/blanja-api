package product

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func AddorEditProduct(writer http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		AddProduct(context.Background(), writer, req)
		return
	case http.MethodGet:
		GetProduct(context.Background(), writer, req)
		return
	case http.MethodPut:
		EditProduct(context.Background(), writer, req)
		return
	default:
		if _, err := fmt.Fprintln(writer, "The thing that you request is not available"); err != nil {
			log.Println(err.Error())
		}
	}
}
