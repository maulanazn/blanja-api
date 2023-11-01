package main

import (
	"net/http"
	"userboilerplate-api/webserver/controller"
	"userboilerplate-api/webserver/middleware"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", controller.RootHandler)
	router.HandleFunc("/register", controller.RegisterCustomer)
	router.HandleFunc("/login", controller.LoginCustomer)
	router.Handle("/customer", middleware.NewEntranceToken(controller.EditCustomer))
	router.Handle("/address", middleware.NewEntranceToken(controller.AddOrEditAddress))

	err := http.ListenAndServe(":3000", router)
	if err != nil {
		panic(err)
	}
}
