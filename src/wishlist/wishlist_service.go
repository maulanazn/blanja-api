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
  token := util.DecodeToken(req.Header.Get("Authorization"));

  if err := util.DecodeRequestAndValidate(writer, req, &addWishlistRequest); err != nil {
    http.Error(writer, err.Error(), http.StatusBadRequest);
    return
  }
  
  wishlist_data := Wishlist{
    UserId: token,
    ProductId: addWishlistRequest.ProductId,
    Quantity: addWishlistRequest.Quantity, 
    CreatedAt: time.Now(),
  }

  if err := InsertWishlist(ctx, &wishlist_data); err != nil {
    http.Error(writer, err.Error(), http.StatusBadRequest);
    return
  }    
  
  writer.WriteHeader(200)
  writer.Write([]byte("Successfully added to your wishlist")) 
  return
}

func PutWishlist(ctx context.Context, writer http.ResponseWriter, req *http.Request) {
  putWishlistRequest := PutWishlistRequest{}

  if err := util.DecodeRequestAndValidate(writer, req, &putWishlistRequest); err != nil {
    http.Error(writer, err.Error(), http.StatusBadRequest);
    return
  }
  
  wishlist_data := &Wishlist{
    Quantity: putWishlistRequest.Quantity, 
    UpdatedAt: time.Now(),
  }

  if err := UpdateWishlist(ctx, *wishlist_data, req.URL.Query().Get("id")); err != nil {
    http.Error(writer, err.Error(), http.StatusBadRequest);
    return
  }    
  
  writer.WriteHeader(202)
  writer.Write([]byte("Successfully updated your wishlist")) 
  return
}

func GetWishlists(ctx context.Context, writer http.ResponseWriter, req *http.Request) {
  if req.URL.Query().Has("id") {
    resultwishlist, err := SelectWishlistById(ctx, req.URL.Query().Get("id"));
    if err != nil {
      http.Error(writer, err.Error(), http.StatusBadRequest);
      return
    }
 
    outputJsonID := ToWishlistId(200, "Here's your wishlist", ResponseWishlistId{
      Data: resultwishlist,
    })

    fmt.Fprint(writer, outputJsonID)
    return
  }

  token := util.DecodeToken(req.Header.Get("Authorization"));

  resultwishlist, err := SelectWishlistByUser(ctx, token);
  if err != nil {
    http.Error(writer, err.Error(), http.StatusBadRequest);
    return
  }

  for _, data := range resultwishlist {
    _, err := json.MarshalIndent(&data, " ", "\n")
    util.PanicIfError(err)
  }
  
  outputJson := ToWishlistUser(200, "Here's your wishlists", ResponseWishlistUser{
    Data: resultwishlist,
  })
  fmt.Fprint(writer, outputJson)
  return
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

  token := util.DecodeToken(req.Header.Get("Authorization"))
  if err := DeleteWishlistByUser(ctx, token); err != nil {
    http.Error(writer, err.Error(), http.StatusInternalServerError)
    return
  }

  writer.WriteHeader(202)
  writer.Write([]byte("Successfully deleted all your wishlist data")) 
  return
}
