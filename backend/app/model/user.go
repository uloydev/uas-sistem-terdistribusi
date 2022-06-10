package model

type UserResponse struct {
	BasicData
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Birthday *int64  `json:"birthday"`
	Phone    *string `json:"phone"`
}

type UserRequest struct {
	Name     string  `json:"name" example:"toqcer"`
	Email    string  `json:"email" example:"user@toqcer.id"`
	Password string  `json:"password" example:"strongpassword"`
	Birthday *int64  `json:"birthday"`
	Phone    *string `json:"phone" example:"08512132332"`
}
