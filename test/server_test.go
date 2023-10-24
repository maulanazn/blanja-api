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
