package model

type LicenseCheckRequest struct {
	Key      string `json:"key"`
	Username string `json:"username"`
}
