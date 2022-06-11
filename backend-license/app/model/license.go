package model

type LicenseResponse struct {
	BasicData
	Username string `json:"username"`
	Key      string `json:"key"`
	IsActive bool   `json:"is_active"`
}

type LicenseRequest struct {
	Username string `json:"username"`
	Key      string `json:"key"`
	IsActive bool   `json:"is_active"`
}

type CheckLicenseRequest struct {
	Username string `json:"username"`
	Key      string `json:"key"`
}
