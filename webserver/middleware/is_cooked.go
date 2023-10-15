package middleware

import (
	"fmt"
	"net/http"
)

type CookieMiddleware struct {
	Handler http.Handler
}

func (cookieMiddleware *CookieMiddleware) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	fmt.Println("before middleware")
	cookieMiddleware.Handler.ServeHTTP(writer, req)
	fmt.Println("after middleware")
}
