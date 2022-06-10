package model

type AuthRequest struct {
	Email    string `json:"email" example:"user@toqcer.id"`
	Password string `json:"password" example:"strongpassword"`
}

type AuthResponse struct {
	ID       uint
	Email    string
	Password string
}
