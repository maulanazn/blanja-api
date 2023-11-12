package main

import (
	"controller"
	"middleware"
	"net/http"
	"repository"
)

func main() {
	router := http.NewServeMux()

	repository.InitDBPostgreSQL()

	router.HandleFunc("/", controller.RootHandler)
	router.HandleFunc("/register", controller.RegisterCustomer)
	router.HandleFunc("/login", controller.LoginCustomer)
	router.Handle("/user", middleware.NewEntranceToken(controller.EditCustomer))
	router.Handle("/address", middleware.NewEntranceToken(controller.AddOrEditAddress))
	router.Handle("/product", middleware.NewEntranceToken(controller.AddorEditProduct))

	err := http.ListenAndServe(":3000", router)
	if err != nil {
		panic(err)
	}
}
