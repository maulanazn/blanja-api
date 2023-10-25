package response

type Data struct {
	UserName string
	Email    string
}

type RegisterResponse struct {
	Status  int
	Message string
	Data    Data
}
