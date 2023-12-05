package wishlist

import (
	"fmt"
	"net/http"
	"util"
)

func WishlistController(writer http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		PostWishlist(req.Context(), writer, req)
		return
	case http.MethodPut:
		PutWishlist(req.Context(), writer, req)
		return
	case http.MethodGet:
		GetWishlists(req.Context(), writer, req)
		return
	case http.MethodDelete:
		DeleteWishlist(req.Context(), writer, req)
		return
	default:
		if _, err := fmt.Fprint(writer, "Not supported"); err != nil {
			util.Log2File(err.Error())
			return
		}
		return
	}
}
