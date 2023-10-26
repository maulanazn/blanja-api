package helper

import (
	"fmt"
	"log"
	"net/http"
)

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func FatalIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func BadStatusIfError(err error, writer http.ResponseWriter) {
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		failedResponse := ToWebResponse(400, "Wrong Password")
		fmt.Fprint(writer, failedResponse)
		return
	}
}

func InternalServerErrorIfError(err error, writer http.ResponseWriter) {
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}
