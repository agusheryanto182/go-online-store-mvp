package dto

type RegisterRequest struct {
	Fullname        string `json:"fullname" validate:"required"`
	Username        string `json:"username" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,eqfield=PasswordConfirm"`
	PasswordConfirm string `json:"password_confirm" validate:"required"`
}

type LoginRequest struct {
	Identifier string `json:"identifier" validate:"required"`
	Password   string `json:"password" validate:"required"`
}
