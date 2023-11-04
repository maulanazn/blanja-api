package response

import "encoding/json"

type Token struct {
	Token string
}

type LoginResponse struct {
	Status  int
	Message string
	Data    Token
}

func ToResponseToken(status int, message string, token Token) string {
	value, err := json.MarshalIndent(&LoginResponse{
		Status:  status,
		Message: string(message),
		Data:    token,
	}, "", "")
	if err != nil {
		panic(err)
	}
	return string(value)
}
