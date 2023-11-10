package main

import (
	"config"
	"controller"
	"entity"
	"middleware"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	config.GetConnection().AutoMigrate(&entity.Users{}, &entity.Address{})

	router.HandleFunc("/", controller.RootHandler)
	router.HandleFunc("/register", controller.RegisterCustomer)
	router.HandleFunc("/login", controller.LoginCustomer)
	router.Handle("/user", middleware.NewEntranceToken(controller.EditCustomer))
	router.Handle("/address", middleware.NewEntranceToken(controller.AddOrEditAddress))

	err := http.ListenAndServe(":3000", router)
	if err != nil {
		panic(err)
	}
}
