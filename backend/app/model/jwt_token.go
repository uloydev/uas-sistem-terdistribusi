package model

type JwtToken struct {
	AccessToken        string `json:"access_token"`
	AccessTokenExpire  int    `json:"access_token_expire"`
	RefreshToken       string `json:"refresh_token"`
	RefreshTokenExpire int    `json:"refresh_token_expire"`
}
