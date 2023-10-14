package webserver

import (
	"belanjabackend/webserver/controller"
	"net/http"
)

func RunWeb() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", controller.RootHandler)
	mux.HandleFunc("/register", controller.RegisterUser)
	mux.HandleFunc("/login", controller.LoginUser)
	mux.HandleFunc("/logout", controller.LogoutUser)
	mux.HandleFunc("/book", controller.InsertBook)

	var server http.Server = http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
