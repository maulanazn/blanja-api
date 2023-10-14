package controller

import (
	"fmt"
	"net/http"
)

func InsertBook(w http.ResponseWriter, req *http.Request) {
	var user_id, title, description, writer, price string

	if req.Method == "POST" {
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
}
