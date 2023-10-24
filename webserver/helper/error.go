package helper

import "net/http"

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func BadStatusIfError(err error, writer http.ResponseWriter) {
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
}

func InternalServerErrorIfError(err error, writer http.ResponseWriter) {
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}
