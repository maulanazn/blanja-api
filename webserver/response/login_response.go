package response

type Token struct {
	Token string
}

type LoginResponse struct {
	Status  int
	Message string
	Data    Token
}
