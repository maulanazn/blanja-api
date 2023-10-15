package test

import (
	"belanjabackend/webserver/controller"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestInsertBook(t *testing.T) {
	router := httprouter.New()
	router.POST("/book", controller.InsertBook)

	data := strings.NewReader("user_id=123&title=iniapaya&description=satu&writer=maulana&price=10")
	request := httptest.NewRequest("POST", "http://localhost:3000/book", data)
	request.Header.Add("content-type", "application/x-www-form-urlencoded")
	cookiedata := &http.Cookie{
		Name:    "name",
		Value:   "maulana@mail.com",
		Path:    "/",
		Secure:  true,
		Expires: time.Now().Add(7 * 24 * time.Hour),
	}

	request.AddCookie(cookiedata)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "Success Adding book", string(body), "Must be valid message")
}

func TestGetBookById(t *testing.T) {
	router := httprouter.New()
	router.GET("/book/:id", controller.GetBookById)

	request := httptest.NewRequest("GET", "http://localhost:3000/book/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	assert.Equal(t, "Book id number : 1", string(body), "Must be valid message")
}
