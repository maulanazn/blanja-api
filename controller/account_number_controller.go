package controller

import (
	"helper"
	"net/http"
	"request"
	"service"
)

func AccountNumber(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		req.Header.Add("Content-Type", "application/json")
		addAccountNumberReq := request.AddAccountNumber{}
		if err := helper.DecodeData(req, &addAccountNumberReq); err != nil {
			http.Error(w, "Failed to decode json", http.StatusExpectationFailed)
			return
		}

		service.InsertAccountNumber(w, addAccountNumberReq, req)

		return
	}
}
