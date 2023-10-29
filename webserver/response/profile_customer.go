package response

type ProfileCustomerData struct {
	Userimage   string `json:"userimage"`
	Username    string `json:"username"`
	Phone       int64  `json:"phone"`
	Gender      string `json:"gender"`
	Dateofbirth string `json:"dateofbirth"`
}

type ProfileCustomer struct {
	Status  int
	Message string
	Data    ProfileCustomerData
}
