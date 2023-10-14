package test

import (
	"belanjabackend/webserver/controller"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRegisterUser(t *testing.T) {
	data := strings.NewReader("username=maulana&email=maulana@mail.com&password=123")
	request := httptest.NewRequest("POST", "http://localhost:3000/register/", data)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	controller.RegisterUser(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Success Register", string(body))
}

func TestLoginUser(t *testing.T) {
	data := strings.NewReader("email=maulana@mail.com&password=123")
	request := httptest.NewRequest("POST", "http://localhost:3000/login", data)
	recorder := httptest.NewRecorder()

	request.Header.Add("content-type", "application/x-www-form-urlencoded")
	cookie := &http.Cookie{
		Name:    "name",
		Value:   "maulana@mail.com",
		Path:    "/",
		Secure:  true,
		Expires: time.Now().Add(7 * 24 * time.Hour),
	}

	request.AddCookie(cookie)

	controller.LoginUser(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Success login", string(body))
}

func TestLogoutUser(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:3000/logout", nil)
	recorder := httptest.NewRecorder()

	request.Header.Add("content-type", "application/x-www-form-urlencoded")
	cookie := &http.Cookie{
		Name:    "name",
		Value:   "maulana@mail.com",
		Path:    "/",
		Secure:  true,
		Expires: time.Now().Add(7 * 24 * time.Hour),
	}

	request.AddCookie(cookie)

	controller.LogoutUser(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	assert.Equal(t, "Logout success", string(body))
}
