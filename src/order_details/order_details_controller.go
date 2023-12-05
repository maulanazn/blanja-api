package order_details

import (
	"fmt"
	"net/http"
	"util"
)

func OrderDetailController(writer http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		PostOrderDetail(req.Context(), writer, req)
		return
	case http.MethodGet:
		GetOrderDetailById(req.Context(), writer, req)
		return
	default:
		if _, err := fmt.Fprint(writer, "No such thing like that, get away"); err != nil {
			util.Log2File(err.Error())
			return
		}
		return
	}
}
