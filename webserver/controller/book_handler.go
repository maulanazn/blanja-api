package controller

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func InsertBook(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	var user_id, title, description, writer, price string

	req.ParseForm()
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	namecookie, err := req.Cookie("name")
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	user_id += req.Form.Get("user_id")
	title += req.PostForm.Get("title")
	description += req.PostForm.Get("description")
	writer += req.PostForm.Get("writer")
	price += req.PostForm.Get("price")

	if user_id == "" || title == "" || description == "" || writer == "" || price == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Request is not valid")
		return
	}

	if namecookie != nil {
		fmt.Fprint(w, "Success Adding book")
		w.WriteHeader(http.StatusCreated)
		return
	}

	fmt.Fprint(w, "You must login")
}

func GetBookById(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id := params.ByName("id")

	fmt.Fprint(w, "Book id number : "+id)
}
