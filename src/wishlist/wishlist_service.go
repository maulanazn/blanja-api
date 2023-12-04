package wishlist

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"util"
)

func PostWishlist(ctx context.Context, writer http.ResponseWriter, req *http.Request) {
	addWishlistRequest := AddWishlistRequest{}
	token := util.DecodeToken(req.Header.Get("Authorization"))

	util.DecodeRequestAndValidate(writer, req, &addWishlistRequest)

	wishlist_data := Wishlist{
		UserId:    token,
		ProductId: addWishlistRequest.ProductId,
		StoreName: addWishlistRequest.StoreName,
		Quantity:  addWishlistRequest.Quantity,
		CreatedAt: time.Now(),
	}

	if err := InsertWishlist(ctx, &wishlist_data); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	writer.WriteHeader(200)
	writer.Write([]byte("Successfully added to your wishlist"))
}

func PutWishlist(ctx context.Context, writer http.ResponseWriter, req *http.Request) {
	putWishlistRequest := PutWishlistRequest{}

	util.DecodeRequestAndValidate(writer, req, &putWishlistRequest)

	wishlist_data := &Wishlist{
		Quantity:  putWishlistRequest.Quantity,
		UpdatedAt: time.Now(),
	}

	if err := UpdateWishlist(ctx, *wishlist_data, req.URL.Query().Get("id")); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	writer.WriteHeader(202)
	writer.Write([]byte("Successfully updated your wishlist"))
}

func GetWishlists(ctx context.Context, writer http.ResponseWriter, req *http.Request) {
	if req.URL.Query().Has("id") {
		resultwishlist, err := SelectWishlistById(ctx, req.URL.Query().Get("id"))
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		outputJsonID := ToWishlistId(200, "Here's your wishlist", ResponseWishlistId{
			Data: resultwishlist,
		})

		fmt.Fprint(writer, outputJsonID)
		return
	}

	if req.URL.Query().Has("store_name") {
		resultwishlist, err := SelectWishlistByStore(ctx, req.URL.Query().Get("store_name"))
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		for _, data := range resultwishlist {
			_, err := json.MarshalIndent(&data, " ", "\n")
			util.PanicIfError(err)
		}

		outputJson := ToWishlists(200, "Here's your wishlists", ResponseWishlistUser{
			Data: resultwishlist,
		})
		fmt.Fprint(writer, outputJson)
		return
	}

	token := util.DecodeToken(req.Header.Get("Authorization"))

	resultwishlist, err := SelectWishlistByUser(ctx, token)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	for _, data := range resultwishlist {
		_, err := json.MarshalIndent(&data, " ", "\n")
		util.PanicIfError(err)
	}

	outputJson := ToWishlists(200, "Here's your wishlists", ResponseWishlistUser{
		Data: resultwishlist,
	})
	fmt.Fprint(writer, outputJson)
}

func DeleteWishlist(ctx context.Context, writer http.ResponseWriter, req *http.Request) {
	if req.URL.Query().Has("id") {
		if err := DeleteWishlistById(ctx, req.URL.Query().Get("id")); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		writer.WriteHeader(202)
		writer.Write([]byte("Successfully deleted the requested data"))
		return
	}

	if req.URL.Query().Has("store_name") {
		if err := DeleteWishlistByStore(ctx, req.URL.Query().Get("store_name")); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		writer.WriteHeader(202)
		writer.Write([]byte("Successfully deleted the requested data"))
		return
	}

	token := util.DecodeToken(req.Header.Get("Authorization"))
	if err := DeleteWishlistByUser(ctx, token); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(202)
	writer.Write([]byte("Successfully deleted all your wishlist data"))
}
