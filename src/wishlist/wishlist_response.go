package wishlist

import (
	"encoding/json"
	"log"
)

type ResponseWishlistUser struct {
  Code int `json:"code"`
  Message string `json:"status"`
  Data []Wishlist `json:"data"`
}

type ResponseWishlistId struct {
  Code int `json:"code"`
  Message string `json:"status"`
  Data Wishlist `json:"data"`
}

func ToWishlistUser(code int, message string, data ResponseWishlistUser) string {
  value, err := json.MarshalIndent(&ResponseWishlistUser{
    Code: code,
    Message: message,
    Data: data.Data,
  }, "", "\t")
  
  if err != nil {
    log.Println(err.Error())
  }

  return string(value)
}

func ToWishlistId(code int, message string, data ResponseWishlistId) string {
  value, err := json.MarshalIndent(&ResponseWishlistId{
    Code: code,
    Message: message,
    Data: data.Data,
  }, "", "\t")
  
  if err != nil {
    log.Println(err.Error())
  }

  return string(value)
}
