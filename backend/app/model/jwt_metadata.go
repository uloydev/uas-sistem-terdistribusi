package model

import "time"

type JWTMetadata struct {
	Expires        time.Time
	UserId         uint
	Email          string
	IsAdmin        bool
	IsRefreshToken bool
}
