package controller

import (
	"fmt"
	"net/http"
	"time"
)

func RegisterUser(writer http.ResponseWriter, req *http.Request) {
	var username, email, password string
	if req.Method == "POST" {
		req.ParseForm()
		req.Header.Add("content-type", "application/x-www-form-urlencoded")

		username += req.PostForm.Get("username")
		email += req.PostForm.Get("email")
		password += req.PostForm.Get("password")

		if username == "" || email == "" || password == "" {
			writer.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(writer, "Request is not valid")
			return
		}

		writer.WriteHeader(http.StatusCreated)
		fmt.Fprint(writer, "Success Register")
	}
}

func LoginUser(writer http.ResponseWriter, req *http.Request) {
	var email, password string
	var cookie http.Cookie = http.Cookie{}
	if req.Method == "POST" {
		req.ParseForm()
		req.Header.Add("content-type", "application/x-www-form-urlencoded")

		email += req.PostForm.Get("email")
		password += req.PostForm.Get("password")

		cookie.Name = "email"
		cookie.Value = email
		cookie.Secure = true
		cookie.Path = "/"
		cookie.Expires = time.Now().Add(7 * 24 * time.Hour)
		http.SetCookie(writer, &cookie)

		_, err := req.Cookie("name")
		if err != nil {
			panic(err)
		} else {
			fmt.Fprint(writer, "You are already login")
			return
		}

		if email == "" || password == "" {
			writer.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(writer, "Request is not valid")
			return
		}

		writer.WriteHeader(http.StatusCreated)
		fmt.Fprint(writer, "Success login")
	}
}

func LogoutUser(writer http.ResponseWriter, req *http.Request) {
	namecookie, err := req.Cookie("name")
	if err != nil {
		fmt.Fprint(writer, err)
		return
	}

	if namecookie != nil {
		fmt.Fprint(writer, "Logout success")
		return
	}

	writer.WriteHeader(http.StatusNetworkAuthenticationRequired)
	fmt.Fprint(writer, "You are not login  yet")
}
