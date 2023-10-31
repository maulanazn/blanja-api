package test

import (
	"belanjabackend/webserver/controller"
	"belanjabackend/webserver/helper"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestRegister(t *testing.T) {
	data := strings.NewReader(`
		{
			"username": "aditya",
			"email": "aditya553@mail.com",
			"password": "aditya553"
		}
	`)
	req := httptest.NewRequest("POST", "http://localhost:3000/register", data)
	recorder := httptest.NewRecorder()
	controller.RegisterCustomer(recorder, req)

	req.Header.Add("Content-Type", "application/json")

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TestLogin(t *testing.T) {
	data := strings.NewReader(`
		{
			"email": "maulanazn19@mail.com",
			"password": "maulanazn123"
		}
	`)
	req := httptest.NewRequest("POST", "http://localhost:3000/login", data)
	recorder := httptest.NewRecorder()
	controller.LoginCustomer(recorder, req)

	req.Header.Add("Content-Type", "application/json")

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TestEditCustomer(t *testing.T) {
	data := strings.NewReader(`
		{
			"email": "maulanazn19@mail.com",
			"password": "maulanazn123"
		}
	`)
	req := httptest.NewRequest("POST", "http://localhost:3000/customer", data)
	recorder := httptest.NewRecorder()
	controller.EditCustomer(recorder, req)

	req.Header.Add("Content-Type", "application/json")

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TestGetAddressUser(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:3000/address", nil)
	recorder := httptest.NewRecorder()

	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im1hdWxhbmF6bjE5QG1haWwuY29tIiwiZXhwIjoiMjAyMy0xMC0yNlQyMTo0ODo1MC45ODc2MjEzNzMrMDc6MDAiLCJ1c2VybmFtZSI6IiJ9.pYwNsO9IJuoM2b2g2q4Z7O-QDzZ5P9zJvnPQzgDUsGI")
	cookie := http.Cookie{
		Name:     "USR_ID",
		Value:    "0151cf30fbd5456aa30a3e5af3ccba18",
		Path:     "/",
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}
	req.AddCookie(&cookie)

	req.Cookie("USR_ID")

	controller.AddOrEditAddress(recorder, req)

	req.Header.Add("Content-Type", "application/json")

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TestEditAddress(t *testing.T) {
	data := strings.NewReader(`
		{
			"addresstype":    "office",
			"recipientname":  "maulanazn",
			"recipientphone": "1892382",
			"addressname":    "tes",
			"postalcode":     "91283A",
			"city":           "KualaLumpur"
		}
	`)
	id := "51ac602e02534e6a813b96c509b9b429"
	req := httptest.NewRequest("PUT", "http://localhost:3000/address?id="+id, data)
	recorder := httptest.NewRecorder()

	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im1hdWxhbmF6bjE5QG1haWwuY29tIiwiZXhwIjoiMjAyMy0xMC0yNlQyMTo0ODo1MC45ODc2MjEzNzMrMDc6MDAiLCJ1c2VybmFtZSI6IiJ9.pYwNsO9IJuoM2b2g2q4Z7O-QDzZ5P9zJvnPQzgDUsGI")
	cookie := http.Cookie{
		Name:     "USR_ID",
		Value:    "0151cf30fbd5456aa30a3e5af3ccba18",
		Path:     "/",
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}
	req.AddCookie(&cookie)

	req.Cookie("USR_ID")

	controller.AddOrEditAddress(recorder, req)

	req.Header.Add("Content-Type", "application/json")

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TestUploadCloudinary(t *testing.T) {
	file, err := os.Open("/home/maulanazn/Pictures/Notification.png")
	resultimage, err := helper.UploadCloudinary(file)
	if err != nil {
		panic(err)
	}
	fmt.Println(resultimage)
}
