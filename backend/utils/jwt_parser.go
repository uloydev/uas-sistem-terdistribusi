package utils

import (
	"os"
	"sister-backend/app/model"
	"sister-backend/exception"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

// ExtractTokenMetadata func to extract metadata from JWT.
func ExtractTokenMetadata(c *fiber.Ctx) *model.JWTMetadata {
	token := verifyToken(c)

	// Setting and checking token and credentials.
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return &model.JWTMetadata{
			Expires:        time.UnixMilli(claims["exp"].(int64)),
			UserId:         claims["user_id"].(uint),
			Email:          claims["email"].(string),
			IsAdmin:        claims["is_admin"].(bool),
			IsRefreshToken: claims["is_refresh_token"].(bool),
		}
	}

	return nil
}

func extractToken(c *fiber.Ctx) string {
	bearToken := c.Get("Authorization")

	// Normally Authorization HTTP header.
	onlyToken := strings.Split(bearToken, " ")
	if len(onlyToken) == 2 {
		return onlyToken[1]
	}

	return ""
}

func verifyToken(c *fiber.Ctx) *jwt.Token {
	tokenString := extractToken(c)

	token, err := jwt.Parse(tokenString, jwtKeyFunc)
	exception.PanicWhenError(err)

	return token
}

func jwtKeyFunc(token *jwt.Token) (any, error) {
	return []byte(os.Getenv("JWT_SECRET_KEY")), nil
}
