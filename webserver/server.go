package webserver

import (
	"belanjabackend/webserver/controller"
	"fmt"
	"net/http"
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

	logMiddleware := LogMiddleware{router}

	err := http.ListenAndServe("localhost:3000", &logMiddleware)
	if err != nil {
		panic(err)
	}
}
