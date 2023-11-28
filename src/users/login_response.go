package users

import "encoding/json"

type Token struct {
	Token string `json:"token"`
}

type LoginResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Token
}

func ToResponseToken(status int, message string, token Token) string {
	value, err := json.MarshalIndent(&LoginResponse{
		Status:  status,
		Message: message,
		Data:    token,
	}, "", "")
	if err != nil {
		panic(err)
	}
	return string(value)
}
