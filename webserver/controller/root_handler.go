package controller

import (
	"fmt"
	"net/http"
)

func RootHandler(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprint(writer, "User boiler plate backend")
}
