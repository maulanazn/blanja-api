package address

import (
	"fmt"
	"net/http"
	"util"
)

func AddOrEditAddress(writer http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		AddAddress(req.Context(), writer, req)
		return
	case http.MethodPut:
		EditAddress(req.Context(), writer, req)
		return
	case http.MethodGet:
		AddressDetail(req.Context(), writer, req)
		return
	default:
		if _, err := fmt.Fprint(writer, "The thing that you request is not available"); err != nil {
			util.Log2File(err.Error())
			return
		}
		return
	}

}
