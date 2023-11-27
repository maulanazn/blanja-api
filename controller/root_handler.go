package controller

import (
	"fmt"
	"log"
	"net/http"
)

func RootHandler(writer http.ResponseWriter, req *http.Request) {
	if _, err := fmt.Fprint(writer, "User boiler plate backend"); err != nil {
		log.Println(err.Error())
	}
}
