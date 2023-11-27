package response

import "encoding/json"

type ProfileCustomerData struct {
	Userimage   string `json:"userimage"`
	Username    string `json:"username"`
	Phone       int64  `json:"phone"`
	Gender      string `json:"gender"`
	Dateofbirth string `json:"dateofbirth"`
}

type ProfileCustomer struct {
	Status  int `json:"status"`
	Message string `json:"message"`
	Data    ProfileCustomerData
}

func ToProfileCustomer(status int, message string, data ProfileCustomer) string {
	value, err := json.MarshalIndent(&ProfileCustomer{
		Status:  status,
		Message: message,
		Data: ProfileCustomerData{
			Userimage:   data.Data.Userimage,
			Username:    data.Data.Username,
			Phone:       data.Data.Phone,
			Gender:      data.Data.Gender,
			Dateofbirth: data.Data.Dateofbirth,
		},
	}, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(value)
}
