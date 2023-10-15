package controller

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func RootHandler(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	fmt.Fprint(writer, "Paybook backend")
}
