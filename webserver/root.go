package webserver

import (
	"fmt"
	"net/http"
	"userboilerplate-api/webserver/controller"
	"userboilerplate-api/webserver/middleware"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (logMiddleware *LogMiddleware) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Token needed")
	logMiddleware.Handler.ServeHTTP(w, req)
}

func RunWeb() {
	router := http.NewServeMux()

	router.HandleFunc("/", controller.RootHandler)
	router.HandleFunc("/register", controller.RegisterCustomer)
	router.HandleFunc("/login", controller.LoginCustomer)
	router.Handle("/customer", middleware.NewEntranceToken(controller.EditCustomer))
	router.Handle("/address", middleware.NewEntranceToken(controller.AddOrEditAddress))

	err := http.ListenAndServe("localhost:3000", router)
	if err != nil {
		panic(err)
	}
}
