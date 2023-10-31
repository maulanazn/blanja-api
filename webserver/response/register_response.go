package response

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
