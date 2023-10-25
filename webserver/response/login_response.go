package response

type Token struct {
	Value string
}

type LoginResponse struct {
	Status  int
	Message string
	Data    Token
}
