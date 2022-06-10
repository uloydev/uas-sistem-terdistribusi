package model

type AdminResponse struct {
	BasicData
	Name  string `json:"name"`
	Email string `json:"email"`
}

type AdminRequest struct {
	Name     string `json:"name" example:"admin"`
	Email    string `json:"email" example:"admin@toqcer.id"`
	Password string `json:"password" example:"strongpassword"`
}
