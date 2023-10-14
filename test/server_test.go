package test

import (
	"belanjabackend/webserver/controller"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootFailed(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	controller.RootHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	formatted := string(body)

	assert.NotEqual(t, formatted, "Hello world")
}

func TestRootSuccess(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	controller.RootHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	formatted := string(body)

	assert.Equal(t, formatted, "Paybook backend")
}

func TestLoginFailed(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:3000/login", nil)
	recorder := httptest.NewRecorder()

	controller.RootHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	formatted := string(body)

	assert.NotEqual(t, formatted, "Login page")
}

func TestLoginSuccess(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:3000/login", nil)
	recorder := httptest.NewRecorder()

	controller.RootHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	formatted := string(body)

	assert.Equal(t, formatted, "Login page")
}
