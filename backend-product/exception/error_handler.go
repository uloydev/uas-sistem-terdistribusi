package exception

import (
	"sister-backend/app/model"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {

	_, ok := err.(ValidationError)
	if ok {
		return ctx.JSON(model.WebResponse{
			Code:   400,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		})
	}

	return ctx.JSON(model.WebResponse{
		Code:   500,
		Status: "INTERNAL_SERVER_ERROR",
		Data:   err.Error(),
	})
}

func JwtError(ctx *fiber.Ctx, err error) error {

	if err.Error() == "Missing or malformed JWT" {
		return ctx.JSON(model.WebResponse{
			Code:   400,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		})
	}
	return ctx.JSON(model.WebResponse{
		Code:   401,
		Status: "UNAUTHORIZED",
		Data:   err.Error(),
	})
}
