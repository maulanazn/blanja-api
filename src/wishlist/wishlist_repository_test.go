package wishlist_test

import (
	"context"
  "encoding/json"
  "testing"
	"time"
	"wishlist"
)

func TestCreateWishlist(t *testing.T) {
  wishlist_data := wishlist.Wishlist{
    UserId    : "1ad31429-abbf-441b-9905-6c2b5b78fc3b",
    ProductId : "655289c52e15c8ba4c517f60",
    Quantity  : 100,
    CreatedAt : time.Now(),
    UpdatedAt : time.Now().AddDate(2023, 10, 9),
    DeletedAt : time.Now().AddDate(2025, 12, 12),
  }

  if err := wishlist.InsertWishlist(context.Background(), &wishlist_data); err != nil {
    t.Log(err.Error())
    return
  } 

  t.Log("Success create wishlist data")
}

func TestSelectWishlistofUser(t *testing.T) {
  timeout, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
  result, err := wishlist.SelectWishlistByUser(timeout, "1ad31429-abbf-441b-9905-6c2b5b78fc3b")
  defer cancel()
  if err != nil {
    t.Error(err.Error())
  }
  
  for _, data := range result {
    formatted, _ := json.Marshal(data)
    t.Log(string(formatted))
  }
}

func TestSelectWishlistId(t *testing.T) {
  timeout, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
  result, err := wishlist.SelectWishlistById(timeout, "b5c2a5c1-9991-41e6-b7f9-137c539f5032")
  defer cancel()
  if err != nil {
    t.Error(err.Error())
  }
  
  if result.WishlistId == "" {
    t.Log("Data Nil")
    return
  }
  
  formatted, _ := json.Marshal(result)
  t.Log(string(formatted))
}

func TestUpdateWishlist(t *testing.T) {
  wishlist_data := wishlist.Wishlist{
    UserId    : "1ad31429-abbf-441b-9905-6c2b5b78fc3b",
    ProductId : "6552bec51e15c8ba4c517f73",
    Quantity  : 29,
    CreatedAt : time.Now().AddDate(2024, 10, 9),
    UpdatedAt : time.Now(),
    DeletedAt : time.Now().AddDate(2025, 12, 12),
  }

  if err := wishlist.UpdateWishlist(context.Background(), wishlist_data, "6e678ef8-5b8e-4002-afa6-9c6f6593ff33"); err != nil {
    t.Log(err.Error())
    return
  } 

  t.Log("Success updating wishlist data")
}

func TestDeleteWishlistofUser(t *testing.T) {
  timeout, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
  if err := wishlist.DeleteWishlistByUser(timeout, "1ad31429-abbf-441b-9905-6c2b5b78fc3b"); err != nil {
    t.Error(err.Error())
  }
  
  defer cancel()
  
  t.Log("Sucessfully delete user wishlist")
}

func TestDeleteWishlistId(t *testing.T) {
  timeout, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
  if err := wishlist.DeleteWishlistById(timeout, "b5c2a5c1-9991-41e6-b7f9-137c539f5032"); err != nil {
    t.Error(err.Error())
  }
  
  defer cancel()
  
  t.Log("Sucessfully delete the requested data")
}