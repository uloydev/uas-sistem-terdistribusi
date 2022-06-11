package controller

import "github.com/gofiber/fiber/v2"

type BaseController interface {
	Route(*fiber.Group)
}
