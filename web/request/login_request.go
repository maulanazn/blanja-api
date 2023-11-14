package request

type LoginRequest struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"alpha,alphanum,contains=$%^&*@"`
}
