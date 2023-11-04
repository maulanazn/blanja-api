package service

import (
	"entity"
	"fmt"
	"helper"
	"net/http"
	"os"
	"repository"
	"request"
	"response"
	"time"
)

func InsertAccountNumber(writer http.ResponseWriter, addAccountNumberReq request.AddAccountNumber, req *http.Request) {
	signtext := os.Getenv("ENC_KEY")
	acc_number, err := helper.Encrypt(signtext, addAccountNumberReq.AccNumber)
	if err != nil {
		http.Error(writer, "Failed to encrypt", 400)
		return
	}

	/* transfer scenario
	* 1992-93-210923-0 --> Encrypt --> check if exists or not --> Decrypt --> send data
	 */
	user_id, userid_err := req.Cookie("USR_ID")
	if userid_err != nil {
		http.Error(writer, "Cannot get user_id from cookie, please login", 400)
		return
	}

	account_number := entity.AccountNumber{
		AccNumber:      acc_number,
		AccOwner:       user_id.Value,
		AccDateCreated: time.Now(),
		AccBalance:     addAccountNumberReq.AccBalance,
	}

	if err := repository.InsertAccountNumber(req.Context(), &account_number); err != nil {
		failedResponse := response.ToWebResponse(400, "Failed to insert data, please repeat", writer)
		fmt.Fprint(writer, failedResponse)
		return
	}

	successResponse := response.ToWebResponse(201, "Successfully added your account number", writer)
	fmt.Fprint(writer, successResponse)
}
