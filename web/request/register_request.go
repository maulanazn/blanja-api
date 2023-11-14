package request

type RegisterRequest struct {
	Username string `json:"username" validate:"alpha,alphanum"`
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"alpha,alphanum,contains=$%^&*@"`
	Roles    string `json:"roles" validate:"alpha,alphanum"`
}
