package request

type RegisterRequest struct {
	Username string `json:"username" validate:"required,alphanum|alpha"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,ascii"`
	Roles    string `json:"roles" validate:"required,alpha"`
}
