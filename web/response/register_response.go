package response

import "encoding/json"

type Data struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Roles    string `json:"roles"`
}

type RegisterResponse struct {
	Status  int
	Message string
	Data    Data
}

func ToResponseData(status int, message string, data Data) string {
	value, err := json.MarshalIndent(&RegisterResponse{
		Status:  status,
		Message: string(message),
		Data:    data,
	}, "", "")
	if err != nil {
		panic(err)
	}
	return string(value)
}
