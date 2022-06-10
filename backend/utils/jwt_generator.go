package utils

import (
	"os"
	"sister-backend/app/model"
	"sister-backend/exception"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

// GenerateNewAccessToken func for generate a new Access token.
func GenerateNewToken(user *model.AuthResponse, is_admin bool) (response model.JwtToken) {
	secretKey := os.Getenv("JWT_SECRET_KEY")

	accessTokenExpire, err := strconv.Atoi(os.Getenv("JWT_ACCESS_TOKEN_EXPIRE"))
	exception.PanicWhenError(err)

	refreshTokenExpire, err := strconv.Atoi(os.Getenv("JWT_REFRESH_TOKEN_EXPIRE"))
	exception.PanicWhenError(err)

	accessTokenClaims := jwt.MapClaims{
		"exp":              time.Now().Add(time.Minute * time.Duration(accessTokenExpire)).UnixMilli(),
		"user_id":          user.ID,
		"email":            user.Email,
		"is_admin":         is_admin,
		"is_refresh_token": false,
	}

	refreshTokenClaims := jwt.MapClaims{
		"exp":              time.Now().Add(time.Minute * time.Duration(refreshTokenExpire)).UnixMilli(),
		"user_id":          user.ID,
		"email":            user.Email,
		"is_admin":         is_admin,
		"is_refresh_token": true,
	}

	response.AccessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims).SignedString([]byte(secretKey))
	exception.PanicWhenError(err)

	response.RefreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims).SignedString([]byte(secretKey))
	exception.PanicWhenError(err)

	response.AccessTokenExpire = accessTokenExpire
	response.RefreshTokenExpire = refreshTokenExpire

	return response
}
