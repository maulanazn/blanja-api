package response

type Data struct {
	Username string
	Email    string
	Roles    string
}

type RegisterResponse struct {
	Status  int
	Message string
	Data    Data
}
