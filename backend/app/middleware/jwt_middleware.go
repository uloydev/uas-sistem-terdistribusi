package middleware

import (
	"os"
	"sister-backend/exception"

	"github.com/gofiber/fiber/v2"

	jwtMiddleware "github.com/gofiber/jwt/v2"
)

func JWTProtected() func(*fiber.Ctx) error {
	config := jwtMiddleware.Config{
		SigningKey:   []byte(os.Getenv("JWT_SECRET_KEY")),
		ContextKey:   "jwt",
		ErrorHandler: exception.JwtError,
	}

	return jwtMiddleware.New(config)
}
