package webserver

import (
	"belanjabackend/webserver/controller"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (logMiddleware *LogMiddleware) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Token needed")
	logMiddleware.Handler.ServeHTTP(w, req)
}

func RunWeb() {
	router := httprouter.New()

	router.PanicHandler = func(w http.ResponseWriter, req *http.Request, i interface{}) {
		fmt.Fprint(w, "Wala")
	}

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "Tidak ditemukan, silahkan cari yang lain")
	})

	router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "Pakai http method yang lain, ini bukan untuk GET, coba pakai POST, PUT, atau DELETE, atau yang lain deh")
	})

	router.GET("/", controller.RootHandler)
	router.POST("/register", controller.RegisterUser)
	router.POST("/login", controller.LoginUser)
	router.GET("/logout", controller.LogoutUser)
	router.POST("/book", controller.InsertBook)
	router.GET("/book/:id", controller.GetBookById)

	logMiddleware := LogMiddleware{router}

	err := http.ListenAndServe("localhost:3000", &logMiddleware)
	if err != nil {
		panic(err)
	}
}
