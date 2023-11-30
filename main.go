package main

import (
	"address"
	"autoload"
	"fmt"
	"log"
	"middleware"
	"net/http"
	"product"
	"users"
	"wishlist"
)

func RootHandler(writer http.ResponseWriter, req *http.Request) {
	if _, err := fmt.Fprint(writer, "User boiler plate backend"); err != nil {
		log.Println(err.Error())
	}
}

func main() {
	router := http.NewServeMux()

	autoload.InitDBPostgreSQL()

	router.HandleFunc("/", RootHandler)
	router.HandleFunc("/register", users.RegisterCustomer)
	router.HandleFunc("/login", users.LoginCustomer)
	router.Handle("/user", middleware.NewEntranceToken(users.PutCustomer))
	router.Handle("/address", middleware.NewEntranceToken(address.AddOrEditAddress))
	router.Handle("/product", middleware.NewEntranceToken(product.AddorEditProduct))
  router.Handle("/wishlist", middleware.NewEntranceToken(wishlist.WishlistController))

	err := http.ListenAndServe(":3000", router)
	if err != nil {
		panic(err)
	}
}
